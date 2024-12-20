package excel

import (
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Address struct {
	RowIdx int
	Value  string
}

func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func LoadAddresses(filePath string) []Address {
	log.Println("Starting to load addresses from file:", filePath)

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Get all the rows in the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		log.Fatal(err)
	}

	var addresses []Address
	for i, row := range rows[1:] { // Skip header row
		address := Address{
			RowIdx: i + 2, // starts with 2
			Value:  removeSpaces(row[0]),
		}
		addresses = append(addresses, address)
	}

	log.Println("Finished loading addresses from file:", filePath)
	return addresses
}
