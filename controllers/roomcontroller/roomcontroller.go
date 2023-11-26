package roomcontroller

import (
	"encoding/json"
	"net/http"
    "gorm.io/gorm"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var rooms []models.Room
	models.DB.Debug().Preload("Location").Find(&rooms)
	models.DB.Preload("Location", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Find(&rooms)

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"rooms": rooms})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	id := mux.Vars(r)["id_ruangan"]

	if err := models.DB.Preload("Location").First(&room, "id_ruangan = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Ruangan tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"room": room})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var room models.Room

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	models.DB.Create(&room)
	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	id := mux.Vars(r)["id_ruangan"]

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if models.DB.Model(&models.Room{}).Where("id_ruangan = ?", id).Updates(&room).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Gagal memperbarui ruangan"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data berhasil diperbarui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id_ruangan"]

    // Check if the room exists
    var existingRoom models.Room
    if err := models.DB.First(&existingRoom, "id_ruangan = ?", id).Error; err != nil {
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Ruangan tidak ditemukan"})
        return
    }

    // Delete the room with the specified ID
    if err := models.DB.Where("id_ruangan = ?", id).Delete(&existingRoom).Error; err != nil {
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus ruangan"})
        return
    }

    helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}
