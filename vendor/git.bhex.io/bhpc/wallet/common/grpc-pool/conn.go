/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: conn.go
 * @Date: 2020/05/14
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package pool

import (
	"google.golang.org/grpc"
)

// Conn single grpc connection inerface
type Conn interface {
	// Value return the actual grpc connection type *grpc.ClientConn.
	Value() *grpc.ClientConn

	// Close decrease the reference of grpc connection, instead of close it.
	// if the pool is full, just close it.
	Close() error
}

// Conn is wrapped grpc.ClientConn. to provide close and value method.
type conn struct {
	cc   *grpc.ClientConn
	pool *pool
	once bool
}

// Value see Conn interface.
func (c *conn) Value() *grpc.ClientConn {
	return c.cc
}

// Close see Conn interface.
func (c *conn) Close() error {
	c.pool.decrRef()
	if c.once {
		return c.reset()
	}
	return nil
}

func (c *conn) reset() error {
	cc := c.cc
	c.cc = nil
	c.pool = nil
	c.once = false
	if cc != nil {
		return cc.Close()
	}
	return nil
}

func (p *pool) wrapConn(cc *grpc.ClientConn, once bool) *conn {
	return &conn{
		cc:   cc,
		pool: p,
		once: once,
	}
}
