package stats

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sort"
	"sync"
)

const statsFile = "stats.json"

type StatsData struct {
	UniqueUsers   map[string]bool `json:"unique_users"`
	Features      map[string]int  `json:"features"`
	TotalRequests int             `json:"total_requests"`
}

var (
	data StatsData
	mu   sync.Mutex
)

func Init() {
	mu.Lock()
	defer mu.Unlock()

	raw, err := os.ReadFile(statsFile)
	if err != nil {
		data = StatsData{
			UniqueUsers: make(map[string]bool),
			Features:    make(map[string]int),
		}
		return
	}
	if err := json.Unmarshal(raw, &data); err != nil {
		data = StatsData{
			UniqueUsers: make(map[string]bool),
			Features:    make(map[string]int),
		}
	}
}

func Record(route string, params url.Values) {
	mu.Lock()
	defer mu.Unlock()

	data.TotalRequests++

	if params == nil {
		save()
		return
	}

	if user := params.Get("user"); user != "" {
		data.UniqueUsers[user] = true
	}

	trackedParams := []string{"wave", "theme", "top_languages", "hide_clan", "hide_title", "animation", "stroke", "name"}
	for _, p := range trackedParams {
		if v := params.Get(p); v != "" {
			key := fmt.Sprintf("%s=%s", p, v)
			data.Features[key]++
		}
	}

	save()
}

func save() {
	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(statsFile, raw, 0644)
}

type FeatureCount struct {
	Name  string
	Count int
}

type StatsView struct {
	TotalRequests   int
	UniqueUserCount int
	Features        []FeatureCount
}

func GetStats() StatsView {
	mu.Lock()
	defer mu.Unlock()

	view := StatsView{
		TotalRequests:   data.TotalRequests,
		UniqueUserCount: len(data.UniqueUsers),
	}

	for k, v := range data.Features {
		view.Features = append(view.Features, FeatureCount{Name: k, Count: v})
	}
	sort.Slice(view.Features, func(i, j int) bool {
		return view.Features[i].Count > view.Features[j].Count
	})

	return view
}
