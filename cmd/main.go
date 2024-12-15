package main

import (
	_ "ai-thing/docs"
	"ai-thing/handler"
	"ai-thing/router"
	_ "ai-thing/schemas"
	"ai-thing/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	h := handler.NewHandler()
	e := gin.Default()

	if !(os.Getenv("GIN_MODE") == "release") || os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
		h.Logger.Info("Running in debug mode")
		if err := utils.LoadEnv(); err != nil {
			h.Logger.Error("Failed to load env", "error", err.Error())
			os.Exit(1)
		}
	}

	router.InitRouter(e, h)

	if err := e.Run(":" + os.Getenv("PORT")); err != nil {
		h.Logger.Error("Failed to start server", "error", err.Error())
		os.Exit(1)
	}
}
