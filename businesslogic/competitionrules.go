// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "errors"

type IRule interface {
	Apply(registration Registration) error
}

type GenderRule struct {
}

func (rule GenderRule) Apply(registration Registration) error {
	return errors.New("not implemented")
}
