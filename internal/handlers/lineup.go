package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createLinup(c *gin.Context) {

}

func (h *Handler) getAllLineups(c *gin.Context) {

}

func (h *Handler) getLinupByAgent(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"agent": c.Param("agent"),
	})
}

func (h *Handler) getLinupByAgentAndMap(c *gin.Context) {

}

func (h *Handler) getLinupWithAgentMapObjective(c *gin.Context) {

}

func (h *Handler) getLinupWithAgentMapObjectiveAbility(c *gin.Context) {

}

func (h *Handler) updateLinup(c *gin.Context) {

}

func (h *Handler) deleteLinup(c *gin.Context) {

}
