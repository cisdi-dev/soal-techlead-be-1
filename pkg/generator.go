package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cisdi-dev/soal-techlead-be-1/models"
)

func GeneratePatients() {
	patients := []models.Patients{
		{
			Id:               1,
			IdentityNumber:   "3298612011850022",
			FamilyCardNumber: "3298640418130021",
			FirstName:        "Abdul",
			LastName:         "Halim",
			Address:          "Jl. Address No.1",
			BpjsNumber:       "0004540111021",
			CreatedAt:        time.Now(),
		},
		{
			Id:               2,
			IdentityNumber:   "3131441702431025",
			FamilyCardNumber: "3131472295041062",
			FirstName:        "Noer",
			LastName:         "Hafizon",
			Address:          "Jl. Address No.2",
			BpjsNumber:       "0003610222032",
			CreatedAt:        time.Now(),
		},
	}

	patientsJSON, _ := json.Marshal(patients)
	err := os.WriteFile("tb_patients.json", patientsJSON, 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func GenerateHospitals() {
	hospitals := []models.Hospitals{
		{
			Id:        1,
			Name:      "RS Mekar Satu",
			Address:   "Jl. Mekar Satu No.1",
			CreatedAt: time.Now(),
		},
		{
			Id:        2,
			Name:      "RS Mekar Dua",
			Address:   "Jl. Mekar Dua No.1",
			CreatedAt: time.Now(),
		},
	}

	hospitalsJSON, _ := json.Marshal(hospitals)
	err := os.WriteFile("tb_hospitals.json", hospitalsJSON, 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func GenerateHospitalsPatients() {
	hospitals := []models.HospitalsPatients{
		{
			Id:         1,
			HospitalId: 1,
			PatientId:  2,
			CreatedAt:  time.Now(),
		},
		{
			Id:         2,
			HospitalId: 2,
			PatientId:  1,
			CreatedAt:  time.Now(),
		},
	}

	hospitalsJSON, _ := json.Marshal(hospitals)
	err := os.WriteFile("tb_hospitals_patients.json", hospitalsJSON, 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
