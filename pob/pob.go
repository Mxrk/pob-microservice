package pob

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"pob_api/watch"
	"strings"
)

type PathOfBuilding struct {
	XMLName xml.Name `xml:"PathOfBuilding"`
	Text    string   `xml:",chardata"`
	Build   struct {
		Text             string `xml:",chardata"`
		Level            string `xml:"level,attr"`
		TargetVersion    string `xml:"targetVersion,attr"`
		PantheonMajorGod string `xml:"pantheonMajorGod,attr"`
		Bandit           string `xml:"bandit,attr"`
		ClassName        string `xml:"className,attr"`
		AscendClassName  string `xml:"ascendClassName,attr"`
		MainSocketGroup  string `xml:"mainSocketGroup,attr"`
		ViewMode         string `xml:"viewMode,attr"`
		PantheonMinorGod string `xml:"pantheonMinorGod,attr"`
		PlayerStat       []struct {
			Text  string `xml:",chardata"`
			Stat  string `xml:"stat,attr"`
			Value string `xml:"value,attr"`
		} `xml:"PlayerStat"`
	} `xml:"Build"`
	Import struct {
		Text              string `xml:",chardata"`
		LastAccountHash   string `xml:"lastAccountHash,attr"`
		LastRealm         string `xml:"lastRealm,attr"`
		LastCharacterHash string `xml:"lastCharacterHash,attr"`
	} `xml:"Import"`
	Calcs struct {
		Text  string `xml:",chardata"`
		Input []struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name,attr"`
			String string `xml:"string,attr"`
			Number string `xml:"number,attr"`
		} `xml:"Input"`
		Section []struct {
			Text      string `xml:",chardata"`
			Collapsed string `xml:"collapsed,attr"`
			ID        string `xml:"id,attr"`
		} `xml:"Section"`
	} `xml:"Calcs"`
	Skills struct {
		Text                string `xml:",chardata"`
		SortGemsByDPSField  string `xml:"sortGemsByDPSField,attr"`
		SortGemsByDPS       string `xml:"sortGemsByDPS,attr"`
		DefaultGemQuality   string `xml:"defaultGemQuality,attr"`
		DefaultGemLevel     string `xml:"defaultGemLevel,attr"`
		ShowSupportGemTypes string `xml:"showSupportGemTypes,attr"`
		ShowAltQualityGems  string `xml:"showAltQualityGems,attr"`
		Skill               []struct {
			Text                 string `xml:",chardata"`
			MainActiveSkillCalcs string `xml:"mainActiveSkillCalcs,attr"`
			Label                string `xml:"label,attr"`
			Enabled              string `xml:"enabled,attr"`
			Slot                 string `xml:"slot,attr"`
			MainActiveSkill      string `xml:"mainActiveSkill,attr"`
			Source               string `xml:"source,attr"`
			Gem                  []struct {
				Text          string `xml:",chardata"`
				EnableGlobal2 string `xml:"enableGlobal2,attr"`
				Level         string `xml:"level,attr"`
				GemId         string `xml:"gemId,attr"`
				SkillId       string `xml:"skillId,attr"`
				EnableGlobal1 string `xml:"enableGlobal1,attr"`
				QualityId     string `xml:"qualityId,attr"`
				Quality       string `xml:"quality,attr"`
				Enabled       string `xml:"enabled,attr"`
				NameSpec      string `xml:"nameSpec,attr"`
				SkillMinion   string `xml:"skillMinion,attr"`
				SkillPart     string `xml:"skillPart,attr"`
			} `xml:"Gem"`
		} `xml:"Skill"`
	} `xml:"Skills"`
	Tree struct {
		Text       string `xml:",chardata"`
		ActiveSpec string `xml:"activeSpec,attr"`
		Spec       []struct {
			Text          string `xml:",chardata"`
			Title         string `xml:"title,attr"`
			AscendClassId string `xml:"ascendClassId,attr"`
			Nodes         string `xml:"nodes,attr"`
			TreeVersion   string `xml:"treeVersion,attr"`
			ClassId       string `xml:"classId,attr"`
			EditedNodes   string `xml:"EditedNodes"`
			URL           string `xml:"URL"`
			Sockets       struct {
				Text   string `xml:",chardata"`
				Socket []struct {
					Text   string `xml:",chardata"`
					NodeId string `xml:"nodeId,attr"`
					ItemId string `xml:"itemId,attr"`
				} `xml:"Socket"`
			} `xml:"Sockets"`
		} `xml:"Spec"`
	} `xml:"Tree"`
	Notes    string `xml:"Notes"`
	TreeView struct {
		Text                string `xml:",chardata"`
		SearchStr           string `xml:"searchStr,attr"`
		ZoomY               string `xml:"zoomY,attr"`
		ShowHeatMap         string `xml:"showHeatMap,attr"`
		ZoomLevel           string `xml:"zoomLevel,attr"`
		ShowStatDifferences string `xml:"showStatDifferences,attr"`
		ZoomX               string `xml:"zoomX,attr"`
	} `xml:"TreeView"`
	Items struct {
		Text               string `xml:",chardata"`
		ActiveItemSet      string `xml:"activeItemSet,attr"`
		UseSecondWeaponSet string `xml:"useSecondWeaponSet,attr"`
		Item               []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Variant  string `xml:"variant,attr"`
			ModRange []struct {
				Text  string `xml:",chardata"`
				Range string `xml:"range,attr"`
				ID    string `xml:"id,attr"`
			} `xml:"ModRange"`
		} `xml:"Item"`
		Slot []struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name,attr"`
			ItemId string `xml:"itemId,attr"`
			Active string `xml:"active,attr"`
		} `xml:"Slot"`
		ItemSet []struct {
			Text               string `xml:",chardata"`
			UseSecondWeaponSet string `xml:"useSecondWeaponSet,attr"`
			Title              string `xml:"title,attr"`
			ID                 string `xml:"id,attr"`
			Slot               []struct {
				Text   string `xml:",chardata"`
				Name   string `xml:"name,attr"`
				ItemId string `xml:"itemId,attr"`
				Active string `xml:"active,attr"`
			} `xml:"Slot"`
		} `xml:"ItemSet"`
	} `xml:"Items"`
	Config struct {
		Text  string `xml:",chardata"`
		Input []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Boolean string `xml:"boolean,attr"`
			String  string `xml:"string,attr"`
			Number  string `xml:"number,attr"`
		} `xml:"Input"`
	} `xml:"Config"`
}

