package nbarouter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/johnatansantana/go/Api_NBA/config/dao"
	. "github.com/johnatansantana/go/Api_NBA/models"
	"gopkg.in/mgo.v2/bson"
)

var dao = NbaDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	nbas, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, nbas)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nba, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid NBA ID")
		return
	}
	respondWithJson(w, http.StatusOK, nba)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var nba Nba
	if err := json.NewDecoder(r.Body).Decode(&nba); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	nba.ID = bson.NewObjectId()
	if err := dao.Create(nba); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, nba)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var nba Nba
	if err := json.NewDecoder(r.Body).Decode(&nba); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], nba); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": nba.Player + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
