package main

import (
	"github.com/orayew2002/json-bench/structures"
	"time"
)

var UserProfileData = structures.UserProfile{
	ID:        "12345",
	Name:      "Alice",
	Age:       30,
	Email:     stringPtr("alice@example.com"),
	CreatedAt: time.Now(),
	LastLogin: nil,
	Preferences: map[string]string{
		"theme":    "dark",
		"language": "en",
	},
	Friends: []structures.FriendProfile{
		{
			ID:        "54321",
			Name:      "Bob",
			Connected: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        "67890",
			Name:      "Charlie",
			Connected: time.Now().Add(-48 * time.Hour),
		},
	},
	Address: structures.Address{
		Country:   "USA",
		City:      "New York",
		Street:    "123 Main St",
		ZipCode:   "10001",
		IsPrimary: true,
	},
	PaymentMethod: &structures.PaymentMethod{
		Type:       "Credit Card",
		CardNumber: "1234-5678-9876-5432",
		Expiry:     "12/25",
	},
	Metadata: map[string]any{
		"last_purchase": "2025-04-30",
		"account_level": 5,
	},
	Tags: []string{"premium", "active", "verified"},
}

func stringPtr(s string) *string {
	return &s
}
