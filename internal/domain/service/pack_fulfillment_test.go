package service

import (
	"reflect"
	"testing"
)

func TestPackFulfillmentService_FulfillOrder(t *testing.T) {
	tests := []struct {
		name      string
		quantity  int
		packSizes []int
		expect    PackFulfillmentResult
	}{
		{
			name:      "Exact match",
			quantity:  1000,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expect:    PackFulfillmentResult{TotalItems: 1000, Packs: map[int]int{1000: 1}},
		},
		{
			name:      "Minimal over-ship",
			quantity:  251,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expect:    PackFulfillmentResult{TotalItems: 500, Packs: map[int]int{500: 1}},
		},
		{
			name:      "Multiple packs",
			quantity:  750,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expect:    PackFulfillmentResult{TotalItems: 750, Packs: map[int]int{500: 1, 250: 1}},
		},
		{
			name:      "Large order",
			quantity:  5200,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expect:    PackFulfillmentResult{TotalItems: 5250, Packs: map[int]int{5000: 1, 250: 1}},
		},
	}

	svc := &PackFulfillmentService{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := svc.FulfillOrder(tt.quantity, tt.packSizes)
			if got.TotalItems != tt.expect.TotalItems {
				t.Errorf("TotalItems: got %d, want %d", got.TotalItems, tt.expect.TotalItems)
			}
			if !reflect.DeepEqual(got.Packs, tt.expect.Packs) {
				t.Errorf("Packs: got %v, want %v", got.Packs, tt.expect.Packs)
			}
		})
	}
}
