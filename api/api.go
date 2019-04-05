package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	// AuthCode Spotify authorization code for getting access tokens.  Returned from /authorize spotify endpoint
	AuthCode = ""
	// scopes restriction of spotify actions.  We only need read and modify current playback state and queue
	scopes  = "user-read-playback-state user-modify-playback-state user-read-playback-state playlist-modify-public user-read-private"
	sAcsTok spotAccessTokenResp
	sTracks spotTrackSearchResponse
	// partyName will be used for creating the playlist
	partyName        = time.Now().Format("2006-01-02 3:4:5")
	sCurrentUser     spotUser
	sCurrentPlaylist spotPlaylist
	sPlayerContext   string
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

// spotify user
type spotUser struct {
	DisplayName string `json:"display_name"`
	UserID      string `json:"id"`
	// AccountLevel is the subscription level, "premium", "free", etc.
	AccountLevel string `json:"product"`
}

// spotify playlist
type spotPlaylist struct {
	PlaylistName string `json:"name"`
	PlaylistURI  string `json:"uri"`
	PlaylistID   string `json:"id"`
}

// search Track structure
type spotTrackSearchResponse struct {
	Tracks spotTracks `json:"tracks"`
}
type spotTracks struct {
	Items []spotTrack `json:"items"`
}
type spotTrack struct {
	TrackName string       `json:"name"`
	Artists   []spotArtist `json:"artists"`
	Explicit  bool         `json:"explicit"`
	TrackID   string       `json:"id"`
	TrackURI  string       `json:"uri"`
}
type spotArtist struct {
	ArtistID   string `json:"id"`
	ArtistName string `json:"name"`
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

// SPOTIFY RELATED ROUTES

// LinkSpotify Login/Authenticates Soundsync with spotify by asking user to sign in
// user must have premium or student spotify account spotify account
// Must have SPOTIFY_ID and SPOTIFY_SECRET environment variables
func LinkSpotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /LinkSpotify")
	var URL *url.URL
	URL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		panic(err)
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
	// IMPORTANT
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
	getCurrentUser()
	createPlaylist()
}

func getSpotToken() {
	var URL *url.URL
	URL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		panic(err)
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
	fmt.Println("Access Tokens:")
	fmt.Printf("\t:Access Token: %s\n", sAcsTok.AccessToken)
	fmt.Printf("\t:Refresh Token: %s\n", sAcsTok.RefreshToken)
	fmt.Printf("\t:Expires In: %s\n", strconv.Itoa(sAcsTok.ExpiresIn))
}

func refreshAccessToken() {
	var URL *url.URL
	URL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		panic(err)
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

	fmt.Println()
	fmt.Println("Refresh")
	fmt.Println(sAcsTok.AccessToken)
	fmt.Println(sAcsTok.RefreshToken)
	fmt.Println(sAcsTok.ExpiresIn)
}

// Play handles play media control
func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /MediaControls")
	action := strings.TrimPrefix(r.URL.Path, "/MediaControls/")
	fmt.Println("Action:", action)

	var URL *url.URL
	URL, err := url.Parse("https://api.spotify.com")
	if err != nil {
		panic(err)
	}
	URL.Path += "/v1/me/player/play"

	client := &http.Client{}
	req := &http.Request{}

	if sPlayerContext != sCurrentPlaylist.PlaylistURI {
		reqBody := map[string]interface{}{
			"context_uri": sCurrentPlaylist.PlaylistURI,
		}
		bytesRepresentation, err := json.Marshal(reqBody)
		if err != nil {
			log.Fatalln(err)
		}
		req, err = http.NewRequest("PUT", URL.String(), bytes.NewBuffer(bytesRepresentation))
		sPlayerContext = sCurrentPlaylist.PlaylistURI
	} else {
		req, err = http.NewRequest("PUT", URL.String(), nil)
	}

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
	URL.Path += "/v1/me/player/pause"

	client := &http.Client{}

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

