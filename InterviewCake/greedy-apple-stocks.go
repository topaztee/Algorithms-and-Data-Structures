// Suppose we could come up with the answer in one pass through the input,
//  by simply updating the 'best answer so far' as we went. What additional values would we need 
//  to keep updated as we looked at each item in our input, in order to be able to update the 'best
//   answer so far' in constant time?"

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	// "testing"
)

func generatePrices() (stockPrices []float64, numStockPrices int)  {
	//930-530
	numStockPrices = 8*60
	stockPrices = make([]float64, numStockPrices)
	rand.Seed(time.Now().UnixNano())
	for i := range stockPrices {
		stockPrices[i] = math.Floor((rand.Float64()*200)*100)/100
	}
	fmt.Println(stockPrices[0])
	return
}

func min(x,y float64) (z float64) {
	if x < y {
		return x
		
	}
	return y
}
func max(x,y float64) (z float64) {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices, _ := generatePrices()
	fmt.Println("$",getMaxProfit(prices))
}

func getMaxProfit (prices []float64) float64 {
	// var prices = []float64{100.0, 10.00, 30.00, 10.00}
	
	if len(prices) < 2 {
		fmt.Println("no enough values!")
		fmt.Errorf("error")
	}

	minPrice := prices[0]
	maxProfit := prices[1] - prices[0]
	var potentialProfit float64

	for i := 0; i < len(prices); i++ {
		currentPrice := prices[i]
		potentialProfit = currentPrice - minPrice
		maxProfit = max(maxProfit, potentialProfit)
		
		minPrice = math.Min(currentPrice, minPrice)
	}

	return maxProfit
}

// func TestEmpty(t *testing.T) {
//     total := Sum(5, 5)
//     if total != 10 {
//        t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
//     }
// }