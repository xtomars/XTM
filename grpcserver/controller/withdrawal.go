/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: withdrawal.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	baascommon "git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	"git.bhex.io/bhpc/wallet/chainnode/util/walletclient"
	"git.bhex.io/bhpc/wallet/common/protos/common"
	"git.bhex.io/bhpc/wallet/common/protos/risk"
	"git.bhex.io/bhpc/wallet/common/protos/wallet"
	"github.com/golang/protobuf/proto"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/connectivity"
)

var batchCount = 50

func getWithdrawlOrdersFromWallet(chain, apiKey string) ([]*common.BaasWithdrawalOrder, error) {
	// get wallet state
	if walletclient.State() != connectivity.Ready {
		log.Errorf("wallet not ready")
		return nil, errors.New("wallet not ready")
	}
	var req wallet.GetWithdrawalOrdersRequest
	req.ApiKey = apiKey
	req.Chain = chain
	log.Infof("getWithdrawlOrdersFromWallet req:%v ", req)
	// do notify
	reply, err := walletClient.GetWithdrawalOrders(context.Background(), &req)
	if err != nil {
		msg := fmt.Sprintf("GetWithdrawalOrders  chain %v, apikey %v error:%v", chain, apiKey, err)
		log.Errorf(msg)
		return nil, err
	}
	if reply == nil {
		msg := fmt.Sprintf("GetWithdrawalOrders  chain %v, apikey %v reply is nil", chain, apiKey)
		log.Error(msg)
		return nil, errors.New(msg)
	}

	if reply.Code != common.ReturnCode_SUCCESS {
		msg := fmt.Sprintf("GetWithdrawalOrders  chain %v, apikey %v  err:%v", chain, apiKey, reply.Msg)
		log.Error(msg)
		return nil, errors.New(msg)
	}

	log.Infof("GetWithdrawalOrders  chain %v, apikey %v success, orders:%v", chain, apiKey, reply.Orders)
	return reply.Orders, nil
}

func getWithdrawalRcSign(orders []*common.BaasWithdrawalOrder, method, url, timestamp string) ([]*risk.BaasWithdrawalRcSign, error) {
	var rcReq risk.BaasWithdrawalWalletRcRequest
	rcReq.Method = method
	rcReq.Url = url
	rcReq.Timestamp = timestamp
	rcReq.Orders = orders

	reply, err := riskClient.BaasWithdrawalWalletRc(context.Background(), &rcReq)
	if err != nil {
		log.Errorf("risk service not avaliable")
		return nil, err
	}

	if reply == nil {
		msg := fmt.Sprintf("BaasWithdrawalWalletRc reply is nil")
		log.Error(msg)
		return nil, errors.New(msg)
	}

	if reply.Code != common.ReturnCode_SUCCESS {
		msg := fmt.Sprintf("BaasWithdrawalWalletRc  orders %v, err %v", orders, reply.Msg)
		log.Error(msg)
		return nil, errors.New(msg)
	}

	return reply.RcSigns, nil
}

// GetWithdrawalOrders return the unhandled orders
func GetWithdrawalOrders(chain, apikey string) []baascommon.WithdrawalOrder {
	// query from wallet and get sign from  rc
	orders, err := getWithdrawlOrdersFromWallet(chain, apikey)
	if err != nil {
		log.Errorf("get order from wallet error %v", err)
		return nil
	}

	method := "GET"
	url := "/api/v1/withdrawal/orders"
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1000)

	rcSigns, err := getWithdrawalRcSign(orders, method, url, timestamp)
	if err != nil {
		log.Errorf("getWithdrawalRcSign from rc error %v", err)
		return nil
	}

	var retOrders []baascommon.WithdrawalOrder
	for i, order := range orders {
		var retOrder baascommon.WithdrawalOrder
		retOrder.OrderID = order.OrderId
		retOrder.TokenID = order.TokenId
		retOrder.To = order.To
		retOrder.Memo = order.Memo
		retOrder.Amount = order.Amount
		retOrder.TimeStamp = timestamp
		retOrder.Signature = rcSigns[i].RcSign
		retOrders = append(retOrders, retOrder)
	}

	log.Infof("GetWithdrawalOrders success chain %v, orders:%v", chain, retOrders)
	return retOrders
}

