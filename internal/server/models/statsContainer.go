package models

type StatsContainer struct {
	Title string
	Stats []Stat
}

func New(title string, stats []Stat) *StatsContainer {
	return &StatsContainer{Title: title, Stats: stats}
}
