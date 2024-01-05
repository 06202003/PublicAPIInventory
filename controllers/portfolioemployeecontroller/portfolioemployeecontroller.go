package portfolioemployeecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var portfolioEmployees []models.PortfolioEmployee
	models.DB.Preload("Portfolio").Preload("Employee").Find(&portfolioEmployees)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"portfolioEmployees": portfolioEmployees})
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id_portfolio_employee"]

	var portfolioEmployee models.PortfolioEmployee
	if err := models.DB.Preload("Portfolio").Preload("Employee").First(&portfolioEmployee, "id_portfolio_karyawan = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Portfolio Employee tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"portfolioEmployee": portfolioEmployee})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var portfolioEmployee models.PortfolioEmployee

	if err := json.NewDecoder(r.Body).Decode(&portfolioEmployee); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&portfolioEmployee).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat Portfolio Employee"})
		return
	}

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var portfolioEmployee models.PortfolioEmployee
	id := mux.Vars(r)["id_portfolio_employee"]

	if err := json.NewDecoder(r.Body).Decode(&portfolioEmployee); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if models.DB.Model(&models.PortfolioEmployee{}).Where("id_portfolio_karyawan = ?", id).Updates(&portfolioEmployee).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat mengupdate Portfolio Employee"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var portfolioEmployee models.PortfolioEmployee
	id := mux.Vars(r)["id_portfolio_employee"]

	if err := models.DB.First(&portfolioEmployee, "id_portfolio_karyawan = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Portfolio Employee tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&portfolioEmployee).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus Portfolio Employee"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}

func ViewByEmployeePortfolio(w http.ResponseWriter, r *http.Request) {
	idKaryawan := mux.Vars(r)["nomor_induk"]

	var portfolioEmployee []models.PortfolioEmployee
	models.DB.Debug().Preload("Portfolio").Preload("Employee").First(&portfolioEmployee, "nomor_induk = ?", idKaryawan)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"portfolioEmployee": portfolioEmployee})
}
