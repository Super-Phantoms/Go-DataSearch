package main

import (
	"log"

	"smartdata/geolocation"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9997", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := geolocation.NewGeoLocationServiceClient(conn)
	response, err := c.GetGeoLocationData(context.Background(), &geolocation.Message{Body: "123 Main Street Louisville"})
	if err != nil {
		log.Fatalf("Error when retrieving GeoLocation Data: %s", err)
	}
	// display the response from the server
	log.Printf("Response from server: %s", response.Body)

}
