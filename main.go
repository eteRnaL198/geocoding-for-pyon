package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eteRnaL198/geocoding-for-pyon/internal/excel"
	"github.com/eteRnaL198/geocoding-for-pyon/internal/geo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the Excel file.")
		return
	}

	excelFilePath := os.Args[1]
	addresses := excel.LoadAddresses(excelFilePath)

	for i := 0; i < len(addresses); i += 10 {
		end := i + 10
		if end > len(addresses) {
			end = len(addresses)
		}
		batch := addresses[i:end]
		coordinates := excel.NewCoordinates(len(batch))

		for _, address := range batch {
			coordinate, err := geo.FetchGeoCoordinate(address.Value)
			if err != nil {
				fmt.Println(err)
				return
			}
			coordinates.Append(address.RowIdx, coordinate)
		}

		// Add a delay to respect the rate limit of 10 requests per second
		time.Sleep(time.Second)
		coordinates.Write(excelFilePath)
		fmt.Printf("Processed %d / %d\n", end, len(addresses))
	}
}
