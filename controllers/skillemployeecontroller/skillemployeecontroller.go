package skillemployeecontroller

import (
	"encoding/json"
	"net/http"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var skillEmployees []models.SkillEmployee
	models.DB.Preload("Skill").Preload("Employee").Find(&skillEmployees)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"skillEmployees": skillEmployees})
}



func Show(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id_skill_employee"]

    var skillEmployees models.SkillEmployee
    if err := models.DB.Preload("Skill").Preload("Employee").First(&skillEmployees, "id_skill_karyawan = ?", id).Error; err != nil {
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Skill Employee tidak ditemukan"})
        return
    }

    // Responding with the fetched category in the JSON format using the helper
    helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"skillEmployee": skillEmployees})
}


func Create(w http.ResponseWriter, r *http.Request) {
	var skillEmployee models.SkillEmployee

	if err := json.NewDecoder(r.Body).Decode(&skillEmployee); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&skillEmployee).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat Skill Employee"})
		return
	}

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var skillEmployee models.SkillEmployee
	id := mux.Vars(r)["id_skill_employee"]

	if err := json.NewDecoder(r.Body).Decode(&skillEmployee); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if models.DB.Model(&models.SkillEmployee{}).Where("id_skill_karyawan = ?", id).Updates(&skillEmployee).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat mengupdate Skill Employee"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var skillEmployee models.SkillEmployee
	id := mux.Vars(r)["id_skill_employee"]

	if err := models.DB.First(&skillEmployee, "id_skill_karyawan = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Skill Employee tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&skillEmployee).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus Skill Employee"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}

func ViewByEmployeeSkill(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["nomor_induk"]

	var skillEmployee []models.SkillEmployee
	models.DB.Debug().Preload("Skill").Preload("Employee").First(&skillEmployee, "nomor_induk = ?", id)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"skillEmployee": skillEmployee})
}
