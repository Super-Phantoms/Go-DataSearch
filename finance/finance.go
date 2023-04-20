package finance

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedFinanceServiceServer
}

var YAHOO_API_KEY string = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
var URL = "https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols="
var STD_ERROR = errors.New("Unable to retrieve quote from the Yahoo API. Please check the symbol or try again later.")

func BuildYahooRequest(symbol string) (*http.Response, error) {
	req, err := http.NewRequest("GET", URL+symbol, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Api-Key", YAHOO_API_KEY)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (s *Server) GetQuoteData(ctx context.Context, in *Message) (*Message, error) {
	log.Println("Incoming Quote Request")
	symbol := in.GetBody()
	resp, err := BuildYahooRequest(symbol)
	if err != nil {
		return &Message{Body: "Error"}, STD_ERROR
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Message{Body: "Error"}, STD_ERROR
	}
	return &Message{Body: string(body)}, nil
}
