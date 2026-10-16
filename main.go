package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gohari-6/database"
	"gohari-6/todo"
)

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func GetName(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}

func GetHelloName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{"message": "Hello, " + name + "!"})
}

func main() {
	r := gin.Default()

	// Koneksi ke database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	defer db.Close()

	// Inisialisasi handler
	handler := todo.TodoHandler{DB: db}

	// Route untuk todo (prefix /api/v1)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", handler.List)         // Daftar todo
		v1.GET("/todos/:id", handler.GetByID)  // Detail todo
		v1.POST("/todos", handler.Create)      // Tambah todo baru
		v1.PUT("/todos/:id", handler.Update)   // Ubah todo
		v1.DELETE("/todos/:id", handler.Delete)// Hapus todo
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
