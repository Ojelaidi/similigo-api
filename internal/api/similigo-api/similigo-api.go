package similigo_api

import (
	"container/heap"
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

// CalculateTopJobSEOFunctionMatches computes the similarity scores for offer titles against
// a list of functions with preprocessed labels and similar Rome codes, returning the top N JobSEO matches.
func (s *Service) CalculateTopJobSEOFunctionMatches(offerTitle string, n int) ([]api.JobSEOMatch, error) {
	if n == 0 {
		n = api.DefaultNSize
	}
	h := api.InitJobSEOMatchHeap()

	// Default options for similarity calculation, adjust as needed
	opts := []similigo.Option{
		similigo.WithNgramSize(3),
		similigo.WithCustomStopWords(api.CustomStpopWords),
	}

	jobSEOMap := make(map[string]api.JobSEO)

	addedJobSEOMap := make(map[string]bool)

	for _, function := range api.EnrichedFunctionsList {
		jobSEOMap[function.ID] = api.JobSEO{
			Keyword:                        function.Keyword,
			AssociatedOccupationalCategory: function.OccupationalCode,
			Weight:                         function.Weight,
		}
	}

	for _, function := range api.EnrichedFunctionsList {
		for _, label := range function.Labels {
			score := similigo.CalculateHybridSimilarity(offerTitle, label, opts...)
			if !addedJobSEOMap[function.ID] {
				heap.Push(h, api.JobSEOMatch{JobSEO: jobSEOMap[function.ID], Score: score})
				addedJobSEOMap[function.ID] = true
			}
		}

		for _, similarRome := range function.SimilarROME {
			for _, ef := range api.EnrichedFunctionsList {
				if ef.ID == similarRome {
					for _, label := range ef.Labels {
						score := similigo.CalculateHybridSimilarity(offerTitle, label, opts...)
						if !addedJobSEOMap[ef.ID] {
							heap.Push(h, api.JobSEOMatch{JobSEO: jobSEOMap[ef.ID], Score: score})
							addedJobSEOMap[ef.ID] = true
						}
					}
				}
			}
		}
	}

	topJobSEOMatches := make([]api.JobSEOMatch, 0, n)
	for i := 0; i < n && h.Len() > 0; i++ {
		m := heap.Pop(h).(api.JobSEOMatch)
		topJobSEOMatches = append(topJobSEOMatches, m)
	}

	return topJobSEOMatches, nil
}
