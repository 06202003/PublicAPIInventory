package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    models.DB.Find(&users)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"User": users})
}

func Show(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var users models.User
    if err := models.DB.First(&users, "id = ?", id).Error; err != nil {
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Data tidak ditemukan"})
        return
    }

    helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"user": users})
}


func Create(w http.ResponseWriter, r *http.Request) {
    var users models.User

    if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
        helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
        return
    }

    if err := models.DB.Create(&users).Error; err != nil {
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat kategori"})
        return
    }

    helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
    var users models.User
    id := mux.Vars(r)["id"]

    if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
        helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
        return
    }

    if models.DB.Model(&models.Category{}).Where("id = ?", id).Updates(&users).RowsAffected == 0 {
        helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat mengupdate kategori"})
        return
    }

    helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
    var users models.User
    id := mux.Vars(r)["id"]

    if err := models.DB.First(&users, "id = ?", id).Error; err != nil {
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Kategori tidak ditemukan"})
        return
    }

    if err := models.DB.Delete(&users).Error; err != nil {
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus kategori"})
        return
    }

    helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data berhasil dihapus"})
}
