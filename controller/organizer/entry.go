// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou

package organizer

import "net/http"

type OrganizerEntryServer struct {
}

func (server OrganizerEntryServer) CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (server OrganizerEntryServer) DeleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (server OrganizerEntryServer) SearchEntryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (server OrganizerEntryServer) UpdateEntryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
