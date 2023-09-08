package main

import (
	"encoding/json"
	"finance/internal/types"
	"net/http"
)

func StockHandler(HostURL,AbuQuir string) error {
	
	companyStock := &types.Company{}
	resp, err := http.Get(HostURL + AbuQuir)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&companyStock); err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}
	companyStock.Manipulator()
	return nil
}
