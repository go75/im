package imerr

import "errors"

// token
var ExpiredTokenErr = errors.New("token过期")
var InvaildTokenErr = errors.New("token无效")

// websocket
var AlreadyExistConnErr = errors.New("websocket连接已存在")
var NotExistConnErr = errors.New("连接不存在")