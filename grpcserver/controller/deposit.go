/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: notify.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	"git.bhex.io/bhpc/wallet/chainnode/util/walletclient"
	pbcommon "git.bhex.io/bhpc/wallet/common/protos/common"
	"git.bhex.io/bhpc/wallet/common/protos/wallet"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/connectivity"
)

// NotifyDeposit call by httpserver
//
func NotifyDeposit(deposit common.NotifyDepositParams, apikey, sign, timestamp string) error {
	log.Infof("NotifyDeposit deposit:%v,apikey:%v,sign:%v,timestamp:%v", deposit, apikey, sign, timestamp)

	status, err := queryTokenStatus(deposit.TokenID)
	if err != nil {
		log.Errorf("NotifyDeposit get token status failed:%v", err)
		return err
	}

	if !common.CheckBlockTime(deposit.BlockTime) {
		log.Errorf("NotifyDeposit ErrinvalidBlockTime")
		return common.ErrinvalidBlockTime
	}

	if !status.CanDeposit {
		return common.ErrPauseDeposit
	}

	if !common.CheckToken(deposit.TokenID) {
		log.Errorf("NotifyDeposit ErrInvalidTokenID")
		return common.ErrInvalidTokenID
	}

	tokenInfo := common.GetTokenInfoByTokenID(deposit.TokenID)
	if tokenInfo == nil {
		log.Errorf("NotifyDeposit ErrInvalidTokenID")
		return common.ErrInvalidTokenID
	}
	log.Infof("GetTokenInfoByTokenID info:%v", *tokenInfo)

	if tokenInfo.APIKey != apikey {
		log.Errorf("NotifyDeposit ErrInvalidAPIKey")
		return common.ErrInvalidAPIKey
	}

	if tokenInfo.NeedMemo && deposit.Memo == "" {
		log.Errorf("NotifyDeposit ErrNeedMemo")
		return common.ErrNeedMemo
	}

	if strings.TrimSpace(deposit.Memo) != deposit.Memo {
		log.Errorf("NotifyDeposit ErrMemo")
		return common.ErrInvalidMemo
	}

	// check amount
	amount, _ := decimal.NewFromString(deposit.Amount)
	if amount.Cmp(decimal.Zero) == 0 || amount.Cmp(decimal.Zero) < 0 {
		log.Errorf("NotifyDeposit ErrInvalidAmount")
		return common.ErrInvalidAmount
	}
	if amount.String() != deposit.Amount {
		log.Errorf("NotifyDeposit ErrInvalidAmount, amount:%v, deposit amount:%v", amount.String(), deposit.Amount)
		return common.ErrInvalidAmount
	}

	// check amount decimal
	ext := int(amount.Exponent())
	log.Infof("NotifyDeposit deposit amount decimal:%v", ext)

	if tokenInfo.Decimals < ext {
		log.Errorf("NotifyDeposit ErrInvalidDecimals")
		return common.ErrInvalidDecimals
	}

	blockTime, _ := decimal.NewFromString(deposit.BlockTime)
	if blockTime.Cmp(decimal.Zero) == 0 || blockTime.Cmp(decimal.Zero) < 0 {
		log.Errorf("NotifyDeposit ErrinvalidBlockTime")
		return common.ErrinvalidBlockTime
	}

	blockHeight, _ := decimal.NewFromString(deposit.BlockHeight)
	if blockHeight.Cmp(decimal.Zero) == 0 || blockHeight.Cmp(decimal.Zero) < 0 {
		log.Errorf("NotifyDeposit ErrInvalidBlockHeight")
		return common.ErrInvalidBlockHeight
	}

	index, _ := decimal.NewFromString(deposit.Index)
	if index.Cmp(decimal.Zero) < 0 || index.Cmp(decimal.New(10000, 0)) > 0 {
		log.Errorf("NotifyDeposit ErrInvalidIndex")
		return common.ErrInvalidIndex
	}

	if deposit.TxHash == "" {
		log.Errorf("NotifyDeposit ErrInvalidTxHash")
		return common.ErrInvalidTxHash
	}

	if deposit.From == "" {
		log.Errorf("NotifyDeposit ErrInvalidFromAddress")
		return common.ErrInvalidFromAddress
	}

	var exist bool

	redisKey := fmt.Sprintf(common.UsedAddressSet, tokenInfo.Chain)
	common.FuncWithRetry(5, func() error {
		exist, err = common.Redis.IsSetMember(redisKey, deposit.To)
		return err
	})

	log.Infof("to address:%v, redis_key:%v, exist:%v", deposit.To, redisKey, exist)

	extension := map[string]string{
		"api_key":   apikey,
		"sign":      sign,
		"method":    "POST",
		"url":       "/api/v1/notify/deposit",
		"timestamp": timestamp,
	}

	if !exist {
		log.Errorf("NotifyDeposit ErrInvalidToAddress")
		return common.ErrInvalidToAddress
	}

	log.Infof("to address:%v in redis", deposit.To)

	height, _ := strconv.ParseUint(deposit.BlockHeight, 10, 64)
	time, _ := strconv.ParseUint(deposit.BlockTime, 10, 64)

	err = common.FuncWithRetry(2, func() error {
		return notifyDeposit(&wallet.NotifyDepositRequest{
			TokenId:     deposit.TokenID,
			TxHash:      deposit.TxHash,
			From:        deposit.From,
			To:          deposit.To,
			Memo:        deposit.Memo,
			Value:       deposit.Amount,
			BlockHeight: height,
			BlockTime:   time,
			Index:       uint32(index.IntPart()),
			Extension:   extension,
		})
	})
	if err != nil {
		log.Errorf("notifyDeposit ErrNeedRetry:%v, timestamp:%v", err, timestamp)

		if err == common.ErrRepeatDeposit {
			return common.ErrRepeatDeposit
		}

		return common.ErrNeedRetry
	}

	err = common.FuncWithRetry(5, func() error {
		return notifyConfirm(&wallet.NotifyConfirmRequest{
			TxHash:  deposit.TxHash,
			Confirm: 100,
		})
	})
	if err != nil {
		log.Errorf("notifyConfirm ErrNeedRetry:%v, timestamp:%v", err, timestamp)

		log.Error(err)
		return common.ErrNeedRetry
	}

	log.Info("NotifyDeposit complete")
	return nil
}

