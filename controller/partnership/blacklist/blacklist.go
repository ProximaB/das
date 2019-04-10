package blacklist

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
	"time"
)

type PartnershipBlacklistViewModel struct {
	Username string    `json:"user"`
	Since    time.Time `json:"since"`
}

type PartnershipRequestBlacklistServer struct {
	auth.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRequestBlacklistRepository
}

// GET /api/partnership/blacklist
func (server PartnershipRequestBlacklistServer) GetBlacklistedAccountHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := server.GetCurrentUser(r)

	blacklist, err := account.GetBlacklistedAccounts(server.IAccountRepository, server.IPartnershipRequestBlacklistRepository)

	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]PartnershipBlacklistViewModel, 0)
	for _, each := range blacklist {
		entry := PartnershipBlacklistViewModel{
			Username: each.FirstName + " " + each.LastName,
		}
		data = append(data, entry)
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}

// POST /api/partnership/blacklist/report
func (server PartnershipRequestBlacklistServer) CreatePartnershipRequestBlacklistReportHandler(w http.ResponseWriter, r *http.Request) {

}
