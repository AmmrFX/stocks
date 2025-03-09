package handlers

// func StockHandler(HostURL, stock_url, company string, lastPercent float64) (string, float64, error) {

// 	companyStock := &types.Company{}
// 	resp, err := http.Get(HostURL + stock_url)
// 	if err != nil {
// 		return "", 0, err
// 	}

// 	decoder := json.NewDecoder(resp.Body)
// 	if err := decoder.Decode(&companyStock); err != nil {
// 		return "", 0, err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", 0, nil
// 	}
// 	message, lastperc := companyStock.Manipulator(lastPercent, company)
// 	return message, lastperc, nil
// }

// // // ------------------------------------------------------------------------------------------------------------------------------------
