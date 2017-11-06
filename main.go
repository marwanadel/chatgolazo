package main

import (

    "fmt"
    "net/http"
    "encoding/json"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"time"
	"strings"
	"strconv"
	"bytes"
	"os"

)

func main() {

    route()
    http.ListenAndServe(os.Getenv("PORT"), nil)

}


//apiFeeds.go

func getTeamProfile(w http.ResponseWriter, r *http.Request, b []byte, code int) {

	req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/teams/" + strconv.Itoa(code), nil)


	if er != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()
	
	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}	

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	var profile teamProfile
	err = json.Unmarshal(b, &profile)
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := message{}

	msg.Message = stringfyTeamProfile(profile)

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
	
}

func getTeamMatches(w http.ResponseWriter, r *http.Request, b []byte, code int) {
	
	req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/matches", nil)

	if er != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()
	
	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	q.Add("team_id", strconv.Itoa(code))
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	var fix []fixture
	err = json.Unmarshal(b, &fix)
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := message{}

	msg.Message = stringfyTeamFixtures(fix)

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func getLeagueTable(w http.ResponseWriter, r *http.Request, b []byte, code int) {
	
	req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/league-tables", nil)

	if er != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()

	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	q.Add("competition_id", strconv.Itoa(code))
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	
	var table []table 
	err = json.Unmarshal(b, &table)
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := message{}

	if(len(table) == 0){

		msg.Message = "No table currently available!"

	}else{

		msg.Message = stringfyTable(table[0])

	}

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}	

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func getLiveScores(w http.ResponseWriter, r *http.Request, b []byte, code int) {

    req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/matches", nil)

    if er != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()

	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	q.Add("competition_id", strconv.Itoa(code))

	timein := time.Now().Local().Add(time.Hour * time.Duration(-3) +
                                 time.Minute * time.Duration(0) +
                                 time.Second * time.Duration(0))



	q.Add("from", timein.Format("2006-01-02T15:04:05-07:00"))
								 
	q.Add("to", time.Now().Format("2006-01-02T15:04:05-07:00"))
	
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	var fix []fixture
	err = json.Unmarshal(b, &fix)

	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := message{}

	msg.Message = stringfyLiveFixtures(fix)

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

//handlers.go


var sessions []string


func home(w http.ResponseWriter, r *http.Request) {

	//redicrect to /welcome
  
	http.Redirect(w, r, "/welcome", 200)

}

func welcome(w http.ResponseWriter, r *http.Request) {

	//reply with uuid to be used in subsequent requests and responses

	// Generate a UUID.
	hasher := md5.New()
	hasher.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	id := hex.EncodeToString(hasher.Sum(nil))

	// Create a session for this UUID
	sessions = append(sessions, id)

    m := welcomeMessage{Message: "Welcome", Uuid: id}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)

}

func findSession(id string) bool{
	
	for i := 0; i < len(sessions); i++ {
			
		if(strings.Compare(id, sessions[i]) == 0){

			return true

		}

	}

	return false

}

func chat(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
	
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	
	}

	uuid := r.Header.Get("Authorization")
	
	if uuid == "" {
	
		http.Error(w, "Missing or empty Authorization header.", http.StatusUnauthorized)
		return
	
	}

	// Make sure a session exists for the extracted UUID
	sessionFound := findSession(uuid)
	
	if !sessionFound {
	
		http.Error(w, fmt.Sprintf("No session found for: %v.", uuid), http.StatusUnauthorized)
		return
	
	}

    b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
	
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return

	}

	// Unmarshal
	msg := message{}
	err = json.Unmarshal(b, &msg)
	
	if err != nil {
		
		msg := message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	filter, code := mappings(msg.Message)

	

	if(strings.Compare(filter, "") == 0){

		msg := message{}

		msg.Message = "Oops! We could not get that. Can you try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return

	}else{

		if(strings.Compare(filter, "team_id") == 0){

			if(code != 0){

				getTeamMatches(w, r, b, code)			
			
			}else{

				msg := message{}

				msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

				output, err := json.Marshal(msg)
				
				if err != nil {
				
					http.Error(w, err.Error(), 500)
					return
				
				}

				w.Header().Set("content-type", "application/json")
				w.Write(output)	
				
				return

			}

		}else{

			if(strings.Compare(filter, "competition_id") == 0){

				if(code != 0){

					getLeagueTable(w, r, b, code)
					
				}else{

					msg := message{}

					msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

					output, err := json.Marshal(msg)
					
					if err != nil {
					
						http.Error(w, err.Error(), 500)
						return
					
					}

					w.Header().Set("content-type", "application/json")
					w.Write(output)	
					
					return

				}
										

			}else{

				if(strings.Compare(filter, "team_profile") == 0){

					if(code != 0){

						getTeamProfile(w, r, b, code)
						
					}else{

						msg := message{}

						msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

						output, err := json.Marshal(msg)
						
						if err != nil {
						
							http.Error(w, err.Error(), 500)
							return
						
						}

						w.Header().Set("content-type", "application/json")
						w.Write(output)	
						
						return

					}
						
					
				}else{

					if(strings.Compare(filter, "live") == 0){

						if(code == 0){

							getLiveScores(w, r, b, 2)
						
						}else{

							msg := message{}

							msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

							output, err := json.Marshal(msg)
							
							if err != nil {
							
								http.Error(w, err.Error(), 500)
								return
							
							}

							w.Header().Set("content-type", "application/json")
							w.Write(output)	
							
							return

						}
						
					
					}else{

						msg := message{}

						msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

						output, err := json.Marshal(msg)
						
						if err != nil {
						
							http.Error(w, err.Error(), 500)
							return
						
						}

						w.Header().Set("content-type", "application/json")
						w.Write(output)	
						
						return

					}

				}

			}

		}

	}

}


