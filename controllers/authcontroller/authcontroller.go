package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/06202003/apiInventory/config"

	"github.com/06202003/apiInventory/helper"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	"github.com/06202003/apiInventory/models"
	"golang.org/x/crypto/bcrypt"
)


func Login(w http.ResponseWriter, r *http.Request) {
	// retrieve json input
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// retrieve user data based on Email and verify if the entered email matches the stored email 
	var user models.User
	if err := models.DB.Where("Email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email atau password salah"} //bisa aja email salah
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// retrieve the static salt from the user model
	staticSalt := "WITAsik"

	// combine entered password and salt
	passwordWithSalt := userInput.Password + staticSalt

	// verify if the entered password matches the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordWithSalt)); err != nil {
		response := map[string]string{"message": "Email atau password salah"} //bisa aja password salah
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(24 * time.Hour) // Set expiration to 24 hours
	claims := &config.JWTClaim{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: expTime.Unix(),
		},
	}

	// declares the algorithm that will be used for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		MaxAge:   24 * 60 * 60,
	})

	response := map[string]string{"message": "login berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}


func Register(w http.ResponseWriter, r *http.Request) {
	// retrieve json input
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// static salt
	staticSalt := "WITAsik"

	// combine password and salt
	passwordWithSalt := userInput.Password + staticSalt

	// hash pass using bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		response := map[string]string{"message": "Error hashing password"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// update password with hashed password
	userInput.Password = string(hashPassword)

	// insert into the database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Successfully created an account"}
	helper.ResponseJSON(w, http.StatusCreated, response)
}


func Logout(w http.ResponseWriter, r *http.Request) {
	// delete the token in the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
