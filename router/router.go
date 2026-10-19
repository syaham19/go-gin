package router

import (
	"gohari-6/todo" // Pastikan nama modul sesuai dengan file go.mod Anda

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter membuat router Gin lengkap dengan semua route.
// Dipakai oleh main.go dan oleh test.
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Membuat grup route dengan prefix /api/v1 (misal: /api/v1/todos)
	v1 := r.Group("/api/v1")
	{
		// Inisialisasi struct TodoHandler dan memasukkan koneksi database GORM
		handler := &todo.TodoHandler{DB: db}

		// Memanggil fungsi RegisterRoutes yang ada di dalam package todo
		todo.RegisterRoutes(v1, handler)
	}

	return r
}
