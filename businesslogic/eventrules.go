package businesslogic

import "errors"

type EventAgeRule struct {
}

func (rule EventAgeRule) Apply(registration Registration) error {
	return errors.New("not implemented")
}