func GetPob(code string) []pobStruct {
	//pob, err := getContent(code)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	return parse(code)
}

func getContent(code string) (string, error) {
	resp, err := http.Get("https://pastebin.com/raw/" + code)
	if err != nil {
		log.Println(err)
		return "", errors.New("couldn't parse code")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", errors.New("couldn't parse code")
	}
	if resp.StatusCode != 200 {
		return "", errors.New("couldn't parse code")
	}
	return string(body), nil
}

type pobStruct struct {
	Name   string    `json:"name"`
	Items  []pobItem `json:"items"`
	Skills []skills  `json:"skills"`
}

type pobItem struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Links int     `json:"links"`
}

type skills struct {
	Name    string  `json:"name"`
	Quality int     `json:"quality"`
	Level   int     `json:"level"`
	Value   float64 `json:"mean"`
}

func parse(pob string) []pobStruct {
	uDec, _ := base64.URLEncoding.DecodeString(pob)
	r, _ := zlib.NewReader(bytes.NewReader(uDec))
	result, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
	}

	xmlStruct := PathOfBuilding{}
	err = xml.Unmarshal(result, &xmlStruct)
	if err != nil {
		log.Println(err)
	}

	var rets []pobStruct

	for _, set := range xmlStruct.Items.ItemSet {
		var ret pobStruct
		ret.Name = set.Title
		for _, slot := range set.Slot {
			if slot.ItemId != "0" {

				for _, item := range xmlStruct.Items.Item {
					if item.ID == slot.ItemId {
						lines := strings.Split(strings.Replace(item.Text, "\r\n", "\n", -1), "\n")

						rarity := strings.ReplaceAll(lines[1], " ", "")
						rarity = strings.TrimSpace(rarity)
						// Rarity:RARE
						s := strings.Split(rarity, ":")
						rarityValue := s[1]
						name := lines[2]

						if rarityValue == "UNIQUE" {
							var items []pobItem

							for _, iterItem := range watch.CombinedCachedComplete["Standard"].Items {
								if iterItem.Name == name {
									var item pobItem
									item.Name = iterItem.Name
									item.Value = iterItem.Mean
									item.Links = iterItem.LinkCount
									items = append(items, item)
								}
							}

							if len(items) > 0 {
								ret.Items = append(ret.Items, items...)
							}
						}

					}
				}
			}
		}
		rets = append(rets, ret)
	}

	var sks []skills

	for _, gem := range xmlStruct.Skills.Skill[0].Gem {

		for _, iterItem := range watch.CombinedCachedComplete["Standard"].Items {
			// filter out items which have similar names but aren't gems
			if iterItem.Category != "gem" {
				continue
			}

			if strings.Contains(iterItem.Name, gem.NameSpec) {
				var sk skills
				sk.Name = iterItem.Name
				sk.Value = iterItem.Mean
				sk.Level = iterItem.GemLevel
				sk.Quality = iterItem.GemQuality
				sks = append(sks, sk)
			}
		}
	}

	rets[0].Skills = sks
	return rets
}
