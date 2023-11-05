package ressources

import _ "embed"

//go:embed "fonctions.json"
var TestFonctions []byte

//go:embed "similarRomeCode.json"
var SimilarROMECodeBytes []byte

//go:embed "new-jobs-seo.json"
var NewJobsSeo []byte

//go:embed "enriched_functions.json"
var EnrichedFunctions []byte
