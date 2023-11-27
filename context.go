package main

import "net/http"

// HandlerFunc는 Handler(처리자)의 함수 원형을 나타내는 타입 입니다.
type HandlerFunc func(c *Context)

// Context는 웹 요청의 처리 상태를 저장하는 구조체 입니다.
// 예를 들면 Client's Request와 Client's Params가 있습니다.
type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}
