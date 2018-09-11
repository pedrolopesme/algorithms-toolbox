package coin_change

import (
	"fmt"
	"os"
	"log"
	"strconv"
)

func main() {

	// Total of coins to be used as change
	coins := map[int]int{50: 100, 25: 100, 10: 100, 5: 100, 1: 100}

	var valueToBeCharged, valueReceived float64
	var err error
	if valueToBeCharged, err = strconv.ParseFloat(os.Args[0], 64); err != nil {
		log.Fatal("It was impossible to convert ", valueToBeCharged, " to float")
	}

	if valueReceived, err = strconv.ParseFloat(os.Args[0], 64); err != nil {
		log.Fatal("It was impossible to convert ", valueReceived, " to float")
	}

	change := valueToBeCharged - valueReceived
	changeCoins := Solution(coins, change)

	var totalCoins int
	var totalChange float64

	fmt.Println("Total value:", valueToBeCharged, "Received value:", valueReceived)
	for coin, qtt := range changeCoins {
		fmt.Println(qtt, "coins of", coin)
		totalChange += float64((coin * qtt) / 100)
		totalCoins++
	}
	fmt.Println("change:", totalChange, ", coins:", 5)
}
