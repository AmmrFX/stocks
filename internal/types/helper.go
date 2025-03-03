package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ------------------------------------------------------------------------------------------------------------------------------------

func parseChangePercentage(percentage string) (float64, error) {
	percentage = strings.TrimSuffix(percentage, "%")
	return strconv.ParseFloat(percentage, 64)
}

// ------------------------------------------------------------------------------------------------------------------------------------

func (comp *Company) Manipulator(lastPercent float64,company string) (string, float64) {
	initialCredit := 983.7
	nowPrice := comp.CompanyTab.PriceBar
	changePercentage, err := parseChangePercentage(nowPrice.ChangePercentage)

	if err != nil {
		fmt.Println("Invalid change percentage format:", err)
		return "", 0
	}

	nowPricefloat, _ := strconv.ParseFloat(nowPrice.Value, 64)
	profit := nowPricefloat*3 - initialCredit
	message := ""
	if changePercentage > 0.5 {
		for {
			InvestmentValue := profit + initialCredit

			if changePercentage-lastPercent > 0.5 {
				fmt.Printf("Chaannnge "+company+": Change percentage is %.2f%% %s, Total Investment %f", changePercentage, nowPrice.Status, InvestmentValue)
				message = fmt.Sprintf(" Chaannnge "+company+": Total Change percentage is %.2f%% %s, Total Investment %f", changePercentage, nowPrice.Status, InvestmentValue)
				message += fmt.Sprintf(" profit %f value %s", profit, nowPrice.Value)
				return message, lastPercent
			}
			fmt.Printf(""+company+": Total Change percentage is %.2f%% %s, Total Investment %f", changePercentage, nowPrice.Status, InvestmentValue)
			time.Sleep(5 * time.Minute)

		}
	}
	InvestmentValue := profit + initialCredit

	if changePercentage < -0.5 {
		message, lastPercent := belowPercentage()
	
		return message, lastPercent
	} else {
		fmt.Printf("Not sending notification:Loss "+company+": Total Change percentage is %.2f%% %s, Total Investment %f", changePercentage, nowPrice.Status, InvestmentValue)
	}
	return "", 0
}

func belowPercentage()()