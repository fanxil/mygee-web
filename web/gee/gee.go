package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
// HandlerFunc 定义了 gee 使用的请求处理程序
// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
//type Engine struct {
//	//key : GET-/、GET-/hello、POST-/hello
//	router map[string]HandlerFunc
//}

type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
//func New() *Engine {
//	return &Engine{router: make(map[string]HandlerFunc)}
//}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//key := method + "-" + pattern
	//log.Printf("Route %4s - %s", method, pattern)
	//engine.router[key] = handler
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

//func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	key := req.Method + "-" + req.URL.Path
//	if handler, ok := engine.router[key]; ok {
//		handler(w, req)
//	} else {
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
//	}
//}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
