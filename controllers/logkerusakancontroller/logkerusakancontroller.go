package logkerusakancontroller

import (
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var logkerusakan []models.LogKerusakan
	if err := models.DB.Debug().Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").Find(&logkerusakan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menarik data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"Kerusakan": logkerusakan})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var logkerusakan models.LogKerusakan
	id := mux.Vars(r)["id"]

	if err := models.DB.Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").First(&logkerusakan, "id = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "history kerusakan tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"Kerusakan": logkerusakan})
}