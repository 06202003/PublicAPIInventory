package reporthistoryperbaikancontroller

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var reportHistories []models.ReportHistoryPerbaikan
	if err := models.DB.Debug().Preload("ReportHistoryKerusakan.Usage").Preload("ReportHistoryKerusakan.Usage.Inventory").Preload("ReportHistoryKerusakan.Usage.Inventory.Category").Preload("ReportHistoryKerusakan.Usage.Room").Preload("ReportHistoryKerusakan.Usage.Room.Location").Preload("ReportHistoryKerusakan.Usage.Employee").Find(&reportHistories).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menarik data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"Perbaikan": reportHistories})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var reportHistory models.ReportHistoryPerbaikan
	id := mux.Vars(r)["id"]

	if err := models.DB.Preload("ReportHistoryKerusakan.Usage").Preload("ReportHistoryKerusakan.Usage.Inventory").Preload("ReportHistoryKerusakan.Usage.Inventory.Category").Preload("ReportHistoryKerusakan.Usage.Room").Preload("ReportHistoryKerusakan.Usage.Room.Location").Preload("ReportHistoryKerusakan.Usage.Employee").First(&reportHistory, id).Error; err != nil {
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

	if err := updateUsageStatusThroughDamageRecord(reportHistoryPerbaikan.IdHistoryKerusakan, "rusak"); err != nil {
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to update usage status"})
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
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Data not found"})
		return
	}

	// Store the ID of the associated damage record
	idHistoryKerusakan := reportHistoryPerbaikan.IdHistoryKerusakan

	// Delete the repair record, which will automatically delete the corresponding damage record
	if err := models.DB.Delete(&reportHistoryPerbaikan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to delete history perbaikan"})
		return
	}

	// Now, update the status of the associated usage item to "baik" using the corresponding damage record
	if err := updateUsageStatusThroughDamageRecord(idHistoryKerusakan, "baik"); err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to update usage status"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data deleted successfully"})
}


// Function to update the usage status through the corresponding damage record
func updateUsageStatusThroughDamageRecord(idHistoryKerusakan string, status string) error {
    var reportHistoryKerusakan models.ReportHistoryKerusakan

    if err := models.DB.First(&reportHistoryKerusakan, "id = ?", idHistoryKerusakan).Error; err != nil {
        return err
    }

    // Assuming there's a direct reference to the Usage record through the ReportHistoryKerusakan record
    usageID := reportHistoryKerusakan.Usage.IdPemakaian

    // Update the status of the associated usage item to "baik"
	updateUsageStatus(usageID, status)
    return nil
}


func updateUsageStatus(usageID string, status string) {
	var usage models.Usage
	if err := models.DB.First(&usage, usageID).Error; err != nil {
		// Handle the error (Usage not found)
		fmt.Println("Usage not found")
		return
	}

	usage.Status = status
	if err := models.DB.Save(&usage).Error; err != nil {
		// Handle the error (failed to update Usage status)
		fmt.Println("Failed to update Usage status")
		return
	}
}