package usagecontroller

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var usages []models.Usage
	models.DB.Debug().Preload("Employee").Preload("Inventory").Preload("Inventory.Category").Preload("Room").Preload("Room.Location").Find(&usages)
	models.DB.Preload("Employee").Preload("Inventory").Preload("Inventory.Category").Preload("Room").Preload("Room.Location").Find(&usages)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"usages": usages})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage
	id := mux.Vars(r)["id_pemakaian"]
	

	if err := models.DB.Preload("Employee").Preload("Inventory").Preload("Inventory.Category").Preload("Room").Preload("Room.Location").First(&usage, "id_pemakaian = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Pemakaian tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"usage": usage})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage

	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	models.DB.Create(&usage)

	createHistoryPemakaian(usage, "", usage.EmployeeID, "", usage.IdRuang)

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Pemakaian Berhasil Dibuat"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage
	id := mux.Vars(r)["id_pemakaian"]

	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	var oldUsage models.Usage
	models.DB.First(&oldUsage, "id_pemakaian = ?", id)

	

	if models.DB.Model(&models.Usage{}).Where("id_pemakaian = ?", id).Updates(&usage).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Gagal memperbarui pemakaian"})
		return
	}

	createHistoryPemakaian(oldUsage, oldUsage.EmployeeID, usage.EmployeeID, oldUsage.IdRuang, usage.IdRuang)

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data berhasil diperbarui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id_pemakaian"]

	// Check if the usage exists
	var existingUsage models.Usage
	if err := models.DB.First(&existingUsage, "id_pemakaian = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Pemakaian tidak ditemukan"})
		return
	}

	// Delete the usage with the specified ID
	if err := models.DB.Where("id_pemakaian = ?", id).Delete(&existingUsage).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus pemakaian"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}

func createHistoryPemakaian(usage models.Usage, oldEmployeeID, newEmployeeID, oldIdRuang, newIdRuang string) {
	historyPemakaian := models.ReportHistoryPemakaian{
		OldEmployeeID: oldEmployeeID,
		NewEmployeeID: newEmployeeID,
		OldRoom:       oldIdRuang,
		NewRoom:       newIdRuang,
		UsageDate:     usage.UpdatedAt,
		IdPemakaian:     usage.IdPemakaian,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	// Create a new history pemakaian record
	models.DB.Create(&historyPemakaian)
}