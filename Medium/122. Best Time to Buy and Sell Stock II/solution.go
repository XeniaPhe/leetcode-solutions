package best_time_to_buy_and_sell_stock_II

func maxProfit(prices []int) int {
    priceBought, currProfit, totalProfit := prices[0], 0, 0
    for i := 1; i < len(prices); i += 1 {
        if priceBought, currProfit = prices[i], prices[i] - priceBought; currProfit > 0 {
            totalProfit += currProfit
        }
    }

    return totalProfit
}