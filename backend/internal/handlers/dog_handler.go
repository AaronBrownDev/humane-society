package handlers

import "net/http"

func GetAllDogs(w http.ResponseWriter, r *http.Request)       {}
func GetDogByID(w http.ResponseWriter, r *http.Request)       {}
func CreateDog(w http.ResponseWriter, r *http.Request)        {}
func UpdateDog(w http.ResponseWriter, r *http.Request)        {}
func DeleteDog(w http.ResponseWriter, r *http.Request)        {}
func GetAvailableDogs(w http.ResponseWriter, r *http.Request) {}
