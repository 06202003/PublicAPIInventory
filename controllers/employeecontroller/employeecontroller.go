package employeecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee
	models.DB.Find(&employees)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"employees": employees})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	id := mux.Vars(r)["nomor_induk"]

	if err := models.DB.First(&employee, "nomor_induk = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Data tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"employee": employee})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	models.DB.Create(&employee)
	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	id := mux.Vars(r)["nomor_induk"]

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if models.DB.Model(&models.Employee{}).Where("nomor_induk = ?", id).Updates(&employee).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat mengupdate karyawan"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	id := mux.Vars(r)["nomor_induk"]

	if err := models.DB.First(&employee, "nomor_induk = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Karyawan tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&employee).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus karyawan"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}
