package web_framework

import "net/http"

type Engine struct {
}

func (engine *Engine) ServeHTTP(http.ResponseWriter, *http.Request) {}

func (engine *Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, engine)
	return err
}
