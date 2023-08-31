/*
 * *******************************************************************
 * @项目名称: walletclient
 * @文件名称: wallet.go
 * @Date: 2018/05/16
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package walletclient

import (
	pool "git.bhex.io/bhpc/wallet/common/grpc-pool"
	"git.bhex.io/bhpc/wallet/common/logger"
	"git.bhex.io/bhpc/wallet/common/protos/wallet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	conn *grpc.ClientConn
	log  = logger.New("walletclient")
)

// New new a wallet client
func New(host, port string) (wallet.WalletClient, error) {
	var err error
	conn, err = pool.Dial(host + ":" + port)
	if err != nil {
		log.Errorf("Could not connect: %v", err)
		return nil, err
	}
	log.Info("Connect wallet server success!")

	return wallet.NewWalletClient(conn), nil
}

// State return the grpc conn state
func State() connectivity.State {
	return conn.GetState()
}

// Close close wallet client
func Close() {
	conn.Close()
}
