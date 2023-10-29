### 服务启动
1. 实现Run方法， 通过调用net/http包的http.ListenAndServe(address, engine.Handler())启动，因为ListenAndServe第二个参数需要传一个http.Handler 的接口，所以engine需要实现这个接口的方法ServeHTTP(ResponseWriter, *Request) 
2. 通过ServeHTTP方法， 调用路由对应的处理方法， 然后返回响应

### 路由
1. 定义路由处理器interface, 主要需要实现两个方法
- 路由添加
- 中间件添加