//mappings.go

func mappings(s string) (string, int){

	s = strings.ToLower(s)

	switch s{

		case "premier league" , "pl": return "competition_id", 2
		
		//team matches
		case "liverpool", "reds": return "team_id", 1
		case "burnley": return "team_id", 480
		case "leicester city": return "team_id", 481
		case "brighton & hove albion", "hove albion": return "team_id", 482
		case "watford": return "team_id", 483
		case "southampton": return "team_id", 69
		case "chelsea", "blues": return "team_id", 7
		case "stoke city", "stoke": return "team_id", 8
		case "west ham united", "west ham", "westham": return "team_id", 202 
		case "manchester united", "manu", "man u", "manutd", "man utd": return "team_id", 2
		case "everton", "toffees": return "team_id", 12
		case "huddersfield town", "huddersfield": return "team_id", 557
		case "bournemouth": return "team_id", 558 
		case "swansea city", "swansea": return "team_id", 15
		case "crystal palace": return "team_id", 337
		case "arsenal", "afc", "gunners", "4th": return "team_id", 18 
		case "newcastle united", "newcastle": return "team_id", 19
		case "manchester city", "man city", "citizens", "city": return "team_id", 20
		case "west bromwich albion", "west brom": return "team_id", 14
		case "tottenham hotspur", "spurs", "tottenham", "they suck": return "team_id", 13

		//team profiles
		case "liverpool profile", "reds profile": return "team_profile", 1
		case "burnley profile": return "team_profile", 480
		case "leicester city profile": return "team_profile", 481
		case "brighton & hove albion profile", "hove albion profile": return "team_profile", 482
		case "watford profile": return "team_profile", 483
		case "southampton profile": return "team_profile", 69
		case "chelsea profile", "blues profile": return "team_profile", 7
		case "stoke city profile", "stoke profile": return "team_profile", 8
		case "west ham united profile", "west ham profile", "westham profile": return "team_profile", 202 
		case "manchester united profile", "manu profile", "man u profile", "manutd profile", "man utd profile": return "team_profile", 2
		case "everton profile", "toffees profile": return "team_profile", 12
		case "huddersfield town profile", "huddersfield profile": return "team_profile", 557
		case "bournemouth profile": return "team_profile", 558 
		case "swansea city profile", "swansea profile": return "team_profile", 15
		case "crystal palace profile": return "team_profile", 337
		case "arsenal profile", "afc profile", "gunners profile", "4th profile": return "team_profile", 18 
		case "newcastle united profile", "newcastle profile": return "team_profile", 19
		case "manchester city profile", "man city profile", "citizens profile", "city profile": return "team_profile", 20
		case "west bromwich albion profile", "west brom profile": return "team_profile", 14
		case "tottenham hotspur profile", "spurs profile", "tottenham profile", "they suck profile": return "team_profile", 13


		case "live scores", "scores": return "live", 0

		default: return "",0

	}
	
}

