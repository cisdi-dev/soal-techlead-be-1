package patients

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cisdi-dev/soal-techlead-be-1/models"
)

type Delivery struct {
	mux     *http.ServeMux
	route   string
	usecase *Usecase
}

func NewDelivery(mux *http.ServeMux, route string, usecase *Usecase) error {
	obj := &Delivery{
		mux:     mux,
		route:   route,
		usecase: usecase,
	}

	obj.mux.HandleFunc(route+"/create", obj.handleCreatePatient)
	obj.mux.HandleFunc(route+"/get/all", obj.handleGetAllPatient)
	obj.mux.HandleFunc(route+"/get/:patientId", obj.handleGetByIdPatient)
	obj.mux.HandleFunc(route+"/register/patient/:patientId/hospital/:hospitalId", obj.handleRegisterPatientToHospital)

	return nil
}

func (d *Delivery) handleCreatePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var requestData models.Patients
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return
	}

	err := d.usecase.CreatePatient(requestData)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *Delivery) handleGetAllPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	output, err := d.usecase.GetAllPatient()
	if err != nil {
		return
	}

	outputJSON, err := json.Marshal(&output)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func (d *Delivery) handleGetByIdPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	paramId := r.URL.Query().Get("patientId")
	if paramId == "" {
		return
	}

	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		return
	}

	output, err := d.usecase.GetByIdPatient(paramIdInt)
	if err != nil {
		return
	}

	outputJSON, err := json.Marshal(&output)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}

func (d *Delivery) handleRegisterPatientToHospital(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	paramId := r.URL.Query().Get("patientId")
	if paramId == "" {
		return
	}

	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		return
	}

	param2Id := r.URL.Query().Get("hospitalId")
	if param2Id == "" {
		return
	}

	param2IdInt, err := strconv.Atoi(param2Id)
	if err != nil {
		return
	}

	err = d.usecase.RegisterPatientToHospital(paramIdInt, param2IdInt)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
}
