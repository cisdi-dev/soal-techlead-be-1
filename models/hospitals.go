package models

import "time"

type Hospitals struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type HospitalsPatients struct {
	Id         int       `json:"id,omitempty"`
	HospitalId int       `json:"hospital_id,omitempty"`
	PatientId  int       `json:"patient_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
