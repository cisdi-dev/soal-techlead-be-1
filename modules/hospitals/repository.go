package hospitals

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/cisdi-dev/soal-techlead-be-1/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateHospital(insertData models.Hospitals) error {
	stmt, _ := r.db.Prepare("INSERT INTO tb_hospitals VALUES (?, ?, ?)")

	stmt.Exec(
		insertData.Name,
		insertData.Address,
		insertData.CreatedAt,
	)

	return nil
}

func (r *Repository) GetAllHospital() ([]models.Hospitals, error) {
	rows, _ := r.db.Query("SELECT * FROM tb_hospitals")

	var results []models.Hospitals
	for rows.Next() {
		var row models.Hospitals

		err := rows.Scan(
			&row.Id,
			&row.Name,
			&row.Address,
			&row.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, row)
	}

	return results, nil
}

func (r *Repository) GetByIdHospital(id int) (models.Hospitals, error) {
	idString := strconv.Itoa(id)

	var result models.Hospitals

	r.db.QueryRow("SELECT * FROM tb_hospitals WHERE id ="+idString).Scan(
		&result.Id,
		&result.Name,
		&result.Address,
		&result.CreatedAt,
	)

	return result, nil
}

func (r *Repository) GetByNameHospital(name string) (models.Hospitals, error) {
	var result models.Hospitals

	r.db.QueryRow("SELECT * FROM tb_hospitals WHERE name ="+name).Scan(
		&result.Id,
		&result.Name,
		&result.Address,
		&result.CreatedAt,
	)

	return result, nil
}

func (r *Repository) RegisterHospitalPatient(patientId int, hospitalId int) error {
	stmt, _ := r.db.Prepare("INSERT INTO tb_hospitals_patients VALUES (?, ?, ?)")

	stmt.Exec(
		hospitalId,
		patientId,
		time.Now(),
	)

	return nil
}

func (r *Repository) GetHospitalPatientsByPatientId(patientId int) (models.HospitalsPatients, error) {
	idString := strconv.Itoa(patientId)

	var result models.HospitalsPatients

	r.db.QueryRow("SELECT * FROM tb_hospitals_patients WHERE patient_id ="+idString).Scan(
		&result.Id,
		&result.HospitalId,
		&result.PatientId,
		&result.CreatedAt,
	)

	return models.HospitalsPatients{}, fmt.Errorf("data not found")
}
