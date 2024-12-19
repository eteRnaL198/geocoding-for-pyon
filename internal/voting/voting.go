package excel

import (
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Address struct {
	RowIdx int
	Value  string
}

func LoadAddresses(filePath string) []Address {
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
			Value:  row[1],
		}
		addresses = append(addresses, address)
	}

	return addresses
}

type Coordinate struct {
	RowIdx int
	Lat    float64
	Long   float64
}

func WriteCoordinates(filePath string, Coordinates []Coordinate) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, c := range Coordinates {
		f.SetCellValue(f.GetSheetName(0), "C"+strconv.Itoa(c.RowIdx), c.Lat)
		f.SetCellValue(f.GetSheetName(0), "D"+strconv.Itoa(c.RowIdx), c.Long)

	}

	if err := f.Save(); err != nil {
		log.Fatal(err)
	}
}
