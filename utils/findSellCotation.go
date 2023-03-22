package utils

func FindSellCotation(investedAmount, coinQuantity, fee, openPrice, expectedProfit float64) (closePrice float64) {
	var buyfeeAmount, sellFeeAmount, amountToClose, wantedProfit float64
	if fee > 0 {
		buyfeeAmount = investedAmount * fee // taxa q foi subtraida na compra
		wantedProfit = ((investedAmount + buyfeeAmount) * (1 + expectedProfit))
		sellFeeAmount = wantedProfit * fee             // taxa no momento da venda
		amountToClose = (wantedProfit + sellFeeAmount) // acrescentar taxa que virá
	} else {
		wantedProfit = (investedAmount * (1 + expectedProfit)) //profit desejado
		amountToClose = wantedProfit
	}

	if coinQuantity > 0 {
		closePrice = amountToClose / coinQuantity
	} else {
		closePrice = 0
	}

	return
}
