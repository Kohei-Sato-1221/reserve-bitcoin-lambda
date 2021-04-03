package utils

import "math"

func RoundDecimal(num float64) float64 {
	return math.Round(num)
}

//少数の切り上げを行う
func roundUp(num, places float64) float64 {
	shift := math.Pow(10, places) //10のplaces乗行う
	return RoundDecimal(num * shift) / shift
}

//予算を元に購入数量を算出する
// BTC: 0.001BTC が　最小注文数量
func CalcAmount(price, budget, minimumAmount, places float64) float64 {
	amount := roundUp(budget/price, places)
	if amount < minimumAmount {
		return  minimumAmount
	} else {
		return amount
	}
}



