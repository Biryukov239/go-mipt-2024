//go:build !solution

package hotelbusiness

import (
	"fmt"
	"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	var result []Load
	if len(guests) == 0 {
		return result
	}
	sort.Slice(guests, func(i, j int) bool {
		return guests[i].CheckInDate < guests[j].CheckInDate
	})
	guestCount := 0
	result = append(result, Load{guests[0].CheckInDate, guestCount + 1})
	guestCount++
	if len(guests) == 1 {
		result = append(result, Load{guests[0].CheckOutDate, guestCount - 1})
		guestCount--
		return result
	}
	for i := 1; i < len(guests); i++ {
		if guests[i-1].CheckInDate == guests[i].CheckInDate {
			result[i-1].GuestCount++
			fmt.Print("We're here!")
		}
		if guests[i].CheckInDate > guests[i-1].CheckOutDate {
			result = append(result, Load{guests[i-1].CheckOutDate, guestCount - 1})
			guestCount--
			result = append(result, Load{guests[i].CheckInDate, guestCount + 1})
			guestCount++
			result = append(result, Load{guests[i].CheckOutDate, guestCount - 1})
			guestCount--
		}
		if guests[i-1].CheckOutDate == guests[i].CheckInDate {
			result = append(result, Load{guests[i].CheckOutDate, guestCount - 1})
			guestCount--
		}
		if guests[i].CheckInDate < guests[i-1].CheckOutDate && guests[i-1].CheckInDate != guests[i].CheckInDate {
			result = append(result, Load{guests[i].CheckInDate, guestCount + 1})
			guestCount++
			result = append(result, Load{guests[i-1].CheckOutDate, guestCount - 1})
			guestCount--
			result = append(result, Load{guests[i].CheckOutDate, guestCount - 1})
			guestCount--
		}
		if guests[i-1].CheckOutDate == guests[i].CheckOutDate {
			result = append(result, Load{guests[i].CheckOutDate, 0})
		}
	}
	return result
}
