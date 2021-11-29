package health

import (
	"encoding/json"
	"net/http"

	"github.com/shipperizer/miniature-monkey/types"
)

func Status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := new(types.DataResponse)
		resp.Message = "ok"

		json.NewEncoder(w).Encode(resp)
	}
}
