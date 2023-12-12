package reporthistorykerusakancontroller

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var reportkerusakanHistories []models.ReportHistoryKerusakan
	if err := models.DB.Debug().Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").Find(&reportkerusakanHistories).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menarik data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"Kerusakan": reportkerusakanHistories})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var reportkerusakanHistories models.ReportHistoryKerusakan
	id := mux.Vars(r)["id"]

	if err := models.DB.Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").First(&reportkerusakanHistories, id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "history kerusakan tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"Kerusakan": reportkerusakanHistories})
}

func Create(w http.ResponseWriter, r *http.Request) {
    var reportkerusakanHistories models.ReportHistoryKerusakan

    if err := json.NewDecoder(r.Body).Decode(&reportkerusakanHistories); err != nil {
        helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
        return
    }

	usageID := reportkerusakanHistories.IdPemakaian

    // Generate the next available ID
    nextID, err := generateNextID(usageID)
    if err != nil {
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat history perbaikan"})
        return
    }

    // Set the IdHistory field with the generated nextID
    reportkerusakanHistories.IdHistory = nextID

    if err := models.DB.Create(&reportkerusakanHistories).Error; err != nil {
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat history perbaikan"})
        return
    }

    // Update the status of the associated usage item to "rusak"
    updateUsageStatus(reportkerusakanHistories.IdPemakaian, "rusak")

    helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}


func Update(w http.ResponseWriter, r *http.Request) {
	var ReportHistoryKerusakan models.ReportHistoryKerusakan
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&ReportHistoryKerusakan); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Model(&ReportHistoryKerusakan).Where("id = ?", id).Updates(&ReportHistoryKerusakan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Gagal memperbarui history kerusakan"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data berhasil diperbarui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
    var reportkerusakanHistories models.ReportHistoryKerusakan
    id := mux.Vars(r)["id"]

    if err := models.DB.First(&reportkerusakanHistories, "id = ?", id).Error; err != nil {
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Data tidak ditemukan"})
        return
    }

    // Update the status of the associated usage item to "baik"
    updateUsageStatus(reportkerusakanHistories.Usage.IdPemakaian, "baik")

    if err := models.DB.Delete(&reportkerusakanHistories).Error; err != nil {
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus history perbaikan"})
        return
    }

    helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}

func updateUsageStatus(usageID string, status string) {
    var usage models.Usage
    if err := models.DB.Where("id_pemakaian = ?", usageID).First(&usage).Error; err != nil {
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


func generateNextID(usageID string) (string, error) {
    var count int64
    if err := models.DB.Model(&models.ReportHistoryKerusakan{}).Where("id_pemakaian = ?", usageID).Error; err != nil {
        return "", err
    }

    // Increment the count to generate the next ID
    nextID := fmt.Sprintf("%s-%03d", usageID, count+1)
    return nextID, nil
}
