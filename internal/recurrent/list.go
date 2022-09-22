package recurrent

import "time"

// RuleList это представление сразу нескольких реккурентных выражений
type RuleList struct {
	rules []Rule
}

func (r *RuleList) Add(rule Rule) {
	r.rules = append(r.rules, rule)
}

func (r *RuleList) IsIntersect(begin time.Time, dur time.Duration) bool {
	for i := range r.rules {
		if r.rules[i].IsIntersect(begin, dur) {
			return true
		}
	}

	return false
}