//routes

func route(){

	http.HandleFunc("/", home)

	http.HandleFunc("/welcome", welcome)

	http.HandleFunc("/chat", chat)

}

//Fixture.go

type fixture struct {

    HomeTeam team
    AwayTeam team
    AwayGoals int
    HomeGoals int
    Start float64
    CurrentState int

}

type team struct {

    Name string

}

func stringfyTeamFixtures(fix []fixture) string{

	var message bytes.Buffer 

	var fixture bytes.Buffer

	if(len(fix) == 0){

		return "No Recent Matches!"

	}

	for i := 0; i < len(fix); i++ {
		
		homeTeam := fix[i].HomeTeam
		awayTeam := fix[i].AwayTeam

		if(fix[i].CurrentState == 0){

			fixture.WriteString(homeTeam.Name)
			fixture.WriteString(" - ")
			fixture.WriteString(awayTeam.Name)
			fixture.WriteString("\n")
				
		}else{

			if(fix[i].CurrentState == 1){

				fixture.WriteString(homeTeam.Name)
				fixture.WriteString(" ")
				fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
				fixture.WriteString(" - ")
				fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
				fixture.WriteString(" ")
				fixture.WriteString(awayTeam.Name)
				fixture.WriteString("	1st. Half")
				fixture.WriteString("\n")

			}else{

				if(fix[i].CurrentState == 2){
				
					fixture.WriteString(homeTeam.Name)
					fixture.WriteString(" ")
					fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
					fixture.WriteString(" - ")
					fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
					fixture.WriteString(" ")
					fixture.WriteString(awayTeam.Name)
					fixture.WriteString("	HT")
					fixture.WriteString("\n")

				}else{

					if(fix[i].CurrentState == 3){

						fixture.WriteString(homeTeam.Name)
						fixture.WriteString(" ")
						fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
						fixture.WriteString(" - ")
						fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
						fixture.WriteString(" ")
						fixture.WriteString(awayTeam.Name)
						fixture.WriteString("	2nd. Half")
						fixture.WriteString("\n")

					}else{

						fixture.WriteString(homeTeam.Name)
						fixture.WriteString(" ")
						fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
						fixture.WriteString(" - ")
						fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
						fixture.WriteString(" ")
						fixture.WriteString(awayTeam.Name)
						fixture.WriteString("	FT")
						fixture.WriteString("\n")

					}

				}

			}

		}	

		message.WriteString(fixture.String()) 

	}

	return message.String()
	
}

func stringfyLiveFixtures(fix []fixture) string{

	var message bytes.Buffer 

	var fixture bytes.Buffer

	if(len(fix) == 0){

		return "No Live Matches!"

	}

	var live int = 0;

	for i := 0; i < len(fix); i++ {
		
		homeTeam := fix[i].HomeTeam
		awayTeam := fix[i].AwayTeam

		if(fix[i].CurrentState == 0){


				
		}else{

			if(fix[i].CurrentState == 1){

				fixture.WriteString(homeTeam.Name)
				fixture.WriteString(" ")
				fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
				fixture.WriteString(" - ")
				fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
				fixture.WriteString(" ")
				fixture.WriteString(awayTeam.Name)
				fixture.WriteString("	1st. Half")
				fixture.WriteString("\n")

				live++

			}else{

				if(fix[i].CurrentState == 2){
				
					fixture.WriteString(homeTeam.Name)
					fixture.WriteString(" ")
					fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
					fixture.WriteString(" - ")
					fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
					fixture.WriteString(" ")
					fixture.WriteString(awayTeam.Name)
					fixture.WriteString("	HT")
					fixture.WriteString("\n")

					live++

				}else{

					if(fix[i].CurrentState == 3){

						fixture.WriteString(homeTeam.Name)
						fixture.WriteString(" ")
						fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
						fixture.WriteString(" - ")
						fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
						fixture.WriteString(" ")
						fixture.WriteString(awayTeam.Name)
						fixture.WriteString("	2nd. Half")
						fixture.WriteString("\n")

						live++

					}else{



					}

				}

			}

		}

		message.WriteString(fixture.String()) 

	}

	if(live == 0){

		return "No Live Matches!"		

		}

	return message.String()
	
}

