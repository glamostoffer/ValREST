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
		api.GET("/lineup", h.getAllLineups)
		api.POST("/lineup/:agent/:map/:objective/:ability", h.createLinup)
		api.PUT("/lineup/:agent/:map/:objective/:ability", h.updateLinup)
		api.DELETE("/lineup/:agent/:map/:objective/:ability", h.deleteLinup)
		api.GET("/lineup/:agent", h.getLinupByAgent)
		api.GET("/lineup/:agent/:map", h.getLinupByAgentAndMap)
		api.GET("/lineup/:agent/:map/:objective", h.getLinupWithAgentMapObjective)
		api.GET("/lineup/:agent/:map/:objective/:ability", h.getLinupWithAgentMapObjectiveAbility)
	}

	return router
}
