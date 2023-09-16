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

// ----------------------------------------------------------------------------

// Pokemon Structures

// Pokemon contains all the information on a single Pokemon
type Pokemon struct {
	Abilities              []PokemonAbility   `json:"abilities"`
	BaseExperience         int                `json:"base_experience"`
	Forms                  []PokemonFormNR    `json:"forms"`
	GameIndices            []VersionGameIndex `json:"game_indices"`
	Height                 int                `json:"height"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	ID                     int                `json:"id"`
	IsDefault              bool               `json:"is_default"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	Name                   string             `json:"name"`
	Order                  int                `json:"order"`
	PastTypes              []PokemonTypePast  `json:"past_types"`
	Species                PokemonSpeciesNR   `json:"species"`
	Sprites                PokemonSprites     `json:"sprites"`
	Stats                  []Stats            `json:"stats"`
	Types                  []PokemonType      `json:"types"`
	Weight                 int                `json:"weight"`
}

type AbilityNR struct {
	NamedAPIResource
}

type PokemonAbility struct {
	Ability  AbilityNR `json:"ability"`
	IsHidden bool      `json:"is_hidden"`
	Slot     int       `json:"slot"`
}

type PokemonFormNR struct {
	NamedAPIResource
}

type VersionGameIndex struct {
	GameIndex int       `json:"game_index"`
	Version   VersionNR `json:"version"`
}

type ItemNR struct {
	NamedAPIResource
}

type PokemonHeldItemVersion struct {
	Rarity  int       `json:"rarity"`
	Version VersionNR `json:"version"`
}

type PokemonHeldItem struct {
	Item           ItemNR                   `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type MoveNR struct {
	NamedAPIResource
}

type MoveLearnMethodNR struct {
	NamedAPIResource
}

type VersionGroupNR struct {
	NamedAPIResource
}

type PokemonMoveVersion struct {
	LevelLearnedAt  int               `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethodNR `json:"move_learn_method"`
	VersionGroup    VersionGroupNR    `json:"version_group"`
}

