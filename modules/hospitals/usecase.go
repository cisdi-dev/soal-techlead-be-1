package hospitals

import (
	"github.com/cisdi-dev/soal-techlead-be-1/models"
)

type Usecase struct {
	repo *Repository
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) CreateHospital(requestData models.Hospitals) error {
	err := u.repo.CreateHospital(requestData)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetAllHospital() ([]models.Hospitals, error) {
	results, err := u.repo.GetAllHospital()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (u *Usecase) GetByIdHospital(id int) (models.Hospitals, error) {
	result, err := u.repo.GetByIdHospital(id)
	if err != nil {
		return models.Hospitals{}, err
	}

	return result, nil
}

func (u *Usecase) GetHospitalPatientsByPatientId(patientId int) (models.HospitalsPatients, error) {
	result, err := u.repo.GetHospitalPatientsByPatientId(patientId)
	if err != nil {
		return models.HospitalsPatients{}, err
	}

	return result, nil
}
