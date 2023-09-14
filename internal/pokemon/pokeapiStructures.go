package pokemon

// NamedAPIResource contains the name and the url to retrieve further information that resource
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationAreas Structures ---------------------------------------------------

// LocationAreas contains the fields returned from location-area endpoint
type LocationAreas struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []ResultsNR `json:"results"`
}

// ResultsNR contains the list of names and urls of each location area
type ResultsNR struct {
	NamedAPIResource
}

// ----------------------------------------------------------------------------

// LocationArea Structures ----------------------------------------------------

// LocationArea contains the information for a single Location Area
type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             LocationNR            `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

// EncounterMethodNR is a Named Resource for EncounterMethod
type EncounterMethodNR struct {
	NamedAPIResource
}

// VersionNR is a Named Resource for Version
type VersionNR struct {
	NamedAPIResource
}

// EncounterVersionDetails contains details on the rate and version of an encounter
type EncounterVersionDetails struct {
	Rate    int       `json:"rate"`
	Version VersionNR `json:"version"`
}

// EncounterMethodRate contains details on the method and rate of an encounter
type EncounterMethodRate struct {
	EncounterMethod EncounterMethodNR         `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

// LocationNR is a Named Resource for Location
type LocationNR struct {
	NamedAPIResource
}

// LanguageNR is a Named Resource for Language
type LanguageNR struct {
	NamedAPIResource
}

// Name contains details on a name and its language information
type Name struct {
	Language LanguageNR `json:"language"`
	Name     string     `json:"name"`
}

// PokemonNR is a Named Resource for Pokemon
type PokemonNR struct {
	NamedAPIResource
}

// EncounterConditionValueNR is a Named Resource for EncounterConditionValue
type EncounterConditionValueNR struct {
	NamedAPIResource
}

// Encounter contains details on an encounter
type Encounter struct {
	Chance          int                         `json:"chance"`
	ConditionValues []EncounterConditionValueNR `json:"condition_values"`
	MaxLevel        int                         `json:"max_level"`
	Method          EncounterMethodNR           `json:"method"`
	MinLevel        int                         `json:"min_level"`
}

// VersionEncounterDetail contains details on the version of a set of encounters
type VersionEncounterDetail struct {
	EncounterDetails []Encounter `json:"encounter_details"`
	MaxChance        int         `json:"max_chance"`
	Version          VersionNR   `json:"version"`
}

// PokemonEncounter contains details on a Pokemon encounter
type PokemonEncounter struct {
	Pokemon        PokemonNR                `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

// ----------------------------------------------------------------------------
