package excel

import (
	"log"
	"strconv"

	"github.com/eteRnaL198/geocoding-for-pyon/internal/geo"
	"github.com/xuri/excelize/v2"
)

type Coordinate struct {
	RowIdx int
	Lat    float64
	Long   float64
}

func ToCoordinate(idx int, c geo.Coordinate) Coordinate {
	return Coordinate{RowIdx: idx, Lat: c.Lat, Long: c.Long}
}

type Coordinates []Coordinate

func NewCoordinates(l int) Coordinates {
	return make(Coordinates, 0, l)
}

func (c *Coordinates) Append(idx int, gc geo.Coordinate) {
	*c = append(*c, ToCoordinate(idx, gc))
}

func (c *Coordinates) Write(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, c := range *c {
		f.SetCellValue(f.GetSheetName(0), "C"+strconv.Itoa(c.RowIdx), c.Lat)
		f.SetCellValue(f.GetSheetName(0), "D"+strconv.Itoa(c.RowIdx), c.Long)
	}

	if err := f.Save(); err != nil {
		log.Fatal(err)
	}
}
