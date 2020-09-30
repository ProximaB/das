package viewmodel

import "github.com/ProximaB/das/businesslogic"

// SearchEntryForm defines the acceptable parameters that can be used to query competition and event entries
type SearchEntryForm struct {
	CompetitionID int `schema:"competitionId"`
	EventID       int `schema:"eventId"`
	FederationID  int `schema:"federationId"`
	DivisionID    int `schema:"divisionId"`
	AgeID         int `schema:"ageId"`
	ProficiencyID int `schema:"proficiencyId"`
	StyleID       int `schema:"styleId"`
	PartnershipID int `schema:"partnershipId,omitempty"`
	AthleteID     int `schema:"athleteId,omitempty"`
}

// AthleteCompetitionEntryViewModel
type AthleteCompetitionEntryViewModel struct {
	EntryID       int                  `json:"id"`
	CompetitionID int                  `json:"competitionId"`
	Athlete       AthleteTinyViewModel `json:"athlete"`
	IsLead        bool                 `json:"isLead"`
	LeadTag       int                  `json:"leadTag"`
}

type CoupleCompetitionEntryViewModel struct {
	EntryID       int                      `json:"id"`
	CompetitionID int                      `json:"competitionId"`
	Couple        PartnershipTinyViewModel `json:"partnership"`
}

type AthleteEventEntryViewModel struct {
	EventID int                  `json:"eventId"`
	Athlete AthleteTinyViewModel `json:"athlete"`
}

type CoupleEventEntryViewModel struct {
	EntryID  int                  `json:"entryId"`
	EventID  int                  `json:"eventId"`
	CoupleID int                  `json:"partnershipId"`
	Lead     AthleteTinyViewModel `json:"lead"`
	Follow   AthleteTinyViewModel `json:"follow"`
}

type CompetitionEntryListViewModel struct {
	Competition    CompetitionViewModel               `json:"competition"`
	AthleteEntries []AthleteCompetitionEntryViewModel `json:"athleteEntries"`
	CoupleEntries  []CoupleCompetitionEntryViewModel  `json:"partnershipEntries"`
}

func AthleteCompetitionEntryToViewModel(entry businesslogic.AthleteCompetitionEntry) AthleteCompetitionEntryViewModel {
	return AthleteCompetitionEntryViewModel{
		EntryID:       entry.ID,
		CompetitionID: entry.Competition.ID,
		Athlete:       AthleteToTinyViewModel(entry.Athlete),
		IsLead:        entry.IsLead,
		LeadTag:       entry.LeadTag,
	}
}

func CompetitionEntriesToViewModel(entries businesslogic.CompetitionEntryList) CompetitionEntryListViewModel {
	view := CompetitionEntryListViewModel{}

	view.Competition = CompetitionDataModelToViewModel(entries.Competition, businesslogic.AccountTypeNoAuth)

	athletes := make([]AthleteCompetitionEntryViewModel, 0)
	for _, each := range entries.AthleteEntries {
		view := AthleteCompetitionEntryViewModel{
			EntryID:       each.ID,
			CompetitionID: each.Competition.ID,
			Athlete:       AthleteToTinyViewModel(each.Athlete),
		}
		athletes = append(athletes, view)
	}

	couples := make([]CoupleCompetitionEntryViewModel, 0)
	for _, each := range entries.CoupleEntries {
		view := CoupleCompetitionEntryViewModel{
			EntryID:       each.ID,
			CompetitionID: each.Competition.ID,
			Couple:        PartnershipToTinyViewModel(each.Couple),
		}
		couples = append(couples, view)
	}

	view.AthleteEntries = athletes
	view.CoupleEntries = couples
	return view
}

type EventEntryListViewModel struct {
	Event          EventViewModel               `json:"event"`
	AthleteEntries []AthleteEventEntryViewModel `json:"athleteEntries"`
	CoupleEntries  []CoupleEventEntryViewModel  `json:"partnershipEntries"`
}

func EventEntriesToViewModel(entries businesslogic.EventEntryList) EventEntryListViewModel {
	view := EventEntryListViewModel{}
	eventView := EventViewModel{}
	eventView.PopulateViewModel(entries.Event)
	view.Event = eventView
	view.CoupleEntries = CoupleEventEntryToViewModel(entries.CoupleEntries)
	return view
}

func CoupleEventEntryToViewModel(entries []businesslogic.PartnershipEventEntry) []CoupleEventEntryViewModel {
	output := make([]CoupleEventEntryViewModel, 0)
	for _, each := range entries {
		evtView := EventViewModel{}
		evtView.PopulateViewModel(each.Event)
		view := CoupleEventEntryViewModel{
			EntryID:  each.ID,
			EventID:  each.Event.ID,
			CoupleID: each.Couple.ID,
			Lead:     AthleteToTinyViewModel(each.Couple.Lead),
			Follow:   AthleteToTinyViewModel(each.Couple.Follow),
		}
		output = append(output, view)
	}
	return output
}
