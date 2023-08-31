/*
 * *******************************************************************
 * @项目名称: baasnode
 * @文件名称: httpserver.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package httpserver

import (
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/config"
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/httpserver/router"
)

// Run a http server
func Run() {
	engine := router.Router
	engine.Run(":" + config.Config.Chainnode.HTTPPort)
}
