package goss

import (
	"github.com/wasanx25/goss/viewer"
)

func Run(body string) error {

	v := viewer.New(body)
	if err := v.Init(); err != nil {
		return err
	}
	v.Start()

	return nil
}
