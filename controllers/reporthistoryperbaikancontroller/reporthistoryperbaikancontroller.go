package reporthistoryperbaikancontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var reportHistories []models.ReportHistoryPerbaikan
	if err := models.DB.Debug().Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").Find(&reportHistories).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menarik data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"Perbaikan": reportHistories})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var reportHistory models.ReportHistoryPerbaikan
	id := mux.Vars(r)["id"]

	if err := models.DB.Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").First(&reportHistory, id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "history perbaikan tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"Perbaikan": reportHistory})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var reportHistoryPerbaikan models.ReportHistoryPerbaikan

	if err := json.NewDecoder(r.Body).Decode(&reportHistoryPerbaikan); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&reportHistoryPerbaikan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat history perbaikan"})
		return
	}

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var reportHistoryPerbaikan models.ReportHistoryPerbaikan
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&reportHistoryPerbaikan); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Model(&reportHistoryPerbaikan).Where("id = ?", id).Updates(&reportHistoryPerbaikan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Gagal memperbarui history perbaikan"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data berhasil diperbarui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var reportHistoryPerbaikan models.ReportHistoryPerbaikan
	id := mux.Vars(r)["id"]

	if err := models.DB.First(&reportHistoryPerbaikan, "id = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Data tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&reportHistoryPerbaikan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus history perbaikan"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}

