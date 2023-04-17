package handlers

import (
	"ValREST/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) register(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("error binding json: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("error binding json: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Name, input.Password)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
