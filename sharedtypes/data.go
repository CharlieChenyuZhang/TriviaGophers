package sharedtypes

type Data struct {
	Title    string
	Body     string
	Question string
	Answer   string
	Players  Players
	Category string
	Teams []string
}

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type TriviaData struct {
	Question string
	Answer   string
}
type Players struct {
	Players []Player `json:"players"`
}