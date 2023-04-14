package handlers

import (
	"ValREST/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) register(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("error binding json: %s", err.Error())
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

}

func (h *Handler) login(c *gin.Context) {

}
