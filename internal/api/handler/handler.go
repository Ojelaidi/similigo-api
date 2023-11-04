package handler

import (
	"github.com/Ojelaidi/similigo-api/internal/api/similigo-api"
	"github.com/Ojelaidi/similigo-api/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	similiService *similigo_api.Service
}

func NewHandler(similiService *similigo_api.Service) *Handler {
	return &Handler{similiService: similiService}
}

func (h *Handler) CalculateHybridSimilarityHandler(c *gin.Context) {
	var req api.SimiligoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seq, err := h.similiService.CalculateHybridSimilarity(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": seq})
}
