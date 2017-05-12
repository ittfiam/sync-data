package sync

import (
	"strings"

	"sync-mysql/errors"
)

type Variables struct {
	Vars map[string]string `json:"vars"`
}

func (v *Variables) Save() error {
	return SaveAssetAsJSON(".variable", v)
}

func (v *Variables) Set(key, value string) {
	v.Vars[key] = value
}

func (v *Variables) GetValue(keys ...*string) error {

	for _, key := range keys {
		var ok bool

		if len(*key) > 2 {
			if strings.HasPrefix(*key, "$") ||
				strings.HasPrefix(*key, "#") ||
				strings.HasPrefix(*key, "%") {
				find := (*key)[1:]

				*key, ok = v.Vars[find]

				if !ok {
					return errors.NewError("variable %s not found.", *key)
				}
			}

		}
	}

	return nil

}

func NewVariables() (variables *Variables, err error) {
	variables = &Variables{
		Vars: make(map[string]string, 0),
	}

	exist, err := AssetExists(".variable")

	if err != nil {
		return
	}

	if exist {
		err = ReadAssetAsJSON(".variable", variables)
		return
	} else {
		err = variables.Save()
		return
	}
}
