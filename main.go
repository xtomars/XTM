/*
 * *******************************************************************
 * @项目名称: baasnode
 * @文件名称: main.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package main

import (
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/grpcserver"
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/httpserver"
)

func main() {
	go common.SyncAllTokens()

	// run http server
	go httpserver.Run()

	// run grpc server
	grpcserver.Run()

}
