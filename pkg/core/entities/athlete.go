package entities

import "log"

type Athlete struct {
	Name    string `json:"athlete" db:"name"`
	Age     int    `json:"age" db:"age"`
	Country string `json:"country" db:"country"`
	Year    int    `json:"year" db:"year"`
	Date    string `json:"date" db:"date"`
	Sport   string `json:"sport" db:"sport"`
	Gold    int    `json:"gold" db:"gold"`
	Silver  int    `json:"silver" db:"silver"`
	Bronze  int    `json:"bronze" db:"bronze"`
}

type MedalPackage struct {
	Gold   int `json:"gold"`
	Silver int `json:"silver"`
	Bronze int `json:"bronze"`
	Total  int `json:"total"`
}

func (p *MedalPackage) Add(gold, silver, bronze int) {
	p.Gold += gold
	p.Silver += silver
	p.Bronze += bronze
	p.Total += gold + silver + bronze
}

type MedalHistory struct {
	Medals  *MedalPackage         `json:"medals"`
	History map[int]*MedalPackage `json:"history"`
}

type AthleteInfo struct {
	Athlete      string                  `json:"athlete"`
	Countries    []string                `json:"countries"`
	SportsMedals map[string]MedalHistory `json:"sports_medals"`
}

func BuildAthleteInfo(entries []Athlete) AthleteInfo {
	if len(entries) == 0 {
		return AthleteInfo{}
	}
	log.Println("start build")

	info := AthleteInfo{Athlete: entries[0].Name, SportsMedals: make(map[string]MedalHistory)}
	countries := make(map[string]struct{})

	for _, entry := range entries {
		countries[entry.Country] = struct{}{}
		if _, ok := info.SportsMedals[entry.Sport]; !ok {
			info.Countries = append(info.Countries, entry.Sport)
			info.SportsMedals[entry.Sport] = MedalHistory{Medals: new(MedalPackage), History: make(map[int]*MedalPackage)}
		}

		log.Println("before add")
		info.SportsMedals[entry.Sport].Medals.Add(entry.Gold, entry.Silver, entry.Bronze)

		if _, ok := info.SportsMedals[entry.Sport].History[entry.Year]; !ok {
			info.SportsMedals[entry.Sport].History[entry.Year] = &MedalPackage{}
		}

		info.SportsMedals[entry.Sport].History[entry.Year].Add(entry.Gold, entry.Silver, entry.Bronze)
	}

	return info
}
