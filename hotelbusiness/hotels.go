//go:build !solution

package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	result := make([]Load, 0)
	if len(guests) == 0 {
		return result
	}
	mapResult := make(map[int]int)
	sort.Slice(guests, func(i, j int) bool {
		return guests[i].CheckInDate < guests[j].CheckInDate
	})
	for _, elem := range guests {
		mapResult[elem.CheckInDate] = 0
		mapResult[elem.CheckOutDate] = 0
	}
	mapResult[guests[0].CheckInDate]++
	for i := range guests[1:] {
		if guests[i-1].CheckOutDate == guests[i].CheckInDate {
			continue
		}
		if guests[i-1].CheckOutDate < guests[i].CheckInDate {
			mapResult[guests[i].CheckInDate]++
			continue
		}
		if guests[i].CheckOutDate > guests[i-1].CheckOutDate {
			mapResult[guests[i].CheckInDate] = mapResult[guests[i-1].CheckInDate] + 1
			mapResult[guests[i-1].CheckOutDate] = mapResult[guests[i].CheckInDate] - 1
			continue
		}
		mapResult[guests[i].CheckInDate] = mapResult[guests[i-1].CheckInDate] + 1
		mapResult[guests[i].CheckOutDate] = mapResult[guests[i].CheckInDate] - 1
	}
	for key, value := range mapResult {
		result = append(result, Load{key, value})
	}
	return result
}
