package main

import "fmt"

type Exporter interface {
	Export() string
}

type SeaShipment struct {
	TrackingID string
}

type AirShipment struct {
	TrackingID string
}

func (s SeaShipment) Export() string {
	return fmt.Sprintf("Sea Shipment - Tracking ID: %s", s.TrackingID)
}

func (s AirShipment) Export() string {
	return fmt.Sprintf("Air Shipment - Tracking ID: %s", s.TrackingID)
}

func ProcessExports(items []Exporter) {
	for _, item := range items {
		fmt.Println(item.Export())
	}
}

func main() {
	sea := SeaShipment{TrackingID: "SEA12345"}
	air := AirShipment{TrackingID: "AIR67890"}
	exports := []Exporter{sea, air}
	ProcessExports(exports)
}
