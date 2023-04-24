package main

import (
	"log"
	"smartdata/finance"

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
	c := finance.NewFinanceServiceClient(conn)
	response, err := c.GetQuoteData(context.Background(), &finance.
		Message{Body: "AAPL"})
	if err != nil {
		log.Fatalf("Error when retrieving the Quote Data: %s", err)
	}
	// display the response from the server
	log.Printf("Response from server: %s", response.Body)
}
