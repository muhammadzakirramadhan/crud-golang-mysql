package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/muhammadzakirramadhan/crud-golang-mysql/mahasiswa"
	"github.com/muhammadzakirramadhan/crud-golang-mysql/models"
	"github.com/muhammadzakirramadhan/crud-golang-mysql/utils"
)

/**
* Setting Port Disini
 */
const PORT = "8090"

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		mahasiswas, err := mahasiswa.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, mahasiswas, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

// Input Data Mahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Mahasiswa

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.Insert(ctx, mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Update Data Mahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Mahasiswa

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		fmt.Println(mhs)

		if err := mahasiswa.Update(ctx, mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Delete Data Mahasiswa
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
 
    if r.Method == "DELETE" {
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var mhs models.Mahasiswa
 
        id := r.URL.Query().Get("id")
 
        if id == "" {
            utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
            return
        }
        mhs.ID, _ = strconv.Atoi(id)
 
        if err := mahasiswa.Delete(ctx, mhs); err != nil {
 
            kesalahan := map[string]string{
                "error": fmt.Sprintf("%v", err),
            }
 
            utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

func main() {

	http.HandleFunc("/api", GetMahasiswa)
	http.HandleFunc("/api/create", PostMahasiswa)
	http.HandleFunc("/api/update", UpdateMahasiswa)
	http.HandleFunc("/api/delete", DeleteMahasiswa)

	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server Running On Port : " + PORT)
}

