package quiz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Questions struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

func LoadQuestions() []Question {
	jsonFile, err := os.Open("quiz/data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var qs Questions
	json.Unmarshal([]byte(byteValue), &qs)
	log.Printf("Loaded %d questions", len(qs.Questions))
	return qs.Questions
}
