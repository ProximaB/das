// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package businesslogic

import "errors"

type IRule interface {
	Apply(registration EventRegistration) error
}

type GenderRule struct {
	AllowSameSex bool
	IAccountRepository
	IPartnershipRepository
}

func (rule GenderRule) Apply(registration EventRegistration) error {
	if partnershipResults, err := rule.SearchPartnership(SearchPartnershipCriteria{
		PartnershipID: registration.PartnershipID,
	}); err != nil {
		return err
	} else {
		partnership := partnershipResults[0]
		if partnership.SameSex && (!rule.AllowSameSex) {
			return errors.New("same sex is not allowed")
		}
	}
	return nil
}
