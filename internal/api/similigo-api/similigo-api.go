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

	opts := []similigo.Option{
		similigo.WithNgramSize(request.NgramSize),
		similigo.WithWordSimWeight(request.WordSimWeight),
		similigo.WithNgramSimWeight(request.NgramSimWeight),
		similigo.WithContainmentSimWeight(request.ContainmentSimWeight),
		similigo.WithCustomStopWords(request.CustomStopWords),
	}

	score := similigo.CalculateHybridSimilarity(request.String1, request.String2, opts...)

	return score, nil
}

func (s *Service) CalculateBestNMatches(request api.SimiligoListRequest) ([]api.Match, error) {

	opts := []similigo.Option{
		similigo.WithNgramSize(request.NgramSize),
		similigo.WithWordSimWeight(request.WordSimWeight),
		similigo.WithNgramSimWeight(request.NgramSimWeight),
		similigo.WithContainmentSimWeight(request.ContainmentSimWeight),
		similigo.WithCustomStopWords(request.CustomStopWords),
	}

	similigoMatches := similigo.FindBestNMatchesInList(request.String1, request.String2, request.N, opts...)

	apiMatches := make([]api.Match, len(similigoMatches))
	for i, m := range similigoMatches {
		apiMatches[i] = api.Match{
			Text:  m.Text,
			Score: m.Score,
		}
	}

	return apiMatches, nil
}
