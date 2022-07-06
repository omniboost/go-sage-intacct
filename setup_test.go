package intacct_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

var (
	client *intacct.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	senderID := os.Getenv("SENDER_ID")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	userID := os.Getenv("USER_ID")
	userPassword := os.Getenv("USER_PASSWORD")
	companyID := os.Getenv("COMPANY_ID")
	locationID := os.Getenv("LOCATION_ID")
	debug := os.Getenv("DEBUG")

	client = intacct.NewClient(nil, senderID, senderPassword, userID, userPassword, companyID)
	client.SetLocationID(locationID)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
