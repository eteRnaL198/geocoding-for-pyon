package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coordinate struct {
	Lat  float64
	Long float64
}

type GeoRecord struct {
	Geometry struct {
		Coordinates [2]float64 `json:"coordinates"`
		Type        string     `json:"type"`
	} `json:"geometry"`
	Type       string `json:"type"`
	Properties struct {
		AddressCode string `json:"addressCode"`
		Title       string `json:"title"`
	} `json:"properties"`
}

func FetchGeoCoordinate(address string) (Coordinate, error) {
	url := fmt.Sprintf("https://msearch.gsi.go.jp/address-search/AddressSearch?q=%s", address)
	resp, err := http.Get(url)
	if err != nil {
		return Coordinate{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Coordinate{}, err
	}

	var records []GeoRecord
	if err := json.Unmarshal(body, &records); err != nil {
		return Coordinate{}, err
	}

	lat := records[0].Geometry.Coordinates[1]
	long := records[0].Geometry.Coordinates[0]
	return Coordinate{Lat: lat, Long: long}, nil
}
