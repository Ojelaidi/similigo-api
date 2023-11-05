package api

import (
	"container/heap"
	"encoding/json"
	"github.com/Ojelaidi/similigo-api/ressources"
	"log"
)

type Fonction struct {
	Children []Fonction `json:"children,omitempty"`
	ID       string     `json:"id"`
	Label    string     `json:"label"`
	Level    *int       `json:"level,omitempty"`
	Slug     string     `json:"slug"`
}

type JobSEO struct {
	Keyword                        string `json:"keyword"`
	AssociatedOccupationalCategory string `json:"associatedOccupationalCategory"`
	Weight                         int    `json:"weight"`
}

type CombinedFunction struct {
	ID               string   `json:"id"`
	Labels           []string `json:"labels"`
	SimilarROME      []string `json:"similarRome,omitempty"`
	Keyword          string   `json:"keyword,omitempty"`
	OccupationalCode string   `json:"occupationalCode,omitempty"`
	Weight           int      `json:"weight,omitempty"`
}

/*var level3Pattern = regexp.MustCompile(`^[A-Za-z]+\d{4}$`)

func IsLevel3Function(fonction Fonction) bool {
	return level3Pattern.MatchString(fonction.ID)
}

func EnrichFunctionsList(functions []Fonction, jobsSEO []JobSEO, similarRome map[string][]string) []CombinedFunction {
	enrichedFunctions := make([]CombinedFunction, 0)

	jobSEOIndex := make(map[string]JobSEO)
	for _, job := range jobsSEO {
		jobSEOIndex[job.AssociatedOccupationalCategory] = job
	}

	for _, function := range functions {
		if !IsLevel3Function(function) {
			continue
		}
		var combinedFunction CombinedFunction

		// Collect all labels from the function and its children.
		labels := extractLabels([]Fonction{function})

		// Retrieve similar ROME codes if they exist.
		similarCodes, found := similarRome[function.ID]

		// Check if the function's ID matches any AssociatedOccupationalCategory in JobSEOList.
		jobSEO, hasJobSEO := jobSEOIndex[function.ID]

		// Create the combined struct.
		combinedFunction = CombinedFunction{
			ID:               function.ID,
			Labels:           labels,
			SimilarROME:      similarCodes,
			Keyword:          jobSEO.Keyword,
			OccupationalCode: jobSEO.AssociatedOccupationalCategory,
			Weight:           jobSEO.Weight,
		}

		if found || hasJobSEO {
			enrichedFunctions = append(enrichedFunctions, combinedFunction)
		} else {
			// Even if there is no associated job or similar ROME codes, include the function.
			combinedFunction.SimilarROME = []string{} // Ensure the field is initialized.
			enrichedFunctions = append(enrichedFunctions, combinedFunction)
		}
	}

	return enrichedFunctions
}

func extractLabels(functions []Fonction) []string {
	var labels []string
	for _, function := range functions {
		labels = append(labels, function.Label)
		labels = append(labels, extractLabels(function.Children)...)
	}
	return labels
}*/

var (
	FunctionsList   []Fonction
	SimilarROMECode = make(map[string][]string)
	JobSEOList      []JobSEO

	EnrichedFunctionsList []CombinedFunction
)

type JobSEOMatch struct {
	JobSEO JobSEO
	Score  float64
}
type JobSEOMatchHeap []JobSEOMatch

func (h JobSEOMatchHeap) Len() int {
	return len(h)
}

func (h JobSEOMatchHeap) Less(i, j int) bool {
	return h[i].Score > h[j].Score
}

func (h JobSEOMatchHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *JobSEOMatchHeap) Push(x interface{}) {
	*h = append(*h, x.(JobSEOMatch))
}

func (h *JobSEOMatchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func InitJobSEOMatchHeap() *JobSEOMatchHeap {
	h := &JobSEOMatchHeap{}
	heap.Init(h)
	return h
}

func init() {
	err := json.Unmarshal(ressources.TestFonctions, &FunctionsList)
	if err != nil {
		log.Fatalf("Failed to unmarshal functions: %v", err)
	}
	_ = json.Unmarshal(ressources.SimilarROMECodeBytes, &SimilarROMECode)
	_ = json.Unmarshal(ressources.NewJobsSeo, &JobSEOList)
	_ = json.Unmarshal(ressources.EnrichedFunctions, &EnrichedFunctionsList)

	/*enrichedFunctions := EnrichFunctionsList(FunctionsList, JobSEOList, SimilarROMECode)

	// Convert enrichedFunctions to JSON for output or further processing.
	enrichedFunctionsJSON, err := json.Marshal(enrichedFunctions)
	if err != nil {
		log.Fatalf("Failed to marshal enriched functions: %v", err)
	}
	err = ioutil.WriteFile("ressources/enriched_functions.json", enrichedFunctionsJSON, 0644)
	if err != nil {
		log.Fatalf("Failed to write enriched functions to file: %v", err)
	}*/

}
