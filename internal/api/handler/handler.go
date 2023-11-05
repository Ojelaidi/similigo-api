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

func (h *Handler) CalculateBestNMatchesHandler(c *gin.Context) {
	var req api.SimiligoListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matches, err := h.similiService.CalculateBestNMatches(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"matches": matches})
}

func (h *Handler) MatchFunctionHandler(c *gin.Context) {
	var req api.JobMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matches, err := h.similiService.CalculateTopJobSEOFunctionMatches(req.OfferTitle, req.N)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}
func (h *Handler) MatchSectorHandler(c *gin.Context) {
	var req api.JobMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matches, err := h.similiService.CalculateTopSecteurMatches(req.OfferTitle, req.N)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}
