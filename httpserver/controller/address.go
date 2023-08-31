/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: address.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"

	"github.com/gin-gonic/gin"
)

// MaxUnusedAddressLength define
const MaxUnusedAddressLength = 10000

// UnusedAddressCount return the unused address count
func UnusedAddressCount(c *gin.Context) {
	chain := c.Query("chain")

	// chain chain
	tokenInfo := common.GetTokenInfoByChain(chain)
	if tokenInfo == nil {
		c.JSON(http.StatusOK, Result{
			Code: INVALID_CHAIN,
			Msg:  "invali chain",
		})
		return
	}

	apiKey := c.Request.Header.Get("BWAAS-API-KEY")
	if tokenInfo.APIKey != apiKey {
		c.JSON(http.StatusOK, Result{
			Code: INVALID_APIKEY,
			Msg:  "invali apikey",
		})
		return
	}

	listLen := common.Redis.Client.SCard(fmt.Sprintf(common.UnusedAddressSet, chain)).Val()
	log.Infof("UnusedAddressCount request_id:%v, address count:%v", getRequestID(c), listLen)

	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
		Data: listLen,
	})
}

type addAddressParams struct {
	Chain    string   `json:"chain"`
	AddrList []string `json:"addr_list"`
}

// AddAddress add new address to cache
func AddAddress(c *gin.Context) {
	requestID := getRequestID(c)

	var para addAddressParams
	if err := c.ShouldBind(&para); err != nil {
		log.Errorf("AddAddress parse params request_id:%v,err:%v", requestID, err)

		c.JSON(http.StatusOK, Result{
			Code: INVALID_PARAMS,
			Msg:  "invalid params",
		})
		return
	}
	log.Infof("AddAddress request_id:%v, para: %v", requestID, para)

	// check chain
	tokenInfo := common.GetTokenInfoByChain(para.Chain)
	if tokenInfo == nil {
		c.JSON(http.StatusOK, Result{
			Code: INVALID_CHAIN,
			Msg:  "invali chain",
		})
		return
	}

	apiKey := c.Request.Header.Get("BWAAS-API-KEY")
	if tokenInfo.APIKey != apiKey {
		c.JSON(http.StatusOK, Result{
			Code: INVALID_APIKEY,
			Msg:  "invali apikey",
		})
		return
	}

	unUsedAddressKey := fmt.Sprintf(common.UnusedAddressSet, para.Chain)
	log.Infof("unused AddAddress request_id:%v, address key: %v", requestID, unUsedAddressKey)

	// check length
	listLen := common.Redis.Client.SCard(unUsedAddressKey).Val()
	if listLen >= MaxUnusedAddressLength {
		c.JSON(http.StatusOK, Result{
			Code: ADDRESS_ENOUGH,
			Msg:  "address enough",
		})
		return
	}

	// distinct addresses
	distinctedAddresses := make(map[string]struct{})
	for _, addr := range para.AddrList {
		if !checkAddressFormat(tokenInfo.AddressRegexp, addr) {
			log.Infof("address:%v, check regexp %vformat err", addr, tokenInfo.AddressRegexp)
			c.JSON(http.StatusOK, Result{
				Code: INVALID_TO_ADDRESS,
				Msg:  "invalid address",
			})
			return
		}

		distinctedAddresses[addr] = struct{}{}
	}

	usedAddressKey := fmt.Sprintf(common.UsedAddressSet, para.Chain)
	log.Infof("used AddAddress request_id:%v, address key: %v", requestID, usedAddressKey)

	for addr := range distinctedAddresses {
		// check adderss exist
		lowerAddr := strings.ToLower(addr)
		existL, _ := common.Redis.IsSetMember(usedAddressKey, lowerAddr)

		existU, _ := common.Redis.IsSetMember(usedAddressKey, addr)
		if !existL && !existU {
			common.Redis.Client.SAdd(unUsedAddressKey, addr)
			listLen++
		}

		if listLen >= MaxUnusedAddressLength {
			break
		}
	}

	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
	})
}

func checkAddressFormat(reg, address string) bool {
	if reg == "" {
		return true
	}

	// check multi regs
	if strings.Contains(reg, ",") {
		subRegs := strings.Split(reg, ",")
		for _, sr := range subRegs {
			addressMatched, _ := regexp.MatchString(sr, address)
			if addressMatched {
				return true
			}
		}
	}

	addressMatched, _ := regexp.MatchString(reg, address)
	return addressMatched
}
