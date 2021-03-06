package partnership

import (
	"encoding/json"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"github.com/ProximaB/das/viewmodel"
	"log"
	"net/http"
)

type PartnershipRoleServer struct {
	businesslogic.IPartnershipRoleRepository
}

func (server PartnershipRoleServer) GetPartnershipRolesHandler(w http.ResponseWriter, r *http.Request) {
	roles, err := server.IPartnershipRoleRepository.GetAllPartnershipRoles()
	if err != nil {
		log.Println(err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "an error occurred while reading the data", nil)
	}

	view := make([]viewmodel.PartnershipRole, 0)
	for _, each := range roles {
		view = append(view, viewmodel.PartnershipRoleDataModelToViewModel(each))
	}
	output, _ := json.Marshal(view)
	w.Write(output)
}
