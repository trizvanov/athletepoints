// segment_example.go provides a simple example to fetch a segment details
// and list the top 10 on the leaderboard.
//
// usage:
//   > go get github.com/strava/go.strava
//   > cd $GOPATH/github.com/strava/go.strava/examples
//   > go run segment_example.go -id=segment_id -token=access_token
//
//   You can find an access_token for your app at https://www.strava.com/settings/api
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/strava/go.strava"
)

func main() {
	var athleteId int64
	var accessToken string

	// Provide an access token, with write permissions.
	// You'll need to complete the oauth flow to get one.
	flag.Int64Var(&athleteId, "id", 0, "Athlete Id")
	flag.StringVar(&accessToken, "token", "", "Access Token")

	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

		flag.PrintDefaults()
		os.Exit(1)
	}


	client := strava.NewClient(accessToken)
	athlete, err := strava.NewCurrentAthleteService(client).Get().Do()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	fmt.Printf("%s %s\n\n", athlete.FirstName, athlete.LastName)

	service := strava.NewCurrentAthleteService(client)

	friends, err := service.ListFriends().Page(1).PerPage(100).Do()

	for _, e := range friends {
		fmt.Printf("%s %s %s\n", e.FirstName, e.LastName, e.Friend)
	}
}
