package quiz

import (
	"log"
	"strings"
)

type answeredIds []int
type UserID int64

var quizes = make(map[UserID]answeredIds, 0)
var questions = make([]Question, 0)

func InitQuiz(userID UserID) string {
	if !isQuizStarted(userID) && !isQuizFinished(userID) {
		quizes[userID] = make(answeredIds, 0)
		questions = append(questions, LoadQuestions()...)
		log.Printf("Quiz started for user %d", userID)
		return strings.Join([]string{StartQuizMsg, getQuestion(userID)}, "\n")
	} else {
		return getQuestion(userID)
	}
}

func ProcessQuiz(userID UserID, text string) (string, bool) {
	switch {
	case isQuizStarted(userID):
		msg, isCorrect := сheckAnswer(userID, text)
		if isCorrect {
			curQ := len(quizes[userID])
			log.Printf("Correct answer for %d from user %d was wrote", curQ, userID)
			for _, q := range questions {
				if q.ID == curQ {
					quizes[userID] = append(quizes[userID], q.ID)
				}
			}

			if isQuizFinished(userID) {
				return FinishMsg, true
			} else {
				return strings.Join([]string{msg, getQuestion(userID)}, "\n"), false
			}
		}
		return msg, false
	case isQuizFinished(userID):
		return FinishMsg, true
	}
	return "", false
}

func GetHint(userID UserID) string {
	log.Printf("Creating hint for user %d", userID)
	currentQ := len(quizes[userID])
	for _, q := range questions {
		if q.ID == currentQ {
			ans := q.Answers[0]
			log.Printf("===> %s ", ans)
			subs := strings.SplitN(ans, "", -1)
			fch, lch := strings.ToUpper(subs[0]), strings.ToUpper(subs[len(subs)-1])
			hint := strings.Join([]string{fch, strings.Repeat(" _ ", len(subs)-2), lch}, " ")
			log.Printf("Hint %s for user %d", hint, userID)
			return hint
		}
	}
	return ""
}

func getQuestion(userID UserID) string {
	currentQ := len(quizes[userID])
	for _, q := range questions {
		if q.ID == currentQ {
			log.Printf("Question %d send to user %d", q.ID, userID)
			return q.Question
		}
	}

	log.Fatalln("Error occurred while sending questions")
	return ""
}

func сheckAnswer(userID UserID, text string) (string, bool) {
	log.Printf("Check answer from user %d", userID)
	currentQ := len(quizes[userID])
	for _, q := range questions {
		if q.ID == currentQ {
			for _, a := range q.Answers {
				if a == text {
					log.Printf("Correct answer %s from user %d", text, userID)
					return RightAnsMsg, true
				}
			}
			log.Printf("Wrong answer %s from user %d", text, userID)
			return WrongAnsMsg, false
		}
	}

	return "Хм...", false
}

func isQuizStarted(userID UserID) bool {
	if _, has := quizes[userID]; has {
		return true
	}
	return false
}

func isQuizFinished(userID UserID) bool {
	if len(quizes[userID]) == len(questions) && len(quizes[userID]) > 0 {
		log.Printf("User %d finished quiz", userID)
		return true
	} else {
		log.Printf("User %d didn't finish quiz", userID)
		return false
	}
}
