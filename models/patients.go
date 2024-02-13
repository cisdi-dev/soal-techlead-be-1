package models

import "time"

type Patients struct {
	Id               int       `json:"id,omitempty"`
	IdentityNumber   string    `json:"identity_number,omitempty"`
	FamilyCardNumber string    `json:"family_card_number,omitempty"`
	FirstName        string    `json:"first_name,omitempty"`
	LastName         string    `json:"last_name,omitempty"`
	BirthDate        time.Time `json:"birth_date,omitempty"`
	Address          string    `json:"address,omitempty"`
	BpjsNumber       string    `json:"bpjs_number,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}
