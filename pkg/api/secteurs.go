package api

import (
	"encoding/json"
	"github.com/Ojelaidi/similigo-api/ressources"
	"log"
)

type Secteur struct {
	Children *[]Secteur `json:"children,omitempty"`
	ID       string     `json:"id"`
	Label    string     `json:"label"`
	Level    *int       `json:"level,omitempty"`
	Slug     string     `json:"slug"`
}

var SecteurList []Secteur

type SecteurMatch struct {
	Secteur Secteur `json:"secteur"`
	Score   float64 `json:"score"`
}

type SecteurMatchHeap []SecteurMatch

func (h SecteurMatchHeap) Len() int           { return len(h) }
func (h SecteurMatchHeap) Less(i, j int) bool { return h[i].Score > h[j].Score }
func (h SecteurMatchHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SecteurMatchHeap) Push(x interface{}) {
	*h = append(*h, x.(SecteurMatch))
}

func (h *SecteurMatchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func init() {
	err := json.Unmarshal(ressources.SecteursBytes, &SecteurList)
	if err != nil {
		log.Fatalf("Failed to unmarshal sectors: %v", err)
	}
}

// Collects all labels from a sector and its children
func CollectLabels(secteur Secteur) []string {
	var labels []string
	labels = append(labels, secteur.Label)

	if secteur.Children != nil {
		for _, child := range *secteur.Children {
			labels = append(labels, CollectLabels(child)...)
		}
	}
	return labels
}
