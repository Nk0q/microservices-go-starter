package types

type Route struct {
	Distance float64     `json:"distance"`
	Duration float64     `json:"duration"`
	Geometry []*Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates []*Coordinate `json:"coordinates"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OSRMApiResponse struct {
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Geometry struct {
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"routes"`
}

type TripPreviewResponse struct {
	Route     Route `json:"route"`
	RideFares any   `json:"rideFares"`
}

func (o *OSRMApiResponse) ToTripPreview() *TripPreviewResponse {
	if o == nil || len(o.Routes) == 0 {
		return &TripPreviewResponse{}
	}

	first := o.Routes[0]

	coords := make([]*Coordinate, 0, len(first.Geometry.Coordinates))
	for _, pair := range first.Geometry.Coordinates {
		if len(pair) < 2 {
			continue
		}
		coords = append(coords, &Coordinate{
			Latitude:  pair[1],
			Longitude: pair[0],
		})
	}

	return &TripPreviewResponse{
		Route: Route{
			Distance: first.Distance,
			Duration: first.Duration,
			Geometry: []*Geometry{
				{
					Coordinates: coords,
				},
			},
		},
		RideFares: []any{},
	}
}
