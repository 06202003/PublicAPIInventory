package categorycontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

// func Index(w http.ResponseWriter, r *http.Request) {
//     var categories []models.Category
//     models.DB.Find(&categories)

//     // Your existing response logic here
//     w.WriteHeader(http.StatusOK)
//     w.Header().Set("Content-Type", "application/json")
//     helper.ResponseJSON(w, http.StatusOK, gin.H{"categories": categories})
// }

func Index(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	models.DB.Find(&categories)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"categories": categories})
}


func Show(w http.ResponseWriter, r *http.Request) {
    // Extracting the "id_kategori" parameter from the URL path
    vars := mux.Vars(r)
    id := vars["id_kategori"]

    // Fetching the category from the database
    var category models.Category
    if err := models.DB.First(&category, "id_kategori = ?", id).Error; err != nil {
        // Handling the case where the category is not found
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Kategori tidak ditemukan"})
        return
    }

    // Responding with the fetched category in the JSON format using the helper
    helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"category": category})
}


func Create(w http.ResponseWriter, r *http.Request) {
    var category models.Category

    // Decode JSON request body into the 'category' variable
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        // Handling the case where the request body cannot be decoded
        helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
        return
    }

    // Create a new category in the database
    if err := models.DB.Create(&category).Error; err != nil {
        // Handling the case where the category creation fails
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat kategori"})
        return
    }

    // Responding with the created category in the JSON format
    helper.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"message": "Data Berhasil Dibuat"})
}


func Update(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    id := mux.Vars(r)["id_kategori"]

    // Decode JSON request body into the 'category' variable
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        // Handling the case where the request body cannot be decoded
        helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
        return
    }

    // Update the category in the database
    if models.DB.Model(&models.Category{}).Where("id_kategori = ?", id).Updates(&category).RowsAffected == 0 {
        // Handling the case where the update operation fails
        helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat mengupdate kategori"})
        return
    }

    // Responding with the updated category in the JSON format
    helper.ResponseJSON(w, http.StatusAccepted, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}


func Delete(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    id := mux.Vars(r)["id_kategori"]

    // Fetch the category from the database
    if err := models.DB.First(&category, "id_kategori = ?", id).Error; err != nil {
        // Handling the case where the category is not found
        helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Kategori tidak ditemukan"})
        return
    }

    // Delete the category from the database
    if err := models.DB.Delete(&category).Error; err != nil {
        // Handling the case where the delete operation fails
        helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus kategori"})
        return
    }

    // Responding with the success message in the JSON format
    helper.ResponseJSON(w, http.StatusNoContent, map[string]interface{}{"message": "Data berhasil dihapus"})
}
