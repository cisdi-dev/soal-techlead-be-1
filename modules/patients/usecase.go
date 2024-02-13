package patients

import (
	"github.com/cisdi-dev/soal-techlead-be-1/models"
	"github.com/cisdi-dev/soal-techlead-be-1/modules/hospitals"
)

type Usecase struct {
	repo          *Repository
	repoHospitals *hospitals.Repository
}

func NewUsecase(repo *Repository, repoHospitals *hospitals.Repository) *Usecase {
	return &Usecase{
		repo:          repo,
		repoHospitals: repoHospitals,
	}
}

func (u *Usecase) CreatePatient(requestData models.Patients) error {
	err := u.repo.CreatePatient(requestData)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetAllPatient() ([]models.Patients, error) {
	results, err := u.repo.GetAllPatient()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (u *Usecase) GetByIdPatient(id int) (models.Patients, error) {
	result, err := u.repo.GetByIdPatient(id)
	if err != nil {
		return models.Patients{}, err
	}

	return result, nil
}

func (u *Usecase) RegisterPatientToHospital(patientId int, hospitalId int) error {
	err := u.repoHospitals.RegisterHospitalPatient(patientId, hospitalId)
	if err != nil {
		return err
	}

	return nil
}
