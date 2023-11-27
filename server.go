package main

import (
	"log"
	"net/http"
)

// Server는 타입으로 사용하여 전체 사용을 추상화 시킵니다.
type Server struct {
	*router
	middlewares  []Middleware
	startHandler HandlerFunc
}

func NewServer() *Server {
	r := &router{make(map[string]map[string]HandlerFunc)}
	s := &Server{router: r}

	s.middlewares = []Middleware{
		logHandler,
	}

	return s
}

func recoverHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next(c)
	}
}

func (s *Server) Run(addr string) error {
	// startHandler를 라우터 핸들러 함수로 지정
	s.startHandler = s.router.handler()

	// 등록된 미들웨어를 라우터 핸들러 앞에 하나씩 추가
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		s.startHandler = s.middlewares[i](s.startHandler)
	}

	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Context 생성
	c := &Context{
		Params:         make(map[string]any),
		ResponseWriter: w,
		Request:        r,
	}

	for k, v := range r.URL.Query() {
		c.Params[k] = v[0]
	}

	s.startHandler(c)
}

func (s *Server) Use(middlewares ...Middleware) {
	s.middlewares = append(s.middlewares, middlewares...)
}
