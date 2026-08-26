package reports

import (
	"fmt"
	"frontend_go/internal/model"
	"sort"
)

type Summary struct {
	Total, Archived, Active int
	ByAction                map[string]int
}

func Build(rs []model.Record) Summary {
	s := Summary{ByAction: map[string]int{}}
	for _, r := range rs {
		s.Total++
		s.ByAction[r.Action]++
		if r.Archived {
			s.Archived++
		} else {
			s.Active++
		}
	}
	return s
}
func Sorted(rs []model.Record) []model.Record {
	out := append([]model.Record{}, rs...)
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out
}
func Format(s Summary) string {
	return fmt.Sprintf("total=%d active=%d archived=%d", s.Total, s.Active, s.Archived)
}
func Completion(s Summary) float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Archived) / float64(s.Total)
}
func Actions(s Summary) []string {
	out := []string{}
	for k := range s.ByAction {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
