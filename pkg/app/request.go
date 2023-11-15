package app

import (
	"gin_blog/pkg/logging"

	"github.com/beego/beego/v2/core/validation"
)

func MarkErr(errs []*validation.Error) {
	for _, err := range errs {
		logging.Info(err.Key, err.Message)
	}
}
