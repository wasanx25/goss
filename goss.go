package goss

import (
	"github.com/wasanx25/goss/viewer"
)

func Run(text string) error {

	v := viewer.New(text)
	if err := v.Init(); err != nil {
		return err
	}

	if err := v.Start(); err != nil {
		return err
	}

	return nil
}
