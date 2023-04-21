package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logrus.Error("empty auth header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		logrus.Error("invalid auth header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		logrus.Errorf("user id not found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return 0, errors.New("user id not found")
	}

	intId, ok := id.(int)
	if !ok {
		logrus.Errorf("can't convert user id into int type")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "can't convert user id into int type"})
		return 0, errors.New("can't convert user id into int type")
	}

	return intId, nil
}
