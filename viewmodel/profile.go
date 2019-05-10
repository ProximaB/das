package viewmodel

import "time"

// SearchAthleteProfileForm
type SearchAthleteProfileForm struct {
	CountryId int    `schema:"countryId" validate:"min=0"`
	StateId   int    `schema:"stateId" validate:"min=0"`
	FirstName string `schema:"firstName" validate:""`
}

type AthleteSummaryProfile struct {
	UID       string
	FirstName string
	LastName  string
	State     string
	Country   string
}

type RatingProfile struct {
	Category    string // Style of dances: Standard, Latin, Smooth, Rhythm
	Rating      float32
	Rank        int
	LastUpdated time.Time
	History     []struct {
		Rating    float32
		Rank      int
		UpdatedOn time.Time
	}
}

type AthleteDetailedProfile struct {
	UID       string
	FirstName string
	LastName  string
	Country   string
	State     string
	Ratings   []RatingProfile
}

// SearchPartnershipProfileForm
type SearchPartnershipProfileForm struct {
	CountryId       int    `schema:"countryId" validate:"min=0"`
	StateId         int    `schema:"stateId" validate:"min=0"`
	LeadFirstName   string `schema:"leadFirstName"`
	LeadLastName    string `schema:"leadLastName"`
	FollowFirstName string `schema:"followFirstName"`
	FollowLastName  string `schema:"followLastName"`
}
