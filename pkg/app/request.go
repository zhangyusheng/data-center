package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/zhangyusheng/data-center/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Error(map[string]interface{}{}, err.String())
	}

	return
}
