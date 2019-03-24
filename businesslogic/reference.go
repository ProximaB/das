package businesslogic

import (
	"errors"
	"time"
)

// Country specifies the data needed to serve as reference data
type Country struct {
	ID              int
	Name            string
	Abbreviation    string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// GetStates retrieves all the states that are associated with the caller Country from the repository
func (country Country) GetStates(repo IStateRepository) ([]State, error) {
	if repo != nil {
		return repo.SearchState(SearchStateCriteria{CountryID: country.ID})
	}
	return nil, errors.New("null IStateRepository")
}

// GetFederations retrieves all the federations that are associated with the caller Country from the repository
func (country Country) GetFederations(repo IFederationRepository) ([]Federation, error) {
	if repo != nil {
		return repo.SearchFederation(SearchFederationCriteria{CountryID: country.ID})
	}
	return nil, errors.New("null IFederationRepository")
}

// SearchCountryCriteria specifies the parameters that can be used to search certain countries in DAS
type SearchCountryCriteria struct {
	CountryID    int    `schema:"id"`
	Name         string `schema:"name"`
	Abbreviation string `schema:"abbreviation"`
}

// ICountryRepository specifies the functions that a repository needs to implement
type ICountryRepository interface {
	CreateCountry(country *Country) error
	SearchCountry(criteria SearchCountryCriteria) ([]Country, error)
	DeleteCountry(country Country) error
	UpdateCountry(country Country) error
}

// State defines the state/province of a country
type State struct {
	ID              int
	Name            string
	Abbreviation    string
	CountryID       int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// SearchStateCriteria defines the criteria that states can be searched by
type SearchStateCriteria struct {
	StateID   int    `schema:"id"`
	Name      string `schema:"name"`
	CountryID int    `schema:"country"`
}

type IStateRepository interface {
	CreateState(state *State) error
	SearchState(criteria SearchStateCriteria) ([]State, error)
	UpdateState(state State) error
	DeleteState(state State) error
}

func (state State) GetCities(repo ICityRepository) ([]City, error) {
	if repo == nil {
		return nil, errors.New("null ICityRepository")
	}
	return repo.SearchCity(SearchCityCriteria{StateID: state.ID})
}

// City contains data to represent a City object in DAS. City is associated with State, which is associated with Country
type City struct {
	ID              int
	Name            string
	StateID         int
	CreateUserID    *int // 2017-10-17 yubing24: use pointer so that if the value is nil, it will be ignored
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// GetSchools retrieves all schools that are in the caller city and from the repository
func (city City) GetSchools(repo ISchoolRepository) ([]School, error) {
	if repo != nil {
		return repo.SearchSchool(SearchSchoolCriteria{CityID: city.ID})
	}
	return nil, errors.New("null ISchoolRepository")
}

// GetStudios retrieves all the studios that are in the caller city and from the repository
func (city City) GetStudios(repo IStudioRepository) ([]Studio, error) {
	if repo != nil {
		return repo.SearchStudio(SearchStudioCriteria{CityID: city.ID})
	}
	return nil, errors.New("null IStudioRepository")
}

// SearchCityCriteria provides the parameter for search City in ICityRepository. This criteria can be used as
// parameters in REST API or internally
type SearchCityCriteria struct {
	CityID  int    `schema:"id"`
	Name    string `schema:"name"`
	StateID int    `schema:"state"`
}

// ICityRepository specifies the interface that data access layer code should implement
type ICityRepository interface {
	CreateCity(city *City) error
	SearchCity(criteria SearchCityCriteria) ([]City, error)
	UpdateCity(city City) error
	DeleteCity(city City) error
}

type Federation struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	YearFounded     int
	CountryID       int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchFederationCriteria struct {
	ID        int    `schema:"id"`
	Name      string `schema:"name"`
	CountryID int    `schema:"country"`
}

type IFederationRepository interface {
	CreateFederation(federation *Federation) error
	SearchFederation(criteria SearchFederationCriteria) ([]Federation, error)
	UpdateFederation(federation Federation) error
	DeleteFederation(federation Federation) error
}

func (federation Federation) GetDivisions(repo IDivisionRepository) ([]Division, error) {
	if repo == nil {
		return nil, errors.New("null IDivisionRepository")
	}
	return repo.SearchDivision(SearchDivisionCriteria{FederationID: federation.ID})
}

type Division struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	FederationID    int
	Note            string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchDivisionCriteria struct {
	ID           int    `schema:"id"`
	Name         string `schema:"name"`
	FederationID int    `schema:"federation"`
}

type IDivisionRepository interface {
	CreateDivision(division *Division) error
	SearchDivision(criteria SearchDivisionCriteria) ([]Division, error)
	UpdateDivision(division Division) error
	DeleteDivision(division Division) error
}

// Age contains data for the age category requirement for events
// Age is associated with Division, which is associated with Federation
type Age struct {
	ID              int
	Name            string
	Description     string
	DivisionID      int
	Enforced        bool // if required, AgeMinimum and AgeMaximum must have non-zero value
	AgeMinimum      int
	AgeMaximum      int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// SearchAgeCriteria provides parameters when searching Age in IAgeRepository
type SearchAgeCriteria struct {
	Name       string `schema:"name"`
	DivisionID int    `schema:"division"`
	AgeID      int    `schema:"id"`
}

// IAgeRepository provides an interface for other businesslogic code to access Age data
type IAgeRepository interface {
	CreateAge(age *Age) error
	SearchAge(criteria SearchAgeCriteria) ([]Age, error)
	UpdateAge(age Age) error
	DeleteAge(age Age) error
}

type Proficiency struct {
	ID              int
	Name            string
	Description     string
	DivisionID      int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTImeUpdated time.Time
}

type SearchProficiencyCriteria struct {
	ProficiencyID int    `schema:"id"`
	DivisionID    int    `schema:"division"`
	Name          string `schema:"name"`
}

type IProficiencyRepository interface {
	SearchProficiency(criteria SearchProficiencyCriteria) ([]Proficiency, error)
	CreateProficiency(proficiency *Proficiency) error
	UpdateProficiency(proficiency Proficiency) error
	DeleteProficiency(proficiency Proficiency) error
}

type Style struct {
	ID              int
	Name            string
	Description     string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchStyleCriteria struct {
	StyleID int    `schema:"id"`
	Name    string `schema:"name"`
}

type IStyleRepository interface {
	CreateStyle(style *Style) error
	SearchStyle(criteria SearchStyleCriteria) ([]Style, error)
	UpdateStyle(style Style) error
	DeleteStyle(style Style) error
}

// Dance is style-dependent. For example, Cha Cha of American Rhythm is different from Cha Cha of International Latin
type Dance struct {
	ID              int
	Name            string
	Description     string
	Abbreviation    string
	StyleID         int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// SearchDanceCriteria specifies the parameters that can be used to to search dances in DAS
type SearchDanceCriteria struct {
	StyleID int    `schema:"style"`
	DanceID int    `schema:"id"`
	Name    string `schema:"name"`
}

// IDanceRepository specifies the interface that needs to be implemented to functions as a repository for dance
type IDanceRepository interface {
	CreateDance(dance *Dance) error
	SearchDance(criteria SearchDanceCriteria) ([]Dance, error)
	UpdateDance(dance Dance) error
	DeleteDance(dance Dance) error
}

// ByDanceID allows sort a slice of dances by their IDs
type ByDanceID []Dance

func (d ByDanceID) Len() int           { return len(d) }
func (d ByDanceID) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDanceID) Less(i, j int) bool { return d[i].ID < d[j].ID }

type School struct {
	ID              int
	Name            string
	CityID          int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchSchoolCriteria struct {
	ID      int    `schema:"id"`
	Name    string `schema:"name"`
	CityID  int    `schema:"city"`
	StateID int    `schema:"state"`
}

type ISchoolRepository interface {
	CreateSchool(school *School) error
	SearchSchool(criteria SearchSchoolCriteria) ([]School, error)
	UpdateSchool(school School) error
	DeleteSchool(school School) error
}

type Studio struct {
	ID              int
	Name            string
	Address         string
	CityID          int
	Website         string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchStudioCriteria struct {
	ID        int    `schema:"id"`
	Name      string `schema:"name"`
	CityID    int    `schema:"city"`
	StateID   int    `schema:"state"`
	CountryID int    `schema:"country"`
}

type IStudioRepository interface {
	CreateStudio(studio *Studio) error
	SearchStudio(criteria SearchStudioCriteria) ([]Studio, error)
	DeleteStudio(studio Studio) error
	UpdateStudio(studio Studio) error
}

const (
	GENDER_MALE    = 2
	GENDER_FEMALE  = 1
	GENDER_UNKNOWN = 3 // registering a new account no longer requires specifying gender
)

type IGenderRepository interface {
	GetAllGenders() ([]Gender, error)
}

type Gender struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}
