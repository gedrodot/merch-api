package model

import "time"

type Purchase struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	ItemName  string    `db:"item_name"`
	CreatedAt time.Time `db:"created_at"`
}

type Item struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

var MerchPrices = map[string]int{
	"t-shirt":    80,
	"cup":        20,
	"book":       50,
	"pen":        10,
	"powerbank":  200,
	"hoody":      300,
	"umbrella":   200,
	"socks":      10,
	"wallet":     50,
	"pink-hoody": 500,
}

func IsValidItem(itemName string) bool {
	_, exists := MerchPrices[itemName]
	return exists
}
