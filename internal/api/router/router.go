package router

import (
	"github.com/Ojelaidi/similigo-api/internal/api/handler"
	"github.com/Ojelaidi/similigo-api/internal/api/similigo-api"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, similigoService *similigo_api.Service) {
	h := handler.NewHandler(similigoService)

	r.POST("/calculateHybridSimilarity", h.CalculateHybridSimilarityHandler)
	r.POST("/calculateBestMatches", h.CalculateBestNMatchesHandler)
	r.POST("/calculateBestFunctionMatch", h.MatchFunctionHandler)
}
