package web_framework

type HandlerFunc func(*Context)

// IRouters 定义全部的路由处理器接口
type IRouters interface {
	AddRouter(string, ...HandlerFunc) IRouters
	AddMiddleWare(...HandlerFunc) IRouters
}
