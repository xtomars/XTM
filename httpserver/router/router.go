/*
 * *******************************************************************
 * @项目名称: router
 * @文件名称: router.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/httpserver/controller"
	"git.bhex.io/bhpc/wallet/common/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	// Router is the global router
	Router *gin.Engine

	log = logger.New("router")

	nextRequestID = func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	Router = gin.New()
	Router.Use(gin.Recovery())

	// support cors request
	Router.Use(corsMiddleware())

	// health check
	Router.GET("/internal/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// api version control
	v1 := Router.Group("/api/v1")
	{
		v1.Use(checkSign())

		v1.GET("/address/unused/count", controller.UnusedAddressCount)
		v1.POST("/address/add", controller.AddAddress)
		v1.POST("/notify/deposit", controller.NotifyDeposit)
		v1.POST("/notify/withdrawal", controller.NotifyWithdrawal)
		v1.POST("/notify/failed", controller.NotifyFailed)
		v1.GET("/withdrawal/orders", controller.GetWithdrawalOrders)
		v1.POST("/asset/verify", controller.AssetVerify)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "origin", "X-Requested-With", "BWAAS-API-KEY", "BWAAS-API-SIGNATURE", "BWAAS-API-TIMESTAMP"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})
}

func checkSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set trace id
		requestID := c.Request.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = nextRequestID()
		}
		c.Request.Header.Set("X-Request-Id", requestID)

		apiKey := c.Request.Header.Get("BWAAS-API-KEY")
		signature := c.Request.Header.Get("BWAAS-API-SIGNATURE")
		timestamp := c.Request.Header.Get("BWAAS-API-TIMESTAMP")

		log.Infof("checkSign request_id:%v, api_key:%v,signature:%v,timestamp:%v", requestID, apiKey, signature, timestamp)

		// check timestamp
		now := time.Now().UnixNano() / 1e6
		inTime, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			log.Infof("checkSign failed, invalid timestamp,request_id:%v, inTime: %v", requestID, timestamp)

			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 10006,
				"msg":  "invalid timestamp",
			})
			return
		}

		if now-inTime > 120000 {
			log.Infof("checkSign failed, timestamp expired, request_id:%v, now:%v, in:%v", requestID, now, inTime)

			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 10020,
				"msg":  "timestamp expired",
			})
			return
		}

		// check apikey
		if !common.ValidAPIKey(apiKey) {
			log.Infof("checkSign failed,request_id:%v, invalid apikey", requestID)

			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 10002,
				"msg":  "invalid apikey",
			})
			return
		}

		// check signature
		var paramsMsg string
		var mapBody gin.H

		if c.Request.Method == "POST" {
			body, _ := ioutil.ReadAll(c.Request.Body)
			json.Unmarshal(body, &mapBody)
			log.Infof("checkSign request_id:%v, reqeustBody: %v", requestID, mapBody)

			// write back to body
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		log.Infof("CreateSignMsg request_id:%v, method:%v ,url:%v,timestamp:%v, body:%v", requestID, c.Request.Method, c.Request.URL.String(), timestamp, mapBody)
		paramsMsg = common.CreateSignMsg(c.Request.Method, c.Request.URL.String(), timestamp, mapBody)
		log.Infof("checkSign request_id:%v, msg:%v", requestID, paramsMsg)

		if !common.VerifySign(apiKey, paramsMsg, signature) {
			log.Infof("checkSign failed request_id:%v", requestID)

			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "invalid signature",
			})
			return
		}

		c.Next()
	}
}
