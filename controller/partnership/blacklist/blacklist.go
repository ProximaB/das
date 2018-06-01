package blacklist

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/controller/util"
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"time"
)

type PartnershipBlacklistViewModel struct {
	Username string    `json:"user"`
	Since    time.Time `json:"since"`
}

type PartnershipRequestBlacklistServer struct {
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRequestBlacklistRepository
}

// GET /api/partnership/blacklist
func (server PartnershipRequestBlacklistServer) GetBlacklistedAccountHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	account, _ := util.GetCurrentUser(r, server.IAccountRepository)

	blacklist, err := businesslogic.GetBlacklistedAccountsForUser(account.ID, server.IAccountRepository, server.IPartnershipRequestBlacklistRepository)

	if err != nil {
		log.Errorf(ctx, "error in getting partnership blacklist for user: %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
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
