package patients

import (
	"database/sql"
	"strconv"

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

func (r *Repository) CreatePatient(insertData models.Patients) error {
	stmt, _ := r.db.Prepare("INSERT INTO tb_patients VALUES (?, ?, ?, ?, ?, ?, ?, ?)")

	stmt.Exec(
		insertData.IdentityNumber,
		insertData.FamilyCardNumber,
		insertData.FirstName,
		insertData.LastName,
		insertData.BirthDate,
		insertData.Address,
		insertData.BpjsNumber,
		insertData.CreatedAt,
	)

	return nil
}

func (r *Repository) GetAllPatient() ([]models.Patients, error) {
	rows, _ := r.db.Query("SELECT * FROM tb_patients")

	var results []models.Patients
	for rows.Next() {
		var row models.Patients

		err := rows.Scan(
			&row.Id,
			&row.IdentityNumber,
			&row.FamilyCardNumber,
			&row.FirstName,
			&row.LastName,
			&row.BirthDate,
			&row.Address,
			&row.BpjsNumber,
			&row.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, row)
	}

	return results, nil
}

func (r *Repository) GetByIdPatient(id int) (models.Patients, error) {
	idString := strconv.Itoa(id)

	var result models.Patients

	r.db.QueryRow("SELECT * FROM tb_patients WHERE id ="+idString).Scan(
		&result.Id,
		&result.IdentityNumber,
		&result.FamilyCardNumber,
		&result.FirstName,
		&result.LastName,
		&result.BirthDate,
		&result.Address,
		&result.BpjsNumber,
		&result.CreatedAt,
	)

	return result, nil
}

func (r *Repository) GetByIdentityNumberPatient(identityNumber string) (models.Patients, error) {
	var result models.Patients

	r.db.QueryRow("SELECT * FROM tb_patients WHERE identity_number ="+identityNumber).Scan(
		&result.Id,
		&result.IdentityNumber,
		&result.FamilyCardNumber,
		&result.FirstName,
		&result.LastName,
		&result.BirthDate,
		&result.Address,
		&result.BpjsNumber,
		&result.CreatedAt,
	)

	return result, nil
}
