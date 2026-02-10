package service

// PackFulfillmentResult holds the result of pack fulfillment.
type PackFulfillmentResult struct {
	TotalItems int
	Packs      map[int]int // pack size -> count
}

// PackFulfillmentService provides pack fulfillment logic.
type PackFulfillmentService struct{}

// FulfillOrder returns the optimal pack distribution for a given quantity and available pack sizes.
func (s *PackFulfillmentService) FulfillOrder(quantity int, packSizes []int) PackFulfillmentResult {
	// Sort packSizes descending
	for i := 0; i < len(packSizes)-1; i++ {
		for j := i + 1; j < len(packSizes); j++ {
			if packSizes[i] < packSizes[j] {
				packSizes[i], packSizes[j] = packSizes[j], packSizes[i]
			}
		}
	}

	minItems := -1
	minPacks := -1
	best := map[int]int{}

	s.dfs(packSizes, 0, quantity, map[int]int{}, &minItems, &minPacks, &best)

	return PackFulfillmentResult{
		TotalItems: minItems,
		Packs:      best,
	}
}

func (s *PackFulfillmentService) dfs(packSizes []int, idx, rem int, packs map[int]int, minItems, minPacks *int, best *map[int]int) {
	// all pack sizes have been considered
	if idx == len(packSizes) {
		// if we still have remaining items, this is not a valid solution
		if rem > 0 {
			return
		}
		items := 0
		packCount := 0
		for size, count := range packs {
			items += size * count
			packCount += count
		}
		if *minItems == -1 || items < *minItems || (items == *minItems && packCount < *minPacks) {
			*minItems = items
			*minPacks = packCount
			*best = make(map[int]int)
			for k, v := range packs {
				(*best)[k] = v
			}
		}
		return
	}
	max := (rem + packSizes[idx] - 1) / packSizes[idx]
	for i := 0; i <= max; i++ {
		newPacks := make(map[int]int)
		for k, v := range packs {
			newPacks[k] = v
		}
		if i > 0 {
			newPacks[packSizes[idx]] = i
		}
		s.dfs(packSizes, idx+1, rem-packSizes[idx]*i, newPacks, minItems, minPacks, best)
	}
}
