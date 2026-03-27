package model

import "time"

type Transaction struct {
	ID         int       `db:"id"`
	FromUserID int       `db:"from_user_id"`
	ToUserID   int       `db:"to_user_id"`
	Amount     int       `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
}

type SendCoinRequest struct {
	ToUser string `json:"toUser" validate:"required"`
	Amount int    `json:"amount" validate:"required,gt=0"`
}

type CoinHistory struct {
	Received []ReceivedTransaction `json:"received"`
	Sent     []SentTransaction     `json:"sent"`
}

type ReceivedTransaction struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type SentTransaction struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}