type PokemonMove struct {
	Move                MoveNR               `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type GenerationNR struct {
	NamedAPIResource
}

type TypeNR struct {
	NamedAPIResource
}

type PokemonType struct {
	Slot int    `json:"slot"`
	Type TypeNR `json:"type"`
}

type PokemonTypePast struct {
	Generation GenerationNR  `json:"generation"`
	Types      []PokemonType `json:"types"`
}

type PokemonSpeciesNR struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// type DreamWorld struct {
// 	FrontDefault string `json:"front_default"`
// 	FrontFemale  any    `json:"front_female"`
// }

// type Home struct {
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type OfficialArtwork struct {
// 	FrontDefault string `json:"front_default"`
// 	FrontShiny   string `json:"front_shiny"`
// }

// type Other struct {
// 	DreamWorld      DreamWorld      `json:"dream_world"`
// 	Home            Home            `json:"home"`
// 	OfficialArtwork OfficialArtwork `json:"official-artwork"`
// }

// type RedBlue struct {
// 	BackDefault      string `json:"back_default"`
// 	BackGray         string `json:"back_gray"`
// 	BackTransparent  string `json:"back_transparent"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontGray        string `json:"front_gray"`
// 	FrontTransparent string `json:"front_transparent"`
// }

// type Yellow struct {
// 	BackDefault      string `json:"back_default"`
// 	BackGray         string `json:"back_gray"`
// 	BackTransparent  string `json:"back_transparent"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontGray        string `json:"front_gray"`
// 	FrontTransparent string `json:"front_transparent"`
// }

// type GenerationI struct {
// 	RedBlue RedBlue `json:"red-blue"`
// 	Yellow  Yellow  `json:"yellow"`
// }

// type Crystal struct {
// 	BackDefault           string `json:"back_default"`
// 	BackShiny             string `json:"back_shiny"`
// 	BackShinyTransparent  string `json:"back_shiny_transparent"`
// 	BackTransparent       string `json:"back_transparent"`
// 	FrontDefault          string `json:"front_default"`
// 	FrontShiny            string `json:"front_shiny"`
// 	FrontShinyTransparent string `json:"front_shiny_transparent"`
// 	FrontTransparent      string `json:"front_transparent"`
// }

// type Gold struct {
// 	BackDefault      string `json:"back_default"`
// 	BackShiny        string `json:"back_shiny"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontTransparent string `json:"front_transparent"`
// }

// type Silver struct {
// 	BackDefault      string `json:"back_default"`
// 	BackShiny        string `json:"back_shiny"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontTransparent string `json:"front_transparent"`
// }

// type GenerationIi struct {
// 	Crystal Crystal `json:"crystal"`
// 	Gold    Gold    `json:"gold"`
// 	Silver  Silver  `json:"silver"`
// }

// type Emerald struct {
// 	FrontDefault string `json:"front_default"`
// 	FrontShiny   string `json:"front_shiny"`
// }

// type FireredLeafgreen struct {
// 	BackDefault  string `json:"back_default"`
// 	BackShiny    string `json:"back_shiny"`
// 	FrontDefault string `json:"front_default"`
// 	FrontShiny   string `json:"front_shiny"`
// }

// type RubySapphire struct {
// 	BackDefault  string `json:"back_default"`
// 	BackShiny    string `json:"back_shiny"`
// 	FrontDefault string `json:"front_default"`
// 	FrontShiny   string `json:"front_shiny"`
// }

// type GenerationIii struct {
// 	Emerald          Emerald          `json:"emerald"`
// 	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
// 	RubySapphire     RubySapphire     `json:"ruby-sapphire"`
// }

// type DiamondPearl struct {
// 	BackDefault      string `json:"back_default"`
// 	BackFemale       any    `json:"back_female"`
// 	BackShiny        string `json:"back_shiny"`
// 	BackShinyFemale  any    `json:"back_shiny_female"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type HeartgoldSoulsilver struct {
// 	BackDefault      string `json:"back_default"`
// 	BackFemale       any    `json:"back_female"`
// 	BackShiny        string `json:"back_shiny"`
// 	BackShinyFemale  any    `json:"back_shiny_female"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type Platinum struct {
// 	BackDefault      string `json:"back_default"`
// 	BackFemale       any    `json:"back_female"`
// 	BackShiny        string `json:"back_shiny"`
// 	BackShinyFemale  any    `json:"back_shiny_female"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type GenerationIv struct {
// 	DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
// 	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
// 	Platinum            Platinum            `json:"platinum"`
// }

// type Animated struct {
// 	BackDefault      string `json:"back_default"`
// 	BackFemale       any    `json:"back_female"`
// 	BackShiny        string `json:"back_shiny"`
// 	BackShinyFemale  any    `json:"back_shiny_female"`
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type BlackWhite struct {
// 	Animated         Animated `json:"animated"`
// 	BackDefault      string   `json:"back_default"`
// 	BackFemale       any      `json:"back_female"`
// 	BackShiny        string   `json:"back_shiny"`
// 	BackShinyFemale  any      `json:"back_shiny_female"`
// 	FrontDefault     string   `json:"front_default"`
// 	FrontFemale      any      `json:"front_female"`
// 	FrontShiny       string   `json:"front_shiny"`
// 	FrontShinyFemale any      `json:"front_shiny_female"`
// }

// type GenerationV struct {
// 	BlackWhite BlackWhite `json:"black-white"`
// }

// type OmegarubyAlphasapphire struct {
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type XY struct {
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type GenerationVi struct {
// 	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
// 	XY                     XY                     `json:"x-y"`
// }

// type Icons struct {
// 	FrontDefault string `json:"front_default"`
// 	FrontFemale  any    `json:"front_female"`
// }

// type UltraSunUltraMoon struct {
// 	FrontDefault     string `json:"front_default"`
// 	FrontFemale      any    `json:"front_female"`
// 	FrontShiny       string `json:"front_shiny"`
// 	FrontShinyFemale any    `json:"front_shiny_female"`
// }

// type GenerationVii struct {
// 	Icons             Icons             `json:"icons"`
// 	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
// }

// type GenerationViii struct {
// 	Icons Icons `json:"icons"`
// }

// type Versions struct {
// 	GenerationI    GenerationI    `json:"generation-i"`
// 	GenerationIi   GenerationIi   `json:"generation-ii"`
// 	GenerationIii  GenerationIii  `json:"generation-iii"`
// 	GenerationIv   GenerationIv   `json:"generation-iv"`
// 	GenerationV    GenerationV    `json:"generation-v"`
// 	GenerationVi   GenerationVi   `json:"generation-vi"`
// 	GenerationVii  GenerationVii  `json:"generation-vii"`
// 	GenerationViii GenerationViii `json:"generation-viii"`
// }

type PokemonSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
	// Other            Other    `json:"other"`
	// Versions         Versions `json:"versions"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
