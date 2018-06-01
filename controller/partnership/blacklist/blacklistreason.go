package blacklist

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

type PartnershipRequestBlacklistReasonServer struct {
	businesslogic.IPartnershipRequestBlacklistReasonRepository
}

// GET /api/partnership/blacklist/reason
func (server PartnershipRequestBlacklistReasonServer) GetPartnershipBlacklistReasonHandler(w http.ResponseWriter, r *http.Request) {
	reasons, err := server.GetPartnershipRequestBlacklistReasons()
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot retrieve partnership request blacklist reason", nil)
		return
	}

	data := make([]viewmodel.PartnershipRequestStatus, 0)
	for _, each := range reasons {
		view := viewmodel.PartnershipRequestStatus{
			ID:   each.ID,
			Name: each.Name,
		}
		data = append(data, view)
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
