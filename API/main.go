package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type question struct {
	ID     string   `json:"id"`
	Text   string   `json:"text"`
	Option []string `json:"option"`
}

type answer struct {
	ID            string `json:"id"`
	Correctanswer string `json:"correctanswer"`
	Scorecount    int    `json:"scorecount"`
}

type userinput struct {
	ID           string `json:"id"`
	Answerbyuser string `json:"answerbyuser"`
}

var useranswered = []userinput{}

//var finalresult = []result{{Score: userResults(useranswered)}}

var questions = []question{
	{ID: "1", Text: "What the capital city of Ghana",
		Option: []string{"Berlin", "Accra", "Kumasi"}},
	{ID: "2", Text: "In which country is Arsenal",
		Option: []string{"Germany", "France", "England"}},
	{ID: "3", Text: "In what year did the second world war end",
		Option: []string{"1945", "1917", "1939"}},
	{ID: "4", Text: "What is the tallest mmountain in the world",
		Option: []string{"Everest", "Kilimanjaro", "Afadjato"}},
	{ID: "5", Text: "What country won the World cup in 1958",
		Option: []string{"Italy", "Brazil", "Uraguay"}},
	{ID: "6", Text: "In which continent is China",
		Option: []string{"Africa", "Europe", "Asia"}},
	{ID: "7", Text: "What is the national dish of Italy",
		Option: []string{"Pizza", "Pasta", "Shawarma"}},
	{ID: "8", Text: "What ocean is between North America and Africa",
		Option: []string{"Indian", "Pacific", "Atlantic"}},
	{ID: "9", Text: "What is the chemical symbol for gold",
		Option: []string{"Ag", "Af", "Au"}},
	{ID: "10", Text: "How many colours make up the rainbow",
		Option: []string{"Seven", "Ten", "Five"}},
}

var answers = []answer{
	{ID: "1", Correctanswer: "Accra", Scorecount: 0},
	{ID: "2", Correctanswer: "England", Scorecount: 0},
	{ID: "3", Correctanswer: "1945", Scorecount: 0},
	{ID: "4", Correctanswer: "Everest", Scorecount: 0},
	{ID: "5", Correctanswer: "Brazil", Scorecount: 0},
	{ID: "6", Correctanswer: "Asia", Scorecount: 0},
	{ID: "7", Correctanswer: "Pizza", Scorecount: 0},
	{ID: "8", Correctanswer: "Atlantic", Scorecount: 0},
	{ID: "9", Correctanswer: "Au", Scorecount: 0},
	{ID: "10", Correctanswer: "Seven", Scorecount: 0},
}

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func getQuestionById(id string) (*question, error) {
	for i, b := range questions {
		if b.ID == id {
			return &questions[i], nil
		}
	}
	return nil, errors.New("question not found")
}

func getAnswers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, answers)
}

func getUserAnswerById(id string) (*userinput, error) {
	for i, b := range useranswered {
		if b.ID == id {
			return &useranswered[i], nil
		}
	}

	return nil, errors.New("answer not found")
}

func getAnswerById(id string) (*answer, error) {
	for i, b := range answers {
		if b.ID == id {
			return &answers[i], nil
		}
	}

	return nil, errors.New("answer not found")
}

func userAnswerById(c *gin.Context) {
	id := c.Param("id")
	userinput, err1 := getUserAnswerById(id)
	answer, err2 := getAnswerById(id)

	if err1 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	if err2 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	if userinput.Answerbyuser == answer.Correctanswer {
		answer.Scorecount += 1

	}
	c.IndentedJSON(http.StatusOK, answer)
}

func questionById(c *gin.Context) {
	id := c.Param("id")
	question, err := getQuestionById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, question)
}

func createUserAnswer(c *gin.Context) {
	var newAnswer userinput

	if err := c.BindJSON(&newAnswer); err != nil {
		return
	}

	useranswered = append(useranswered, newAnswer)
	c.IndentedJSON(http.StatusCreated, newAnswer)
}

func getFinalResults(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, answers)

}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	router.GET("/questions", getQuestions)
	router.GET("/answeredscore", getAnswers)
	router.GET("/testscore/:id", userAnswerById)
	router.GET("/questions/:id", questionById)
	router.POST("/useranswered", createUserAnswer)
	router.GET("/finalresults", getFinalResults)
	router.Run("localhost:8080")
}