// SearchSpotify searches spotify by the track name
func SearchSpotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /SearchSpotify")
	params := mux.Vars(r)
	query := params["query"]
	fmt.Println("Query: " + query)

	var URL *url.URL
	URL, err := url.Parse("https://api.spotify.com")
	if err != nil {
		panic(err)
	}
	URL.Path += "/v1/search"
	parameters := url.Values{}
	parameters.Add("q", query)
	parameters.Add("type", "track")
	// limit the number of results.  Default: 20
	parameters.Add("limit", "20")
	URL.RawQuery = parameters.Encode()

	// get html for user to sign in to
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sAcsTok.AccessToken)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &sTracks)
	if err != nil {
		panic(err)
	}

	// Search results for first track in list and first artist linked to that track
	fmt.Println()
	fmt.Println("Search Information:")
	fmt.Printf("\t:Track Name: %s\n", sTracks.Tracks.Items[0].TrackName)
	fmt.Printf("\t:Track ID: %s\n", sTracks.Tracks.Items[0].TrackID)
	fmt.Printf("\t:Is Explicit: %s\n", strconv.FormatBool(sTracks.Tracks.Items[0].Explicit))
	fmt.Printf("\t:Artist Name: %s\n", sTracks.Tracks.Items[0].Artists[0].ArtistName)
	fmt.Printf("\t:Artist ID: %s\n", sTracks.Tracks.Items[0].Artists[0].ArtistID)
}

// AddSong adds a song to the queue ((playlist))
func AddSong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request: /AddSong")
	params := mux.Vars(r)
	encodedTrackURI := params["songURI"]
	trackURI, err := url.QueryUnescape(encodedTrackURI)
	if err != nil {
		panic(err)
	}
	fmt.Println("Track URI: " + trackURI)

	var URL *url.URL
	URL, err = url.Parse("https://api.spotify.com")
	if err != nil {
		panic(err)
	}
	URL.Path += "v1/playlists/" + sCurrentPlaylist.PlaylistID + "/tracks"

	reqBody := make(map[string][]string)
	reqBody["uris"] = append(reqBody["uris"], trackURI)
	reqBody["uris"] = append(reqBody["uris"], "spotify:track:3e9HZxeyfWwjeyPAMmWSSQ")
	reqBody["uris"] = append(reqBody["uris"], "spotify:track:5uIRujGRZv5t4fGKkUTv4n")
	reqBody["uris"] = append(reqBody["uris"], "spotify:track:6zWU7YALeEDMcPGhKKZJhV")
	reqBody["uris"] = append(reqBody["uris"], "spotify:track:6XQHlsNu6so4PdglFkJQRJ")
	bytesRepresentation, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sAcsTok.AccessToken)
	if err != nil {
		panic(err)
	}

	// resp body contains playlist "snapshot_id"
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

// gets the currently logged in user
func getCurrentUser() {
	var URL *url.URL
	URL, err := url.Parse("https://api.spotify.com")
	if err != nil {
		panic(err)
	}
	URL.Path += "v1/me"

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sAcsTok.AccessToken)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &sCurrentUser)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("User Information:")
	fmt.Printf("\t:User Name: %s\n", sCurrentUser.DisplayName)
	fmt.Printf("\t:User ID: %s\n", sCurrentUser.UserID)
	fmt.Printf("\t:User account type: %s\n", sCurrentUser.AccountLevel)

	defer resp.Body.Close()
}

// creates a playlist on the currently logged in users (hosts) account
func createPlaylist() {
	var URL *url.URL
	URL, err := url.Parse("https://api.spotify.com")
	if err != nil {
		panic(err)
	}
	URL.Path += "v1/users/" + sCurrentUser.UserID + "/playlists"

	reqBody := map[string]interface{}{
		"name":        partyName,
		"description": "My party created by SoundSync",
	}

	bytesRepresentation, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sAcsTok.AccessToken)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &sCurrentPlaylist)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Playlist Information:")
	fmt.Printf("\t:Playlist Name: %s\n", sCurrentPlaylist.PlaylistName)
	fmt.Printf("\t:Playlist URI: %s\n", sCurrentPlaylist.PlaylistURI)

	defer resp.Body.Close()
}

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
