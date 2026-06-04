package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"baibaoxiang/config"
	"baibaoxiang/database"
	"baibaoxiang/handlers"
	"baibaoxiang/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	// Initialize database
	if err := database.Init(cfg.DBPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Setup Gin
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// Health check
	r.GET("/api/ping", handlers.Ping)

	// Tool API routes (to be implemented in later days)
	// api := r.Group("/api/tools")
	// {
	// 	api.POST("/json-format", handlers.JsonFormat)
	// 	api.POST("/base64", handlers.Base64Convert)
	// 	api.POST("/timestamp", handlers.TimestampConvert)
	// 	api.POST("/qrcode", handlers.QRCodeGenerate)
	// 	api.POST("/color", handlers.ColorConvert)
	// 	api.POST("/text", handlers.TextProcess)
	// 	api.POST("/uuid", handlers.UUIDGenerate)
	// }

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down gracefully...")
		database.Close()
		os.Exit(0)
	}()

	addr := ":" + cfg.Port
	log.Printf("百宝箱 backend starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
