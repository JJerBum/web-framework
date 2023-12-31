package main

import (
	"log"
	"time"
)

// Middleware는 handler를 처리하기 전에 거치는 함수 단위를 정의하는 타입 입니다.
type Middleware func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		// next(c)를 실행하기 전에 현재 시간을 기록
		t := time.Now()

		// 다음 핸들러 수행
		next(c)

		// 웹 요청 정보와 전체 소요 시간을 로그로 남김
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	}
}
