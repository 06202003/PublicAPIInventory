package usagecontroller

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"strconv"
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

// func Create(w http.ResponseWriter, r *http.Request) {
// 	var usage models.Usage

// 	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
// 		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	models.DB.Create(&usage)

// 	createHistoryPemakaian(usage, "", usage.EmployeeID, "", usage.IdRuang)

// 	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Pemakaian Berhasil Dibuat"})
// }

func Create(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage

	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Check if the provided IdPemakaian already exists
	var existingUsage models.Usage
	if err := models.DB.First(&existingUsage, "id_pemakaian = ?", usage.IdPemakaian).Error; err == nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "IdPemakaian sudah ada"})
		return
	}

	// If IdPemakaian is not provided or empty, generate a new one
	if usage.IdPemakaian == "" {
		// Fetch the associated inventory
		var inventory models.Inventory
		if err := models.DB.Model(&models.Inventory{}).Where("kode_aset = ?", usage.AssetCode).First(&inventory).Error; err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Asset tidak ditemukan"})
			return
		}

		// Fetch the category ID from the inventory
		categoryID := inventory.CategoryID

		// Generate the new IDPemakaian
		lastUsage := getLastUsageByCategory(categoryID)
		usage.IdPemakaian = GenerateIDPemakaian(categoryID, lastUsage)
	}

	// Create the new entry
	models.DB.Create(&usage)

	// Pass an empty Usage as the oldUsage parameter since it's a new entry
	createHistoryPemakaian(models.Usage{}, usage)

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Pemakaian Berhasil Dibuat"})
}

// Update updates an existing usage record.
func Update(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage
	id := mux.Vars(r)["id_pemakaian"]

	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	var oldUsage models.Usage
	if err := models.DB.First(&oldUsage, "id_pemakaian = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Pemakaian tidak ditemukan"})
		return
	}

	// Check if the record has been modified
	if !isUsageModified(oldUsage, usage) {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Data pemakaian tidak berubah"})
		return
	}

	// Update the record
	models.DB.Model(&models.Usage{}).Where("id_pemakaian = ?", id).Updates(&usage)

	// Retrieve the updated record
	var updatedUsage models.Usage
	models.DB.First(&updatedUsage, "id_pemakaian = ?", id)

	// Create history
	createHistoryPemakaian(oldUsage, updatedUsage)

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data berhasil diperbarui"})
}

// isUsageModified checks if the usage record has been modified.
func isUsageModified(oldUsage, newUsage models.Usage) bool {
	// Compare relevant fields to determine if the record has been modified
	return oldUsage.EmployeeID != newUsage.EmployeeID || oldUsage.IdRuang != newUsage.IdRuang
}

// getLastUsageByCategory retrieves the last usage record for a given categoryID.
func getLastUsageByCategory(categoryID string) models.Usage {
	var lastUsage models.Usage
	models.DB.Order("id_pemakaian DESC").First(&lastUsage, "id_ruangan = ?", categoryID)
	return lastUsage
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

func createHistoryPemakaian(oldUsage, newUsage models.Usage) {
	// Check if the history entry already exists
	var existingHistory models.ReportHistoryPemakaian
	if err := models.DB.First(&existingHistory, "id_pemakaian = ?", newUsage.IdPemakaian).Error; err == nil {
		// If it exists, update the existing entry
		existingHistory.OldEmployeeID = oldUsage.EmployeeID
		existingHistory.NewEmployeeID = newUsage.EmployeeID
		existingHistory.OldRoom = oldUsage.IdRuang
		existingHistory.NewRoom = newUsage.IdRuang
		existingHistory.UsageDate = newUsage.UpdatedAt
		existingHistory.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

		models.DB.Save(&existingHistory)
	} else {
		// If it doesn't exist, create a new entry
		historyPemakaian := models.ReportHistoryPemakaian{
			OldEmployeeID: oldUsage.EmployeeID,
			NewEmployeeID: newUsage.EmployeeID,
			OldRoom:       oldUsage.IdRuang,
			NewRoom:       newUsage.IdRuang,
			UsageDate:     newUsage.UpdatedAt,
			IdPemakaian:   newUsage.IdPemakaian,
			CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		}

		// Create a new history pemakaian record
		models.DB.Create(&historyPemakaian)
	}
}


func GenerateIDPemakaian(categoryID string, lastUsage models.Usage) string {
	if lastUsage.IdPemakaian == "" {
		return fmt.Sprintf("%s-001", categoryID)
	}

	// Extract the last three characters after the '-' and convert to int
	lastIndex, err := strconv.Atoi(lastUsage.IdPemakaian[len(categoryID)+1:])
	if err != nil {
		return "" // or handle the error accordingly
	}

	newIndex := lastIndex + 1

	// Format the new index with leading zeros and concatenate with the category ID
	return fmt.Sprintf("%s-%03d", categoryID, newIndex)
}

func ViewByRoom(w http.ResponseWriter, r *http.Request) {
	idRuang := mux.Vars(r)["id_ruangan"]

	var usages []models.Usage
	models.DB.Debug().Preload("Employee").Preload("Inventory").Preload("Inventory.Category").Preload("Room").Preload("Room.Location").Find(&usages, "id_ruangan = ?", idRuang)

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"usages": usages})
}

// ViewByEmployee returns a list of usages filtered by employee ID.
func ViewByEmployee(w http.ResponseWriter, r *http.Request) {
	idKaryawan := mux.Vars(r)["nomor_induk"]

	var usages []models.Usage
	models.DB.Debug().Preload("Employee").Preload("Inventory").Preload("Inventory.Category").Preload("Room").Preload("Room.Location").Find(&usages, "nomor_induk = ?", idKaryawan)

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"usages": usages})
}