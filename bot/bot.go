package bot

import (
	"bot/http"
	"log"
	"strings"

	"bot/quiz"
)

type Bot struct {
	httpClient http.TelegramAPIClient
}

func (b Bot) ProcessUpdate(update Update) {
	msg := strings.ToLower(update.Message.Text)
	userId := update.Message.Chat.ID
	if strings.HasPrefix(msg, "/") {
		switch msg {
		case "/start":
			log.Printf("Get command /start from %d", userId)
			b.SendMessage(userId, HelloMsg)
			b.SendMessage(userId, WantQuizMsg)
		case "/quiz":
			log.Printf("Get command /quiz from %d", userId)
			ans := quiz.InitQuiz(quiz.UserID(userId))
			b.SendMessage(userId, ans)
		case "/hint":
			log.Printf("Get command /hint from %d", userId)
			b.SendMessage(userId, quiz.HintMsg)
			b.SendMessage(userId, quiz.GetHint(quiz.UserID(userId)))
		case "/restart":
			log.Printf("Get command /restart from %d", userId)
			b.SendMessage(userId, "Сорри, моя хозяйка не успела меня научить этой команде :(")

		default:
			//сорри
		}
	} else {
		ans, finished := quiz.ProcessQuiz(quiz.UserID(userId), msg)
		if !finished {
			b.SendMessage(userId, ans)
		} else {
			b.SendPhoto(userId, PresentMsg)
		}
	}
}

func NewBot(httpClient http.TelegramAPIClient) *Bot {
	return &Bot{httpClient: httpClient}
}
