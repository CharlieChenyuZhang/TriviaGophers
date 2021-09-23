package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"io/ioutil"
	"sort"
	"strings"
	"strconv"
	"github.com/manulife-ca/aff-trivia-gophers/api/db"
	"github.com/manulife-ca/aff-trivia-gophers/api/trivia"
	"github.com/manulife-ca/aff-trivia-gophers/sharedtypes"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

// Renders the templates
func renderTemplate(w http.ResponseWriter, tmpl string, page *sharedtypes.Data) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCheckboxValue(formValue []string) string {
	if len(formValue) > 0 {
		return formValue[0]
	} else {
		return ""
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var players sharedtypes.Players
	r.ParseForm()
	var teamName string
	if(r.FormValue("teamName") == ""){
		teamName = "endzymex" // FIXME: @bruce reset this to all
	}else{
		teamName = r.FormValue("teamName")
	}
	players = db.GetScore(teamName)
	sort.Slice(players.Players, func(i, j int) bool {
		return players.Players[i].Score > players.Players[j].Score
	})
	items, _ := ioutil.ReadDir("./teams")
	teams := []string{}
    for _, item := range items {
        if item.IsDir() {
            subitems, _ := ioutil.ReadDir(item.Name())
            for _, subitem := range subitems {
                if !subitem.IsDir() {
                    fmt.Println(item.Name() + "/" + subitem.Name())
                }
            }
        } else {
			teams = append(teams,strings.TrimSuffix(item.Name(), ".json"))
            fmt.Println(strings.TrimSuffix(item.Name(), ".json"))
        }
    }
	page := &sharedtypes.Data{Title: "Leaders Board", Body: "Welcome to the Trivia Game", Players: players,Teams:teams}

	renderTemplate(w, "leadersboard", page)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var winners []string
	var players sharedtypes.Players
	players = db.GetScore("endzymex")
	count := 0
	for count < len(players.Players) {
		winners = append(winners, GetCheckboxValue(r.Form["participant" + strconv.Itoa(count)]))
		count++
	}
	fmt.Println("winners", winners)
	db.UpdatePlayers(winners, false)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	var winners []string
	db.UpdatePlayers(winners, true)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func SwtichTeamHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	teamName := r.FormValue("teamName")
	fmt.Println("teamName", teamName)
	//db.UpdatePlayers(winners)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func TriviaHandler(w http.ResponseWriter, r *http.Request) {
	category := "general"
	r.ParseForm()
	if len(r.Form["category"]) > 0 {
		category = r.Form["category"][0]
	}
	triviaDataJson := trivia.GetAQuestion(category)

	var players sharedtypes.Players
	players = db.GetScore("endzymex")

	var triviaData sharedtypes.TriviaData
	json.Unmarshal([]byte(triviaDataJson), &triviaData)
	page := &sharedtypes.Data{Title: "Trivia page", Body: "This is our trivia page.", Question: triviaData.Question, Answer: triviaData.Answer, Players: players, Category: category}
	renderTemplate(w, "trivia", page)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	page := &sharedtypes.Data{Title: "Admin page"}
	renderTemplate(w, "admin", page)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/admin/", AdminHandler)
	http.HandleFunc("/trivia/", TriviaHandler)
	http.HandleFunc("/update", UpdateHandler)
	http.HandleFunc("/switch", SwtichTeamHandler)
	http.HandleFunc("/reset", ResetHandler)
	http.HandleFunc("/scores/", db.GetScores)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := "8080"
	fmt.Printf("running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
