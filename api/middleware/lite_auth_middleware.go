package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// LiteAuth 미들웨어 설정
func LiteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 헤더에서 Authorization 값 가져오기
		token := c.GetHeader("Authorization")

		// Authorization 헤더가 없는 경우
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Bearer 토큰인지 확인
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// 실제 토큰 값 추출
		token = strings.TrimPrefix(token, "Bearer ")

		// 토큰 검증 로직 (여기서는 예제이므로 간단히 길이 체크만 함)
		if len(token) != 32 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 인증이 통과된 경우 다음 핸들러로 요청 전달
		c.Next()
	}
}
