package getchange

import (
	"math"
	"sort"
)

func GetChange(
	payed, price float64,
	billsAndCoinsTypes []float64, // .5 is 50 cents, .25 is 25 cents, etc.
) []int {

	sort.Float64s(billsAndCoinsTypes)

	billsAndCoins := make([]int, len(billsAndCoinsTypes))

	change := payed - price

	for i := len(billsAndCoinsTypes) - 1; i >= 0; i-- {
		billOrCoinValue := billsAndCoinsTypes[i]
		nbBillsOrCoins := math.Floor(change / billOrCoinValue)
		billsAndCoins[i] = int(nbBillsOrCoins)
		// round to 2 decimals
		change = math.Round((change-nbBillsOrCoins*billOrCoinValue)*100) / 100
	}

	return billsAndCoins
}
