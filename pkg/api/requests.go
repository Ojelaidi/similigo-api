package api

type SimiligoRequest struct {
	String1              string          `json:"string1" swag:"example,title"`
	String2              string          `json:"string2" swag:"example,keyword"`
	NgramSize            int             `json:"ngramSize"`
	WordSimWeight        float64         `json:"wordSimWeight"`
	NgramSimWeight       float64         `json:"ngramSimWeight"`
	ContainmentSimWeight float64         `json:"containmentSimWeight"`
	CustomStopWords      map[string]bool `json:"customStopWords"`
}