// NotifyFailed notify withdrawal failed
func NotifyFailed(order baascommon.NotifyFailedParams) error {
	if order.OrderID == "" {
		return baascommon.ErrInvalidOrderID
	}

	status, err := queryTokenStatus(order.TokenID)
	if err != nil {
		log.Errorf("NotifyDeposit get token status failed:%v", err)
		return err
	}

	if !status.CanWithdrawal {
		return baascommon.ErrPauseWithdrawal
	}

	tokenInfo := baascommon.GetTokenInfoByTokenID(order.TokenID)
	if tokenInfo == nil {
		return baascommon.ErrInvalidTokenID
	}
	log.Infof("NotifyFailed token:%v", tokenInfo)

	err = notifyWithdrawal(&wallet.NotifyWithdrawalRequest{
		BusinessNumber: order.OrderID,
		TokenId:        order.TokenID,
		TxHash:         "failed order",
		Extension: map[string]string{
			"success": "false",
			"reason":  order.Reason,
		},
	})
	if err != nil {
		msg := fmt.Sprintf("[ALERT] NotifyFailed failed, business number:%v，err:%v", order.OrderID, err)
		log.Error(msg)
		return baascommon.ErrInvalidOrderID
	}

	return nil
}

// NotifyWithdrawal notify withdrawal finished
func NotifyWithdrawal(order baascommon.NotifyWithdrawalParams) error {
	if order.OrderID == "" {
		return baascommon.ErrInvalidOrderID
	}

	status, err := queryTokenStatus(order.TokenID)
	if err != nil {
		log.Errorf("NotifyDeposit get token status failed:%v", err)
		return err
	}

	if !status.CanWithdrawal {
		return baascommon.ErrPauseWithdrawal
	}

	tokenInfo := baascommon.GetTokenInfoByTokenID(order.TokenID)
	if tokenInfo == nil {
		return baascommon.ErrInvalidTokenID
	}

	// check amount
	amount, _ := decimal.NewFromString(order.Amount)
	if amount.Cmp(decimal.Zero) == 0 || amount.Cmp(decimal.Zero) < 0 {
		return baascommon.ErrInvalidAmount
	}
	if amount.String() != order.Amount {
		log.Errorf("NotifyWithdrawal ErrInvalidAmount, amount:%v, order amount:%v", amount.String(), order.Amount)
		return baascommon.ErrInvalidAmount
	}

	// check amount decimal
	ext := int(amount.Exponent())
	log.Infof("NotifyWithdrawal deposit amount decimal:%v", ext)

	if tokenInfo.Decimals < ext {
		return baascommon.ErrInvalidDecimals
	}

	blockTime, _ := decimal.NewFromString(order.BlockTime)
	if blockTime.Cmp(decimal.Zero) == 0 || blockTime.Cmp(decimal.Zero) < 0 {
		return baascommon.ErrinvalidBlockTime
	}

	blockHeight, _ := decimal.NewFromString(order.BlockHeight)
	if blockHeight.Cmp(decimal.Zero) == 0 || blockHeight.Cmp(decimal.Zero) < 0 {
		return baascommon.ErrInvalidBlockHeight
	}

	if order.TxHash == "" {
		return baascommon.ErrInvalidTxHash
	}

	err = baascommon.FuncWithRetry(5, func() error {
		return notifyWithdrawal(&wallet.NotifyWithdrawalRequest{
			TokenId:        order.TokenID,
			BusinessNumber: order.OrderID,
			TxHash:         order.TxHash,
			BlockHeight:    uint64(blockHeight.IntPart()),
			BlockTime:      uint64(blockTime.IntPart()),
			Extension: map[string]string{
				"success": "true",
			},
		})
	})
	if err != nil {
		msg := fmt.Sprintf("[ALERT] NotifyWithdrawal failed, hash:%v, business number:%v", order.TxHash, order.OrderID)
		log.Error(msg)
		return baascommon.ErrNeedRetry
	}

	err = baascommon.FuncWithRetry(5, func() error {
		return notifyConfirm(&wallet.NotifyConfirmRequest{
			TxHash:  order.TxHash,
			Confirm: 100,
		})
	})
	if err != nil {
		log.Error(err)
		return baascommon.ErrNeedRetry
	}

	return nil
}

