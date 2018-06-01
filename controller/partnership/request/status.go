package request

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

type PartnershipRequestStatusServer struct {
	businesslogic.IPartnershipRequestStatusRepository
}

// GET /api/partnership/status
func (server PartnershipRequestStatusServer) GetPartnershipRequestStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := server.GetPartnershipRequestStatus()
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot retrieve partnership request status list", nil)
		return
	}

	data := make([]viewmodel.PartnershipRequestStatus, 0)
	for _, each := range status {
		view := viewmodel.PartnershipRequestStatus{
			ID:   each.ID,
			Name: each.Description,
		}
		data = append(data, view)
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
