package recurrent

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	// проверим только, что пасрер работает
	_, err := Parse("FREQ=DAILY;INTERVAL=2;COUNT=5", time.Now(), 10*time.Second)
	assert.NoError(t, err, "correct rrule expression must be parsed without errors")
}

func testIntersections(t *testing.T, r Rule, startTime time.Time) {
	// ищем пересечение с интервалом, который был задолго до startTime
	assert.False(t, r.IsIntersect(startTime.Add(-2*time.Hour), 1*time.Hour), "must not intersect")

	// пересечение с интервалом, который заканчивается сразу перед startTime
	assert.False(t, r.IsIntersect(startTime.Add(-1*time.Hour), 1*time.Hour), "must not intersect")

	// пересечение с интервалом, который заканчивается чуть позже чем startTime
	assert.True(t, r.IsIntersect(startTime.Add(-1*time.Hour+1*time.Second), 1*time.Hour), "must intersect")

	// пересечение с интервалом, который заканчивается сильно больше, чем startTime
	assert.True(t, r.IsIntersect(startTime.Add(-30*time.Minute), 1*time.Hour), "must intersect")

	// пересечение с интервалом, который заканчивается на границк startTime + dur
	assert.True(t, r.IsIntersect(startTime.Add(-1*time.Second), 1*time.Hour+1*time.Second), "must intersect")

	// пересечение одинаковых интервалов
	assert.True(t, r.IsIntersect(startTime, 1*time.Hour), "must intersect")

	// интервал лежит внутри интервала
	assert.True(t, r.IsIntersect(startTime.Add(1*time.Second), 1*time.Hour-1*time.Second), "must intersect")

	// начала интервала лежит внутри, а конец снаружи
	assert.True(t, r.IsIntersect(startTime.Add(30*time.Minute), 30*time.Minute+time.Second), "must intersect")

	// начало интервала лежит на границе startTime+dur
	assert.False(t, r.IsIntersect(startTime.Add(1*time.Hour), time.Hour), "must not intersect")

	// интервал лежит после startTime + dur
	assert.False(t, r.IsIntersect(startTime.Add(2*time.Hour), time.Hour), "must not intersect")
}

func TestRule_IsIntersect(t *testing.T) {
	startTime := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	r, err := Parse("FREQ=DAILY;INTERVAL=2;COUNT=5", startTime, 1*time.Hour)
	assert.NoError(t, err, "correct rrule expression must be parsed without errors")

	// пробежимся по всем 5 повторам прибавляя 2 дня ы startTime
	for i := 0; i < 5; i, startTime = i+1, startTime.Add(2*24*time.Hour) {
		testIntersections(t, r, startTime)
	}

	// убедимся, что когда повторение окончено - пересечений нет
	assert.False(t, r.IsIntersect(startTime.Add(30*time.Minute), 30*time.Minute+time.Second), "must not intersect")

	// попробуем для неповторяющейся задачи
	r = Once(startTime, 1*time.Hour)
	testIntersections(t, r, startTime)
}
