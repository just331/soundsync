package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joshuaj1397/soundsync/model"
)

const redirectURI = "http://localhost:3005/callback"

var (
	mySigningKey = []byte("ASuperSecretSigningKeyCreatedByTheAliensFromArrival")
	// State spotify session state, should be unique for each party instance i.e. not static
	State = "mySuperCoolState"
	// myClientID spotify client id environment variable
	myClientID = os.Getenv("SPOTIFY_ID")
	// mySecretShh spotify client secret environment variable
	mySecretShh = os.Getenv("SPOTIFY_SECRET")
	// scopes restriction of spotify actions.  We only need read and modify current playback state and queue
	scopes = "user-read-playback-state user-modify-playback-state"
	// AuthCode Spotify authorization code for getting access tokens.  Returned from /authorize spotify endpoint
	AuthCode = ""
	sAcsTok  spotAccessTokenResp
)

var html = `
<br/>
<a href="/MediaControls/Play">Play</a><br/>
<a href="/MediaControls/Pause">Pause</a><br/>
<a href="/MediaControls/Next">Next track</a><br/>
<a href="/MediaControls/Previous">Previous Track</a><br/>
`

type spotAccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

var GetToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "Joshua Johnson"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)
	fmt.Println("I singed a token senpai!")

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
})

// CreateParty returns the party code so the host can send it out to others
var CreateParty = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	phoneNum := params["phoneNum"]
	nickname := params["nickname"]
	partyName := params["partyName"]

	partyCode, err := model.CreateParty(partyName, phoneNum, nickname)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(partyCode)
})

var JoinParty = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	partyCode := params["partyCode"]
	nickname := params["nickname"]
	phoneNum := params["phoneNum"]
	err := model.JoinParty(partyCode, nickname, phoneNum)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Party joined")
})

// func Verify(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	phoneNum := params["phoneNum"]
// 	name := params["name"]
// 	authCode := params["authCode"]
// 	status, err := model.Verify(phoneNum, name, authCode)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	json.NewEncoder(w).Encode(status)
// }

// Spotify related routes

// LinkSpotify Login/Authenticates Soundsync with spotify by asking user to sign in
// user must have premium or student spotify account spotify account
// Must have SPOTIFY_ID and SPOTIFY_SECRET environment variables
func LinkSpotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /LinkSpotify")
	var URL *url.URL
	URL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		panic("boom")
	}
	URL.Path += "/authorize"
	parameters := url.Values{}
	parameters.Add("client_id", myClientID)
	parameters.Add("response_type", "code")
	parameters.Add("redirect_uri", redirectURI)
	parameters.Add("state", State)
	parameters.Add("scope", scopes)
	URL.RawQuery = parameters.Encode()

	fmt.Println("Visit this to log into spotify:")
	fmt.Println(URL.String())

	// get html for user to sign in to
	resp, err := http.Get(URL.String())
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println()
	log.Println()
	// display this to the user.  it is where they will sign in to spotify
	// this is just html, we could just inject this into the frontend
	log.Println(string(body))
	log.Println()
	log.Println()
}

// SpotifyCallback Redirected here after authorization to receive authCode
// authCode can be exchanged for an access token
func SpotifyCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /SpotifyCallback")
	urlParams := r.URL.Query()
	AuthCode = urlParams.Get("code")

	getSpotToken()
}

func getSpotToken() {
	var URL *url.URL
	URL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		panic("boom")
	}
	URL.Path += "/api/token"

	// body of post request
	formData := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {AuthCode},
		"redirect_uri":  {redirectURI},
		"client_id":     {myClientID},
		"client_secret": {mySecretShh},
	}

	resp, err := http.PostForm(URL.String(), formData)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// read the payload
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// pass a pointer for the response type
	err = json.Unmarshal(body, &sAcsTok)

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Access")
	fmt.Println(sAcsTok.AccessToken)
	fmt.Println(sAcsTok.RefreshToken)
	fmt.Println(sAcsTok.ExpiresIn)
	fmt.Println()
}

func refreshAccessToken() {
	var URL *url.URL
	URL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		panic("boom")
	}
	URL.Path += "/api/token"

	// body of post request
	formData := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {sAcsTok.RefreshToken},
		"client_id":     {myClientID},
		"client_secret": {mySecretShh},
	}

	response, err := http.PostForm(URL.String(), formData)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// read the payload
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// pass a pointer for the response type
	err = json.Unmarshal(body, &sAcsTok)

	if err != nil {
		panic(err)
	}

	fmt.Println("Refresh")
	fmt.Println(sAcsTok.AccessToken)
	fmt.Println(sAcsTok.RefreshToken)
	fmt.Println(sAcsTok.ExpiresIn)
	fmt.Println()
}

// PlayPause handels play and pause controls
func PlayPause(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /MediaControls")
	action := strings.TrimPrefix(r.URL.Path, "/MediaControls/")
	fmt.Println("Action:", action)

	var URL *url.URL
	URL, err := url.Parse("https://api.spotify.com")
	if err != nil {
		panic(err)
	}
	URL.Path += "/v1/me/player"

	client := &http.Client{}

	switch action {
	case "Play":
		URL.Path += "/play"
	case "Pause":
		URL.Path += "/pause"
	}

	req, err := http.NewRequest("PUT", URL.String(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sAcsTok.AccessToken)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

}

// NextPrev handels next and previous controls
func NextPrev(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /MediaControls")
	action := strings.TrimPrefix(r.URL.Path, "/MediaControls/")
	fmt.Println("Action:", action)

	var URL *url.URL
	URL, err := url.Parse("https://api.spotify.com")
	if err != nil {
		panic(err)
	}
	URL.Path += "/v1/me/player"

	client := &http.Client{}

	switch action {
	case "Next":
		URL.Path += "/next"
	case "Previous":
		URL.Path += "/previous"
	}

	req, err := http.NewRequest("POST", URL.String(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sAcsTok.AccessToken)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

}

// SearchSpotify
// func SearchSpotify(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	query := params["query"]

// 	config := &clientcredentials.Config{
// 		ClientID:     os.Getenv("SPOTIFY_ID"),
// 		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
// 		TokenURL:     spotify.TokenURL,
// 	}
// // 	token, err := config.Token(context.Background())
// 	if err != nil {
// 		log.Fatalf("couldn't get token: %v", err)
// 	}

// 	client := spotify.Authenticator{}.NewClient(token)
// 	// search for playlists and albums containing query
// 	results, err := client.Search(query, spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// handle album results
// 	if results.Albums != nil {
// 		fmt.Println("Albums:")
// 		for _, item := range results.Albums.Albums {
// 			fmt.Println("   ", item.Name)
// 		}
// 	}
// 	// handle playlist results
// 	if results.Playlists != nil {
// 		fmt.Println("Playlists:")
// 		for _, item := range results.Playlists.Playlists {
// 			fmt.Println("   ", item.Name)
// 		}
// 	}
// }

//
// func AddSong(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	songId := params["songId"]
// 	partyId := params["partyId"]
// 	status, err := model.AddSong(songId, partyId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(status)
// }
//
// func SongQueue(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	partyId := params["partyId"]
// 	songQueue, err := model.SongQueue(partyId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(songQueue)
// }
//
// func RemoveSong(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	songId := params["songId"]
// 	partyId := params["partyId"]
// 	status, err := model.RemoveSong(songId, partyId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(status)
// }
