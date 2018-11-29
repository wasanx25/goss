package goss

import (
	"github.com/wasanx25/goss/viewer"
)

func Run(text string) error {

	v := viewer.New(text)
	if err := v.Init(); err != nil {
		return err
	}
	v.Start()

	return nil
}
