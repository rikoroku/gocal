package googlecalendar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const tokFile = "token.json"

func Setup() {
	tok, _ := tokenFromTokFile()
	if tok != nil {
		log.Fatalf("You have already done. Setup is first time only.")
	}

	config := getConfig()

	tok = getTokenFromWeb(config)
	saveToken(tokFile, tok)
}

func getConfig() *oauth2.Config {
	credentialPath := os.Getenv("CLIENT_SECRET_FILE")
	b, err := ioutil.ReadFile(credentialPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config
}

func getService() *calendar.Service {
	tok, err := tokenFromTokFile()
	if err != nil {
		log.Fatalf("Couldn't get a refresh token. First, You must do setup command.")
	}

	config := getConfig()
	client := config.Client(context.Background(), tok)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return srv
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromTokFile() (*oauth2.Token, error) {
	f, err := os.Open(tokFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
