package middleware

import (
	"github.com/go-chi/cors"
)

// Cors creates a new CORS middleware handler
func Cors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // ここにフロントエンドのオリジンを設定
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5分
	})
}
