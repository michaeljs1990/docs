## Stock Price

See the full problem [here](https://www.interviewcake.com/question/python/stock-price).
If you click through it will also show you a full solution and go into more depth than
I will on the problem. I have outlined the question below for posterity.

### Question

 First, I wanna know how much money I could have made yesterday if I'd been trading Apple 
 stocks all day.

So I grabbed Apple's stock prices from yesterday and put them in a list called stock_prices, where:

  - The indices are the time (in minutes) past trade opening time, which was 9:30am local time.
  - The values are the price (in US dollars) of one share of Apple stock at that time.

So if the stock cost $500 at 10:30am, that means stock_prices[60] = 500.

Write an efficient function that takes stock_prices and returns the best profit I could have
made from one purchase and one sale of one share of Apple stock yesterday.

For example:

```python
stock_prices = [10, 7, 5, 8, 11, 9]

get_max_profit(stock_prices)
# Returns 6 (buying for $5 and selling for $11)
```

No "shorting" - you need to buy before you can sell. Also, you can't buy and sell in the
same time step-at least 1 minute has to pass.

### Solution

The following was done in golang and the full program can be seen [here](). 

 - Once you find that by keeping track of the min and diff you can move over the arry in order. This removes any need to think about the indices and greatly simplifies the problem.
 - We want to always return an answer even if it's negative so we always can calculate the diff from the first two indices of the array safely.
 - I start ranging at [2:] just so we don't recompute the diff we did at the start.
 - tmpDiff could be moved inside the first if statement to make this more efficient but I think it also makes the code harder to read.

```go
package main

import (
  "fmt"
)

func solver(prices []int) int {
  diff := prices[1] - prices[0]
  min := prices[0]
  for _, price := range prices[2:] {
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
```

### Why this is a bad question

Overall I think this is a good question however the question can be reworded to not require
the person to think about share prices and stocks at all. Someone who is familiar with the
stock market works would likely be much more confident right away where someone who does not
know the market opens at 9:30 AM would likely have to read this more than once.

### Why this is a good question

This does not really rely on someone memorizing and algorithm that they may or may not have
gone over in prep for the interview. Additionally candidates can always start with a solution
that loops over everything and then while doing that most will get to the next step that lets
them solve it in O(n) time.