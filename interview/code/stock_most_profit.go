// https://www.interviewcake.com/question/python/stock-price
//
// First, I wanna know how much money I could have made yesterday if I'd been trading Apple stocks all day.
// 
// So I grabbed Apple's stock prices from yesterday and put them in a list called stock_prices, where:
// 
//  - The indices are the time (in minutes) past trade opening time, which was 9:30am local time.
//  - The values are the price (in US dollars) of one share of Apple stock at that time.
// 
// So if the stock cost $500 at 10:30am, that means stock_prices[60] = 500.
// 
// Write an efficient function that takes stock_prices and returns the best profit I could have made 
// from one purchase and one sale of one share of Apple stock yesterday. 

package main

import (
  "fmt"
)

func solver(prices []int) int {
  diff := prices[1] - prices[0]
  min := prices[0]
  for _, price := range prices[1:] {
    tmpDiff := price - min
    if price > min && tmpDiff > diff {
      diff = tmpDiff
    }

    if price < min {
      min = price
    }
  }

  return diff
}

func main() {
  sharePrices := []int{10, 7, 5, 8, 11, 9}
  fmt.Println("(ans: 6) \t out: ", solver(sharePrices))

  sharePrices = []int{10, 8, 6, 4, 1}
  fmt.Println("(ans: -2) \t out: ", solver(sharePrices))

  sharePrices = []int{10, 10, 10, 10, 10, 10}
  fmt.Println("(ans: 0) \t out: ", solver(sharePrices))
}
