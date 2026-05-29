func max(x , y int)int{
	if x>y{
		return x
	}

	return y
}

func min (x , y int) int {
	if x<y{
		return x
	}

	return y
}

func maxProfit(prices []int) int {
	maxProfit :=0

	buyDay := prices[0]

	for i:=1; i<len(prices); i++{

		if prices[i] - buyDay >0{
			maxProfit = max(maxProfit , prices[i]-buyDay)
		}

		buyDay = min(buyDay  , prices[i])
	}

	return maxProfit
}
