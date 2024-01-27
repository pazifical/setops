package setops

func Intersection[K comparable](collections [][]K) []K {
	occurrenceCounter := make(map[K]int)
	for _, collection := range collections {
		for _, item := range collection {
			occurrenceCounter[item] += 1
		}
	}
	intersection := make([]K, 0)
	for item, count := range occurrenceCounter {
		if count == len(collections) {
			intersection = append(intersection, item)
		}
	}
	return intersection
}

func Union[K comparable](collections [][]K) []K {
	unionMap := make(map[K]bool)
	for _, collection := range collections {
		for _, item := range collection {
			unionMap[item] = true
		}
	}
	union := make([]K, 0)
	for hit, _ := range unionMap {
		union = append(union, hit)
	}
	return union
}