func notifyDeposit(req *wallet.NotifyDepositRequest) error {
	log.Infof("notifyDeposit req; %v", *req)

	if walletclient.State() != connectivity.Ready {
		time.Sleep(1 * time.Second)
		return errors.New("wallet server not ready")
	}
	// do notify
	reply, err := walletClient.NotifyDeposit(context.Background(), req)
	if err != nil {
		log.Warn("Wallet service not avaliable")
		return err
	}

	if reply == nil {
		msg := fmt.Sprintf("wallet reply nil, tx hash:%v", req.TxHash)
		log.Error(msg)
		return errors.New(msg)
	}

	if reply.Code != pbcommon.ReturnCode_SUCCESS {
		msg := fmt.Sprintf("NotifyDeposit tx %v error :%v", req.TxHash, reply.Msg)
		log.Error(msg)

		if reply.Msg == "repeat deposit" {
			return common.ErrRepeatDeposit
		}

		if strings.Contains(reply.Msg, "wrong api key") {
			return common.ErrInvalidAPIKey
		}

		return errors.New(msg)
	}

	log.Infof("Notify deposit:%v success.", req.TxHash)
	return nil
}

func notifyConfirm(req *wallet.NotifyConfirmRequest) error {
	if walletclient.State() != connectivity.Ready {
		log.Warn("Wallet service not avaliable")
		return errors.New("wallet server not ready")
	}
	// do notify
	reply, err := walletClient.NotifyConfirm(context.Background(), req)
	if err != nil {
		log.Errorf("Wallet notify tx %v error:%v", req.TxHash, err)
		return err
	}

	if reply == nil {
		log.Errorf("wallet notify tx %v, return nil", req.TxHash)
		return errors.New("wallet notify response nil")
	}

	if reply.Code != pbcommon.ReturnCode_SUCCESS {
		msg := fmt.Sprintf("NotifyConfirm err:%v", reply.Msg)
		log.Error(msg)
		return errors.New(msg)
	}

	log.Infof("Notify confirm:%v success.", req.TxHash)
	return nil
}

type tokenStatus struct {
	CanDeposit    bool `json:"can_deposit"`
	CanWithdrawal bool `json:"can_withdrawal"`
}

func queryTokenStatus(tokenID string) (*tokenStatus, error) {
	// query wallet db
	reqData := map[string]interface{}{
		"token_id": tokenID,
	}

	data, err := json.Marshal(reqData)
	if err != nil {
		log.Error("queryTokenStatus json marshal err:%v", err)
		return nil, common.ErrInvalidParams
	}

	reply := &wallet.QueryWalletDBReply{}
	err = common.FuncWithRetry(5, func() error {
		reply, err = walletClient.QueryWalletDB(context.Background(), &wallet.QueryWalletDBRequest{
			Method:  "BaasGetTokenStatus",
			ReqData: data,
		})

		return err
	})
	if err != nil {
		log.Error("queryTokenStatus ErrNeedRetry query wallet db err:%v", err)
		return nil, common.ErrNeedRetry
	}

	if reply.Code != pbcommon.ReturnCode_SUCCESS {
		log.Error("queryTokenStatus query wallet db err:%v", reply.Msg)
		return nil, common.ErrInvalidParams
	}
	log.Infof("queryTokenStatus token_id:%v, status:%v", tokenID, string(reply.GetResData()))

	var result tokenStatus
	json.Unmarshal(reply.ResData, &result)

	return &result, nil
}
