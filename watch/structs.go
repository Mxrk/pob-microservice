package watch

func Init(){
	CombinedCachedComplete = make(map[string]CombinedEndpoint)
}

// CombinedCachedComplete is the cache as map for the complete /combined endpoint.
var CombinedCachedComplete map[string]CombinedEndpoint

// CombinedEndpoint is the json representation for /combined
type CombinedEndpoint struct {
	Items []DefaultItemData `json:"items"`
}

type DefaultItemData struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	Group         string    `json:"group"`
	Frame         int       `json:"frame"`
	Influences    string    `json:"influences"`
	Icon          string    `json:"icon"`
	Mean          float64   `json:"mean"`
	Min           float64   `json:"min"`
	Max           float64   `json:"max"`
	Exalted       float64   `json:"exalted"`
	Daily         int       `json:"daily"`
	Change        float64   `json:"change"`
	History       []float64 `json:"history"`
	LowConfidence bool      `json:"lowConfidence"`
	Implicits     []string  `json:"implicits"`
	Explicits     []string  `json:"explicits"`
	ItemLevel     int       `json:"itemLevel"`
	PerfectPrice  float64   `json:"perfectPrice,omitempty"`
	PerfectAmount int       `json:"perfectAmount,omitempty"`
	Width         int       `json:"width"`
	Height        int       `json:"height"`

	//gems
	GemLevel     int    `json:"gemLevel,omitempty"`
	GemQuality   int    `json:"gemQuality,omitempty"`
	GemCorrupted bool   `json:"gemIsCorrupted,omitempty"`
	Type         string `json:"type,omitempty"`
	//Watchstones
	Uses int `json:"uses,omitempty"`

	//essences
	Tier int `json:"tier,omitempty"`

	//Weapons, Armours
	LinkCount int `json:"linkCount,omitempty"`

	//Maps
	MapSeries int `json:"mapSeries,omitempty"`
	MapTier   int `json:"mapTier,omitempty"`

	//Only relevant for the /hot page
	ChangeByPrice float64 `json:"changeByPrice,omitempty"`

	//Rewards for divs
	Reward      string  `json:"reward,omitempty"`
	RewardPrice float64 `json:"reward_price,omitempty"`
	RewardID    int     `json:"reward_id,omitempty"`
}