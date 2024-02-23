package server

import (
	"net/http"
)

func StartServer() {
	// Gin 라우터 초기화
	router := gin.Default()

	// 정적 파일 서빙
	router.Static("/static", "./static")

	// 라우트 및 핸들러 정의
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 웹 서버 시작
	router.Run(":8080")
}
