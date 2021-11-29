package check

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shipperizer/miniature-monkey/webiface"
	"go.uber.org/zap"
)

type Blueprint struct {
	logger *zap.SugaredLogger
}

func (bp *Blueprint) Routes(router *mux.Router) {
	router.HandleFunc("/api/v1/check", bp.check).Methods(http.MethodGet, http.MethodOptions).Name("check")
}

func (bp *Blueprint) check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := new(CheckResponse)
	resp.Status = http.StatusOK
	resp.Message = "purple bro"

	json.NewEncoder(w).Encode(resp)
}

func NewBlueprint(logger *zap.SugaredLogger) webiface.BlueprintInterface {
	bp := new(Blueprint)

	bp.logger = logger
	if bp.logger == nil {
		panic("empty logger in blueprint")
	}

	return bp
}
