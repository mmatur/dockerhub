package docker

import "time"

type OfficialResults []OfficialResult

func (a OfficialResults) Len() int { return len(a) }
func (a OfficialResults) Swap(i, j int) {
	a[i].PullCount, a[j].PullCount = a[j].PullCount, a[i].PullCount
}
func (a OfficialResults) Less(i, j int) bool { return a[i].PullCount > a[j].PullCount }

type OfficialPage struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Results  []OfficialResult `json:"results"`
}

type OfficialResult struct {
	User           string    `json:"user"`
	Name           string    `json:"name"`
	Namespace      string    `json:"namespace"`
	RepositoryType string    `json:"repository_type"`
	Status         int       `json:"status"`
	Description    string    `json:"description"`
	IsPrivate      bool      `json:"is_private"`
	SAutomated     bool      `json:"s_automated"`
	CanEdit        bool      `json:"can_edit"`
	StarCount      int       `json:"star_count"`
	PullCount      int64     `json:"pull_count"`
	LastUpdated    time.Time `json:"last_updated"`
	IsMigrated     bool      `json:"is_migrated"`
}

type Results []Result

func (a Results) Len() int           { return len(a) }
func (a Results) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Results) Less(i, j int) bool { return a[i].Popularity > a[j].Popularity }

type Page struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"summaries"`
}

type Result struct {
	Slug       string `json:"slug"`
	Source     string `json:"source"`
	Popularity int64  `json:"popularity"`
	StarCount  int    `json:"star_count"`
}
