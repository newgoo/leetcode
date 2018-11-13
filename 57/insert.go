package _57

/**
 * Definition for an interval.
 */
type Interval struct {
	Start int
	End   int
}

//16ms
func insert(intervals []Interval, newInterval Interval) []Interval {
	nIn := newInterval
	list := make([]Interval, 0)
	b := false
	st, ed := false, false
	for index, interval := range intervals {
		if newInterval.Start > interval.End {
			list = append(list, interval)
			continue
		}

		if interval.End >= newInterval.Start && !st {
			nIn.Start = min(newInterval.Start, interval.Start)
			st = true
		}

		if interval.Start <= newInterval.End {
			nIn.End = max(interval.End, newInterval.End)
			ed = true
		}

		if interval.Start > nIn.End && ed && !b {
			list = append(list, nIn)
			b = true
		}

		if newInterval.End < interval.Start {
			if !b {
				list = append(list, nIn)
				b = true
			}
			list = append(list, intervals[index:]...)
			break
		}
	}
	if !b {
		list = append(list, nIn)
	}
	return list
}

//12ms
func insert2(intervals []Interval, newInterval Interval) []Interval {
	i := 0
	resp := make([]Interval, 0)
	for i < len(intervals) && intervals[i].End < newInterval.Start {
		resp = append(resp, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i].Start <= newInterval.End {
		newInterval.Start = min(intervals[i].Start, newInterval.Start)
		newInterval.End = max(intervals[i].End, newInterval.End)
		i++
	}
	resp = append(resp, newInterval)
	for i < len(intervals) {
		resp = append(resp, intervals[i])
		i++
	}
	return resp
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
