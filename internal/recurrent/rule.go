package recurrent

import (
	"github.com/teambition/rrule-go"
	"time"
)

// Rule - представление рекуррентно повторяющих событий. Также может представлять событие, срабатывающее только один раз
type Rule struct {
	r       *rrule.RRule
	dtStart time.Time
	dur     time.Duration
}

// Parse парсит выражение в формате RFC5545 и возвращает его представление
func Parse(expr string, startTime time.Time, dur time.Duration) (Rule, error) {
	r := Rule{
		dtStart: startTime,
		dur:     dur,
	}

	opts, err := rrule.StrToROption(expr)
	if err != nil {
		return r, err
	}

	opts.Dtstart = startTime
	r.r, err = rrule.NewRRule(*opts)
	return r, err
}

// Once возвращает представление события, которое выполнится только 1 раз
func Once(startTime time.Time, dur time.Duration) Rule {
	return Rule{
		dtStart: startTime,
		dur:     dur,
	}
}

// IsIntersect проверить, пересекается ли интевал с выражением
func (r Rule) IsIntersect(begin time.Time, dur time.Duration) bool {
	if r.r != nil {
		// смотрим ситуацию, когда повторяющаяся задача стартует до указанного интервала
		nearestBeforePoint := r.r.Before(begin, true)
		if !nearestBeforePoint.IsZero() && nearestBeforePoint.Add(r.dur).After(begin) {
			return true
		}

		// смотрим середину
		nearestAfterPoint := r.r.After(begin, true)
		if !nearestAfterPoint.IsZero() && begin.Add(dur).After(nearestAfterPoint) {
			return true
		}
	} else {
		end := begin.Add(dur)
		rEnd := r.dtStart.Add(dur)

		res := false
		// 1. начало запрошенного интервала лежит в R
		res = res || begin.Unix() >= r.dtStart.Unix() && begin.Unix() < rEnd.Unix()

		// 2. конец интервала лежит в R
		res = res || end.Unix() > r.dtStart.Unix() && end.Unix() <= rEnd.Unix()

		// 3. запрошенный интервал включает в себя R
		res = res || begin.Unix() <= r.dtStart.Unix() && end.Unix() >= rEnd.Unix()

		return res
	}

	return false
}
