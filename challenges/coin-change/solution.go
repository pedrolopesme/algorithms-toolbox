package coin_change

// Solution for the Coins Change problem, where
// 	- 	coin 			: coin value
// 	- 	availableCoins 	: available coins of the given value
// 	- 	val 			: value to be calculated
func calculateCoins(coin int, availableCoins int, val int) (rest int, neededCoins int) {
	neededCoins = val / coin
	if neededCoins > availableCoins {
		neededCoins = availableCoins
	}

	rest = val - (neededCoins * coin)
	return
}

// 	- 	coins 	: coins used to compose the change, value -> available qtt
// 	- 	val 	: value to be calculated
//
// 	result:
//		map containing a list of coins -> needed quantity
func Solution(coins map[int]int, val float64) (result map[int]int) {
	intVal := int(val * 100)
	for coin, availableCoins := range coins {
		rest, neededCoins := calculateCoins(coin, availableCoins, intVal)
		result[coin] = neededCoins

		if rest <= 0 {
			return
		}
	}
	return
}
