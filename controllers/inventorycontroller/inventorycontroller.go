package inventorycontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var inventory []models.Inventory
	models.DB.Debug().Preload("Category").Find(&inventory)
	models.DB.Preload("Category").Find(&inventory)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"inventory": inventory})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	id := mux.Vars(r)["kode_aset"]

	models.DB.Debug().Preload("Category").Find(&inventory, "kode_aset = ?", id)
	if err := models.DB.Preload("Category").Where("kode_aset = ?", id).First(&inventory).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Aset tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"inventory": inventory})
}

// func Create(w http.ResponseWriter, r *http.Request) {
// 	var inventory models.Inventory

// 	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
// 		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Create a new inventory record
// 	models.DB.Create(&inventory)

// 	// Create a new history pemakaian record
// 	createHistoryPemakaian(inventory, "", inventory.EmployeeID, "", inventory.RoomID)

// 	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data Berhasil Dibuat"})
// }

func Create(w http.ResponseWriter, r *http.Request) {
	var inventories []models.Inventory

	if err := json.NewDecoder(r.Body).Decode(&inventories); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if len(inventories) == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "No inventory items provided"})
		return
	}

	// Use a database transaction
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Loop through each inventory item and perform necessary calculations
	for i := range inventories {
		if inventories[i].UsefulLife == 0 {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "UsefulLife cannot be zero"})
			return
		}

		calculateDepreciation(&inventories[i])

		// Create a new inventory record
		if err := tx.Create(&inventories[i]).Error; err != nil {
			tx.Rollback()
			helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to create inventory record"})
			return
		}
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Aset Berhasil Dibuat"})
}

func calculateDepreciation(inv *models.Inventory) {
	inv.Depreciation = (inv.Price - (inv.Price / 4)) / inv.UsefulLife
	inv.Year1 = inv.Price - inv.Depreciation
	inv.Year2 = inv.Year1 - inv.Depreciation
	inv.Year3 = inv.Year2 - inv.Depreciation
	inv.Year4 = inv.Year3 - inv.Depreciation
}

func Update(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	id := mux.Vars(r)["kode_aset"]

	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Update the inventory record
	if models.DB.Model(&models.Inventory{}).Where("kode_aset = ?", id).Updates(&inventory).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat memperbarui Aset"})
		return
	}

	helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	id := mux.Vars(r)["kode_aset"]

	if err := models.DB.First(&inventory, "kode_aset = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Aset tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&inventory).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus Aset"})
		return
	}

	helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}
