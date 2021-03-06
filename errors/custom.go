package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

/**

 * Created by cxnky on 24/08/2018 at 14:21
 * errors
 * https://github.com/cxnky/

**/

// NewError is a custom error which pretty prints an error
func NewError(message string) error {

	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	return fmt.Errorf("[%s:%d] ERROR: %s\n", file, line, message)

}
