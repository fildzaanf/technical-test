package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/joho/godotenv"
)

type CountryCities struct {
	CountryName string `json:"country_name"`
	CityCount   int    `json:"city_count"`
}

func SaveToCSV(data []CountryCities, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error creating csv file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"CountryName", "CityCount"})

	for _, entry := range data {
		writer.Write([]string{entry.CountryName, strconv.Itoa(entry.CityCount)})
	}
}

func SaveToJSON(data []CountryCities, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error creating json file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file:", err)
		return
	}

	csvURL := os.Getenv("CSV_URL")
	if csvURL == "" {
		fmt.Println("CSV_URL is not set in .env file")
		return
	}

	response, err := http.Get(csvURL)
	if err != nil {
		fmt.Println("error downloading the csv file:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("failed to download file, status:", response.Status)
		return
	}

	reader := csv.NewReader(response.Body)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading csv file:", err)
		return
	}

	cityCountMap := make(map[string]int)

	for _, record := range records[1:] {
		country := record[7]
		cityCountMap[country]++
	}

	var countryCities []CountryCities
	for country, count := range cityCountMap {
		countryCities = append(countryCities, CountryCities{CountryName: country, CityCount: count})
	}

	sort.Slice(countryCities, func(i, j int) bool {
		return countryCities[i].CityCount < countryCities[j].CityCount
	})

	SaveToCSV(countryCities, "output.csv")

	SaveToJSON(countryCities, "output.json")

	fmt.Println("processing completed")
}
