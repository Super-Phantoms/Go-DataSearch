package geolocation

import (
	context "context"
	"encoding/json"
	"errors"
	"log"

	"googlemaps.github.io/maps"
)

type Server struct {
	UnimplementedGeoLocationServiceServer
}

var GOOGLE_API_KEY string = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
var STD_ERROR = errors.New("Unable to retrieve GeoLocation Data from the Google Maps API. Please check the address or try again later.")

func (s *Server) GetGeoLocationData(ctx context.Context, in *Message) (*Message, error) {
	log.Println("Incoming GeoLocation Request")
	c, err := maps.NewClient(maps.WithAPIKey(GOOGLE_API_KEY))
	if err != nil {
		return &Message{Body: "Error"}, STD_ERROR
	}
	r := &maps.GeocodingRequest{
		Address: in.GetBody(), // you retrieve the address from the
		// Message in body
	}
	geocode, err := c.Geocode(context.Background(), r)
	if err != nil {
		return &Message{Body: "Error"}, STD_ERROR
	}
	if len(geocode) < 1 { // the geocode result returned is empty due
		// to erroneous address
		return &Message{Body: "Error"}, STD_ERROR
	}
	geocodeJson, err := json.Marshal(geocode) // convert the geocode result to
	// json string so you can return it.
	if err != nil {
		return &Message{Body: "Error"}, STD_ERROR
	}
	return &Message{Body: string(geocodeJson)}, nil
}
