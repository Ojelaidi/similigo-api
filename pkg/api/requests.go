package api

type SimiligoRequest struct {
	String1              string   `json:"string1" swag:"example,title"`
	String2              string   `json:"string2" swag:"example,keyword"`
	NgramSize            int      `json:"ngramSize"`
	WordSimWeight        float64  `json:"wordSimWeight"`
	NgramSimWeight       float64  `json:"ngramSimWeight"`
	ContainmentSimWeight float64  `json:"containmentSimWeight"`
	CustomStopWords      []string `json:"customStopWords"`
}

type SimiligoListRequest struct {
	String1              string   `json:"string1"`
	String2              []string `json:"string2"` // Now an array of strings to match against
	N                    int      `json:"n"`       // Number of best matches to find
	NgramSize            int      `json:"ngramSize"`
	WordSimWeight        float64  `json:"wordSimWeight"`
	NgramSimWeight       float64  `json:"ngramSimWeight"`
	ContainmentSimWeight float64  `json:"containmentSimWeight"`
	CustomStopWords      []string `json:"customStopWords"` // Use an array for ease of use
}

type MatchFunctionRequest struct {
	OfferTitle string `json:"offerTitle"`
	N          int    `json:"n"` // Number of best matches to find
}
