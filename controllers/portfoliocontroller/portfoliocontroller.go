package portfoliocontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var portfolios []models.Portfolio
	models.DB.Find(&portfolios)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"portfolios": portfolios})
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id_portfolio"]

	var portfolio models.Portfolio
	if err := models.DB.First(&portfolio, "id_portfolio = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Portfolio tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"portfolio": portfolio})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var portfolio models.Portfolio

	if err := json.NewDecoder(r.Body).Decode(&portfolio); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&portfolio).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat portfolio"})
		return
	}

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var portfolio models.Portfolio
	id := mux.Vars(r)["id_portfolio"]

	if err := json.NewDecoder(r.Body).Decode(&portfolio); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if models.DB.Model(&models.Portfolio{}).Where("id_portfolio = ?", id).Updates(&portfolio).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat mengupdate portfolio"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var portfolio models.Portfolio
	id := mux.Vars(r)["id_portfolio"]

	if err := models.DB.First(&portfolio, "id_portfolio = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Portfolio tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&portfolio).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus portfolio"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}
