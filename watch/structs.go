package watch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Init(){
	CombinedCachedComplete = make(map[string]CombinedEndpoint)
	go updateCache()
}

// updateCache updates the combined cache every ten minutes for all active leagues.
func updateCache(){
	activeLeagues := getCurrentLeagues()
	for{
		for _, league := range activeLeagues {
			getCombined(league.Name)
		}
	time.Sleep(10 * time.Minute)
	}
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

type leagues []struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// getCombined is getting the current combined json from the combined endpoint.
func getCombined(league string){
	url := "https://api.poe.watch" + "/combined?league=" + league
	url = strings.ReplaceAll(url, " ", "%20")

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data CombinedEndpoint
	err = json.Unmarshal(body, &data)
	if len(data.Items) > 0 {
		CombinedCachedComplete[league] = data
	}
}

// getCurrentLeagues is getting the current active leagues from the poe.watch api.
func getCurrentLeagues() leagues{
	url := "https://api.poe.watch/leagues"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return leagues{}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return leagues{}
	}

	var data leagues
	err = json.Unmarshal(body, &data)
	if err != nil{
		log.Println(err)
	}

	return data
}