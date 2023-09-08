package types

import (
	"fmt"
	"strconv"
	"strings"
)

// ------------------------------------------------------------------------------------------------------------------------------------

func parseChangePercentage(percentage string) (float64, error) {
	percentage = strings.TrimSuffix(percentage, "%")
	return strconv.ParseFloat(percentage, 64)
}

// ------------------------------------------------------------------------------------------------------------------------------------

func (comp *Company) Manipulator() {
	
	nowPrice := comp.CompanyTab.PriceBar
	changePercentage, err := parseChangePercentage(nowPrice.ChangePercentage)
	if err != nil {
		fmt.Println("Invalid change percentage format:", err)
		return
	}

	if changePercentage > 1.0 {
		fmt.Printf("Sending notification: Change percentage is %.2f%% %s\n", changePercentage, nowPrice.Status)
	} else {
		fmt.Printf(" not Sending notification: Change percentage is %.2f%% %s\n", changePercentage, nowPrice.Status)
	}
}

