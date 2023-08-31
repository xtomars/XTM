/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: init.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/config"
	"git.bhex.io/bhpc/wallet/chainnode/util/walletclient"
	"git.bhex.io/bhpc/wallet/common/protos/risk"
	"git.bhex.io/bhpc/wallet/common/protos/wallet"

	"google.golang.org/grpc"
)

var (
	walletClient wallet.WalletClient
	riskClient   risk.RiskClient
)

func init() {
	var err error

	// wallet client
	walletClient, err = walletclient.New(config.Config.Wallet.Host, config.Config.Wallet.Port)
	if err != nil {
		log.Error("Connect wallet server err:", err)
		panic(err)
	}

	riskClient, err = newRiskClient(config.Config.Risk.Host, config.Config.Risk.Port)
	if err != nil {
		log.Error("[ALERT] Connect risk server err:", err)
		panic(err)
	}
}

func newRiskClient(host, port string) (risk.RiskClient, error) {
	var dialOpts = []grpc.DialOption{
		grpc.WithInsecure(),
		// grpc.WithTimeout(5 * time.Second),
		// grpc.WithBlock(),
	}

	riskConn, err := grpc.Dial(host+":"+port, dialOpts...)
	if err != nil {
		log.Errorf("Could not connect: %v", err)
		return nil, err
	}
	log.Info("Connect risk server success!")

	return risk.NewRiskClient(riskConn), nil
}
