package handlers

import (
	"ValREST/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/register", h.register)
	}

	api := router.Group("/lineup", h.userIdentity)
	{
		api.GET("/", h.getAllLineups)

		api.POST("/:agent/:map/:objective/:ability", h.createLinup)
		api.PUT("/:agent/:map/:objective/:ability", h.updateLinup)
		api.DELETE("/:agent/:map/:objective/:ability", h.deleteLinup)
		api.GET("/:agent", h.getLinupByAgent)
		api.GET("/:agent/:map", h.getLinupByAgentAndMap)
		api.GET("/:agent/:map/:objective", h.getLinupWithAgentMapObjective)
		api.GET("/:agent/:map/:objective/:ability", h.getLinupWithAgentMapObjectiveAbility)
	}

	return router
}
