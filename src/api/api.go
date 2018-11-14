package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tictactoe"
)

var GAME *tictactoe.TicTacToeGame

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	jsonResp, err := json.Marshal(payload)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResp)
}

func respondJSONError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, GAME)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	if GAME.State == tictactoe.InProgress {
		respondJSONError(w, http.StatusBadRequest, "Game is in progress")
		return
	}

	GAME = tictactoe.NewTicTacToeGame()
	resp := map[string]string{
		"message": "Game started. Choose your move and play",
	}
	respondJSON(w, http.StatusOK, resp)
}

func PlayGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	resp := PlayReq{}

	defer r.Body.Close()
	err := decoder.Decode(&resp)
	if err != nil {
		//FIXME Things here need to be handled more explicitly
		log.Print(err.Error())
		respondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = resp.Validate()
	if err != nil {
		log.Print(err.Error(), " ", resp)
		respondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = GAME.Play(resp.Col, resp.Row, resp.Move)
	if err != nil {
		log.Print(err.Error())
		respondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, GAME)
}


func Run() {
	GAME = tictactoe.NewTicTacToeGame()
	router := mux.NewRouter()
	router.HandleFunc("/game", GetGame).Methods("GET")
	router.HandleFunc("/game", CreateGame).Methods("POST")
	router.HandleFunc("/game", PlayGame).Methods("PUT")
	log.Fatal(http.ListenAndServe(":5000", router))
}
