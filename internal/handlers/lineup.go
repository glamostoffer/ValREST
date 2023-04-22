package handlers

import (
	"ValREST/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createLinup(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.Lineup
	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Lineup.Create(userId, input)
	if err != nil {
		logrus.Errorf("error during creating lineup service: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLineups(c *gin.Context) {

}

func (h *Handler) getLinupByAgent(c *gin.Context) {
	sources, err := h.services.Lineup.GetByAgent(c.Param("agent"))
	if err != nil {
		logrus.Errorf("error during getting lineups sources: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sources)
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
