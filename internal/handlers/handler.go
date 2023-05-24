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

	api := router.Group("/api")
	{
		api.GET("/", h.getAllLineups)

		api.POST("/", h.userIdentity, h.createLinup)
		api.PUT("/", h.userIdentity, h.updateLinup)
		api.DELETE("/:id", h.userIdentity, h.deleteLinup)
		api.GET("/:agent", h.getLinupByAgent)
		api.GET("/:agent/:map", h.getLinupByAgentAndMap)
		api.GET("/:agent/:map/:objective", h.getLinupWithAgentMapObjective)
		api.GET("/:agent/:map/:objective/:ability", h.getLinupWithAgentMapObjectiveAbility)
	}

	return router
}