//Message.go

type message struct {
    
    Message string `json:"message"`
 
}

type welcomeMessage struct {
    
    Message string `json:"message"`
    Uuid string `json:"uuid"`

}

//Table.go

type table struct {

    LeagueTable []teamStanding
    Subround string

}

type teamStanding struct {

    Name string
	Points int
	Wins int 
	Losses int 
	Draws int
	GamesPlayed int
	GoalsFor int
	GoalsAgainst int
	GoalsDiff int

}

func stringfyTeamStanding(tableRow teamStanding, position int) string{

	var teamStanding bytes.Buffer

	teamStanding.WriteString(strconv.Itoa(position))
	teamStanding.WriteString(". ")
	teamStanding.WriteString(tableRow.Name)
	teamStanding.WriteString(" Points: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Points))
	teamStanding.WriteString(", W: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Wins)) 
	teamStanding.WriteString(", L: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Losses))
	teamStanding.WriteString(", D: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Draws))
	teamStanding.WriteString(", GP: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GamesPlayed))
	teamStanding.WriteString(", GF: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GoalsFor))
	teamStanding.WriteString(", GA: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GoalsAgainst))
	teamStanding.WriteString(", GD: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GoalsDiff))

	return teamStanding.String()
	
}

func stringfyTable(table table) string{

	var leagueTable bytes.Buffer

	for i := 0; i < len(table.LeagueTable); i++ {
		
		row := table.LeagueTable[i]
		leagueTable.WriteString(stringfyTeamStanding(row, i+1))
		leagueTable.WriteString("\n")

	}

	return leagueTable.String()
	
}

//TeamProfile.go

type teamProfile struct {
    
    Name string
    ShortCode  string
    Players []player
    DefaultHomeVenue stadium

}


type player struct {

    Name string
    Number  int
    Position string

}

type stadium struct {

    Name string
    Capacity  int

}

func stringfyPlayer(player player) string{
	
	var plr bytes.Buffer

	plr.WriteString("	")
	plr.WriteString(strconv.Itoa(player.Number))
	plr.WriteString(". ")
	plr.WriteString(player.Name)

	plr.WriteString(" (")
	plr.WriteString(player.Position)
	plr.WriteString(")")

	return plr.String()

}

func stringfyPlayers(players []player) string{
	
	var plrs bytes.Buffer

	for i := 0; i < len(players); i++ {
		
		plrs.WriteString(stringfyPlayer(players[i]))
		
		if(i != len(players) - 1){

			plrs.WriteString("\n")
		
		}

	}

	return plrs.String()

}


func stringfyStadium(stadium stadium) string{
	
	var std bytes.Buffer

	std.WriteString("	Name: ")
	std.WriteString(stadium.Name)
	std.WriteString("\n")
	std.WriteString("	Capacity: ")
	std.WriteString(strconv.Itoa(stadium.Capacity))

	return std.String()

}

func stringfyTeamProfile(team teamProfile) string{
	
	var profile bytes.Buffer

	profile.WriteString("Name: ")
	profile.WriteString(team.Name)
	profile.WriteString(" (")
	profile.WriteString(team.ShortCode)
	profile.WriteString(")")
	profile.WriteString("\n")
	profile.WriteString("Home Turf:\n")
	profile.WriteString(stringfyStadium(team.DefaultHomeVenue))
	profile.WriteString("\n")
	profile.WriteString("Players:\n")
	profile.WriteString(stringfyPlayers(team.Players))
	

	return profile.String()

}

