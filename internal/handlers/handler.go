package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/register")
	}

	api := router.Group("/api")
	{
		api.GET("/lineup")
		api.POST("/lineup/:agent/:map/:objective/:ability")
		api.PUT("/lineup/:agent/:map/:objective/:ability")
		api.GET("/lineup/:agent")
		api.GET("/lineup/:agent/:map")
		api.GET("/lineup/:agent/:map/:objective")
		api.GET("/lineup/:agent/:map/:objective/:ability")
	}

	return router
}
