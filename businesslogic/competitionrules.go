// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
