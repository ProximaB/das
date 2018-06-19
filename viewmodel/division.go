// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

type SearchDivisionViewModel struct {
	DivisionID   int    `schema:"id"`
	Name         string `schema:"name"`
	FederationID int    `schema:"federation"`
}

type DivisionViewModel struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Federation int    `json:"federation"`
}
