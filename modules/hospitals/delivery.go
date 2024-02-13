package hospitals

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

	obj.mux.HandleFunc(route+"/create", obj.handleCreateHospital)
	obj.mux.HandleFunc(route+"/get/all", obj.handleGetAllHospital)
	obj.mux.HandleFunc(route+"/get/:hospitalId", obj.handleGetByIdHospital)
	obj.mux.HandleFunc(route+"/get/patient/:patientId", obj.handleGetHospitalPatientsByPatientId)

	return nil
}

func (d *Delivery) handleCreateHospital(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var requestData models.Hospitals
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return
	}

	err := d.usecase.CreateHospital(requestData)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *Delivery) handleGetAllHospital(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	output, err := d.usecase.GetAllHospital()
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

func (d *Delivery) handleGetByIdHospital(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	paramId := r.URL.Query().Get("hospitalId")
	if paramId == "" {
		return
	}

	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		return
	}

	output, err := d.usecase.GetByIdHospital(paramIdInt)
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

func (d *Delivery) handleGetHospitalPatientsByPatientId(w http.ResponseWriter, r *http.Request) {
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

	output, err := d.usecase.GetHospitalPatientsByPatientId(paramIdInt)
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