// withdrawal commit
// func withdrawal(tokenID string, businessNumber string, req *common.SendMWTransactionRequest) error {
// 	log.Infof("withdrawal start, token_id %v, bussiness number %v", tokenID, businessNumber)

// 	// check in rc
// 	rcSign, err := getRcSign(req)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	tokenInfo := baascommon.GetTokenInfoByTokenID(req.TokenId)
// 	if tokenInfo == nil {
// 		log.Errorf("[ALERT] get tokeninfo err, token_id:%v", req.TokenId)
// 		return errors.New("invalid token id")
// 	}

// 	// add withdrawal order to list
// 	if _, err := baascommon.Redis.Client.RPush(fmt.Sprintf(baascommon.WithdrawalOrderList, tokenInfo.Chain), baascommon.WithdrawalOrder{
// 		OrderID:   req.BusinessNumber,
// 		TokenID:   req.TokenId,
// 		To:        req.To,
// 		Memo:      req.Memo,
// 		Amount:    req.Amount,
// 		Signature: rcSign,
// 	}).Result(); err != nil {
// 		log.Errorf("[ALERT] Save withdrawal order to redis err:%v", err)
// 		return err
// 	}

// 	return nil
// }

func getRcSign(req *common.SendOnlineWalletTransactionRequest) (string, error) {
	var rcReq risk.SignTxRcRequest
	log.Info("verifyWithdrawRc req: ", *req)

	txHashRc, err := getTxHashRc(req)
	if err != nil {
		errInfo := fmt.Sprintf("getTxHashRc req %v ,err %v", *req, err)
		log.Error(errInfo)
		return "", err
	}

	// do check
	rcReq.TxHash = txHashRc
	reply, err := riskClient.SignTxRc(context.Background(), &rcReq)
	if err != nil {
		log.Warn("risk service not avaliable")
		return "", err
	}

	if reply == nil {
		errInfo := fmt.Sprintf("risk reply nil, tx hash:%v  ,err %v", rcReq.TxHash, err)
		log.Error(errInfo)
		return "", err
	}

	if reply.Code != 0 {
		errInfo := fmt.Sprintf("verifyWithdrawRc tx %v error :%v", rcReq.TxHash, reply.Msg)
		log.Error(errInfo)
		return "", errors.New(errInfo)
	}

	return reply.Extension["signature"], nil
}

func getTxHashRc(req proto.Message) (string, error) {
	b, err := proto.Marshal(req)
	if err != nil {
		log.Errorf("getTxHashRc marshal reqest msg error, err:%v, msg:%v", err, req.String())
		return "", err
	}

	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:]), nil
}

func notifyWithdrawal(req *wallet.NotifyWithdrawalRequest) error {
	// get wallet state
	if walletclient.State() != connectivity.Ready {
		log.Errorf("wallet not ready")
		return errors.New("wallet not ready")
	}

	log.Infof("Notify Withdrawal req:%v ", req)
	// do notify
	reply, err := walletClient.NotifyWithdrawal(context.Background(), req)
	if err != nil {
		msg := fmt.Sprintf("Wallet reply tx %v error:%v", req.TxHash, err)
		log.Errorf(msg)
		return err
	}
	if reply == nil {
		msg := fmt.Sprintf("wallet reply nil, tx:%v", req.TxHash)
		log.Error(msg)
		return errors.New(msg)
	}

	if reply.Code != common.ReturnCode_SUCCESS {
		msg := fmt.Sprintf("NotifyWithdrawal tx %v err:%v", req.TxHash, reply.Msg)
		log.Error(msg)
		return errors.New(msg)
	}

	log.Infof("Notify Withdrawal:%v success.", req.TxHash)
	return nil
}
