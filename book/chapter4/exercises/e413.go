package exercises

import "golang.org/x/exp/slices"

type event struct {
	time  int
	class string
}

func partyMembers(events []event) int {
	maxCount := 0
	count := 0
	slices.SortFunc(events, func(a, b event) int { return a.time - b.time })
	for _, ev := range events {
		if ev.class == "entry" {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count--
		}
	}

	return maxCount
}
