package structures

import "time"

//go:generate easyjson -all
type UserProfile struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Age           int               `json:"age"`
	Email         *string           `json:"email,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	LastLogin     *time.Time        `json:"last_login,omitempty"`
	Preferences   map[string]string `json:"preferences"`
	Friends       []FriendProfile   `json:"friends"`
	Address       Address           `json:"address"`
	PaymentMethod *PaymentMethod    `json:"payment_method,omitempty"`
	Metadata      map[string]any    `json:"metadata"`
	Tags          []string          `json:"tags"`
}

type FriendProfile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Connected time.Time `json:"connected"`
}

type Address struct {
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	ZipCode   string `json:"zip_code"`
	IsPrimary bool   `json:"is_primary"`
}

type PaymentMethod struct {
	Type       string `json:"type"`
	CardNumber string `json:"card_number"`
	Expiry     string `json:"expiry"`
}
