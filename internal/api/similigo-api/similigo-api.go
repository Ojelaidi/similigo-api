package similigo_api

import (
	"github.com/Ojelaidi/similigo"
	"github.com/Ojelaidi/similigo-api/pkg/api"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CalculateHybridSimilarity(request api.SimiligoRequest) (float64, error) {
	var customStopWords []string
	for word := range request.CustomStopWords {
		customStopWords = append(customStopWords, word)
	}

	opts := []similigo.Option{
		similigo.WithNgramSize(request.NgramSize),
		similigo.WithWordSimWeight(request.WordSimWeight),
		similigo.WithNgramSimWeight(request.NgramSimWeight),
		similigo.WithContainmentSimWeight(request.ContainmentSimWeight),
		similigo.WithCustomStopWords(customStopWords),
	}

	score := similigo.CalculateHybridSimilarity(request.String1, request.String2, opts...)

	return score, nil
}
