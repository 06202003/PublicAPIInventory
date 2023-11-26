package reporthistorypemakaiancontroller

import (
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var reportHistories []models.ReportHistoryPemakaian
	if err := models.DB.Debug().Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").Find(&reportHistories).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menarik data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"ReportHistories": reportHistories})
}

func Show(w http.ResponseWriter, r *http.Request) {
    var reportHistory models.ReportHistoryPemakaian
    id := mux.Vars(r)["id"]

    if err := models.DB.Preload("Usage").Preload("Usage.Inventory").Preload("Usage.Inventory.Category").Preload("Usage.Room").Preload("Usage.Room.Location").Preload("Usage.Employee").First(&reportHistory, "id = ?", id).Error; err != nil {
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "History pemakaian tidak ditemukan"})
        return
    }

    helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"ReportHistory": reportHistory})
}
