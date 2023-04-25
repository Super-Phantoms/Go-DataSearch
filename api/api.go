package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"smartdata/finance"
	"smartdata/geolocation"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our API for Smart Data!")
	fmt.Println("Endpoint: /")
}

func getGeoLocationData(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming API GeoLocation Request")
	vars := mux.Vars(r)
	address := vars["address"]
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9997", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Fprintln(w, "Error. Please try again later")
		return
	}
	c := geolocation.NewGeoLocationServiceClient(conn)
	response, err := c.GetGeoLocationData(context.Background(),
		&geolocation.Message{Body: address})
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, response.Body)
}
func getQuote(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming API Quote Request")
	vars := mux.Vars(r)
	symbol := vars["symbol"]
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9997", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		// issues with connecting to gRPC
		fmt.Fprintln(w, "Error. Please try again later") //provide standard
		// error messages instead of technical errors
		return
	}
	c := finance.NewFinanceServiceClient(conn)
	response, err := c.GetQuoteData(context.Background(), &finance.Message{Body: symbol})
	if err != nil {
		// issues with the symbol (symbol doesn't exist) or the Yahoo API is not working.
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, response.Body) //return the quote in the body of the response
}

// handleRequests will process HTTP requests and redirect them to
// the appropriate Handle function
func handleRequests() {
	// create a router to handle our requests from the mux package.
	router := mux.NewRouter().StrictSlash(true)
	// access root page
	router.HandleFunc("/", homePage)
	router.HandleFunc("/getGeoLocationData/{address}", getGeoLocationData)
	router.HandleFunc("/getQuote/{symbol}", getQuote)
	log.Fatal(http.ListenAndServe(":11112", router))
}
func main() {
	handleRequests()
}
