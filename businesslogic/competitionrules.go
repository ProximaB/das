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
