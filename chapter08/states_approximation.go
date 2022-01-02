package chapter08

// States will find coverage for stations.
// The book has a number of problems here, and formatting issues to boot. Luckily, I can read the code
// from the book's GitHub repo.
func States() map[string]struct{} {
	statesNeeded := map[string]struct{}{
		"mt": {},
		"wa": {},
		"or": {},
		"id": {},
		"nv": {},
		"ut": {},
		"ca": {},
		"az": {},
	}
	stations := map[string]map[string]struct{}{
		"kone": {
			"id": {},
			"nv": {},
			"ut": {},
		},
		"ktwo": {
			"wa": {},
			"id": {},
			"mt": {},
		},
		"kthree": {
			"or": {},
			"nv": {},
			"ca": {},
		},
		"kfour": {
			"nv": {},
			"ut": {},
		},
		"kfive": {
			"ca": {},
			"az": {},
		},
	}
	finalStations := make(map[string]struct{})
	for len(statesNeeded) > 0 {
		bestStation := ""
		statesCovered := make(map[string]struct{})
		for station, states := range stations {
			covered := intersect(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}
		finalStations[bestStation] = struct{}{}
		remove(statesNeeded, statesCovered)
	}
	return finalStations
}

func remove(needed map[string]struct{}, covered map[string]struct{}) {
	for k := range covered {
		delete(needed, k)
	}
}

func intersect(needed map[string]struct{}, states map[string]struct{}) map[string]struct{} {
	ints := make(map[string]struct{})
	from, to := needed, states
	if len(needed) < len(states) {
		from, to = states, needed
	}
	for k := range from {
		if _, ok := to[k]; ok {
			ints[k] = struct{}{}
		}
	}
	return ints
}
