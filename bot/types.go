package bot

type getMeResponse struct {
	Ok     bool
	Result *User
}

type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type Update struct {
	ID      int64    `json:"update_id"`
	Message *Message `json:"message"`
}

type Message struct {
	ID   int64  `json:"message_id"`
	From *User  `json:"from"`
	Chat *Chat  `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ID       int64  `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

const (
	HelloMsg    = "Привет! Я вижу у тебя сегодня День Рождения? 🥰"
	WantQuizMsg = "Поздравляю от всей души!🎉 Хочу вручить тебе подарок, но перед этим давай поиграем! Вызывай скорее команду для квиза!!! \n\nПы.сы. Если тебе нужна будет помощь, воспользуйся подсказкой!"
	PresentMsg  = `С Днем Рождения! Я желаю тебе здоровья, гармонии, удачи во всех начинаниях. Желаю тебе продолжать двигаться своим неповторимым путем, освещать другим путь и замечать все красоты этой невероятной жизни. Ты моя любовь, моя опора, моя радость. Я желаю тебе всего наилучшего. Пусть сбудется все, о чем ты мечтаешь ❤️

Моим подарком будет массаж в месте, которое я уже больше полугода вспоминаю с теплотой. Надеюсь, что этот час ты проведешь в гармонии разумом и телом и сполна насладишься атмосферой покоя и расслабления. 
	
Тебя будут ждать по адресу: Thai Thai, Ленинградский просп., 29, корп. 3. 
	
	https://yandex.com/maps/org/tay_tay/53958677295`
)
