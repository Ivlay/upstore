package upstore

type User struct {
	ChatId    int64  `json:"chat_id"`
	UserId    int    `json:"user_id"`
	FirstName string `json:"first_name" db:"firstname"`
	UserName  string `json:"user_name" db:"username"`
	CreatedAt string `json:"created_at"`
}

type UserList struct {
	Id        int
	UserId    int
	ProductId string
}

type UserWithProducts struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	ChatId   int64  `json:"chat_id"`
	Products []Product
}
