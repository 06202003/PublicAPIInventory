package locationcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var locations []models.Location
	models.DB.Find(&locations)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"locations": locations})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var location models.Location
	id := mux.Vars(r)["id_lokasi"]

	if err := models.DB.First(&location, "id_lokasi = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Lokasi tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"location": location})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var location models.Location

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	models.DB.Create(&location)
	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Lokasi Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var location models.Location
	id := mux.Vars(r)["id_lokasi"]

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if models.DB.Model(&models.Location{}).Where("id_lokasi = ?", id).Updates(&location).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Gagal memperbarui lokasi"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data berhasil diperbarui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id_lokasi"]

	// Check if the location exists
	var existingLocation models.Location
	if err := models.DB.First(&existingLocation, "id_lokasi = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Lokasi tidak ditemukan"})
		return
	}

	// Delete the location with the specified ID
	if err := models.DB.Where("id_lokasi = ?", id).Delete(&existingLocation).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus lokasi"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}
