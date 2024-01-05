package skillcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var skills []models.Skill
	models.DB.Find(&skills)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"skills": skills})
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id_skill"]

	var skill models.Skill
	if err := models.DB.First(&skill, "id_skill = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Skill tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"skill": skill})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill

	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&skill).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat skill"})
		return
	}

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill
	id := mux.Vars(r)["id_skill"]

	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if models.DB.Model(&models.Skill{}).Where("id_skill = ?", id).Updates(&skill).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat mengupdate skill"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill
	id := mux.Vars(r)["id_skill"]

	if err := models.DB.First(&skill, "id_skill = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Skill tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&skill).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus skill"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}
