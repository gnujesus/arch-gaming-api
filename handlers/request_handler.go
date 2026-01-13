package handlers

import (
	"database/sql"
	"github.com/gnujesus/arch-gaming-api/config"
	"log"
	"net/http"
)

type RequestHandler struct {
	Service 
	Logger *log.Logger
}

func (*r RequestHandler) SubmitGameRequest(w http.ResponseWriter, r *http.Request) {

}
