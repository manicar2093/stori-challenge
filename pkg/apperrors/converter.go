package apperrors

import (
	"fmt"
	"net/http"

	"github.com/manicar2093/stori-challenge/pkg/logger"
	"github.com/manicar2093/stori-challenge/pkg/validator"

	"github.com/coditory/go-errors"

	"github.com/labstack/echo/v4"
)

type (
	BasicError struct {
		Code  int         `json:"code"`
		Error interface{} `json:"error"`
	}
	ValidationError struct {
		Code                       int `json:"code"`
		*validator.ValidationError `json:",inline"`
	}
)

func HandlerWEcho(err error, ctx echo.Context) {
	code, body := handleErrorType(err) //nolint:varnamelen
	stack, ok := hasStackTrace(err)
	if ok {
		logger.GetLogger().WithField("stack", stack).Error(err)
	} else {
		logger.GetLogger().Error(err)
	}
	ctx.JSON(code, body) //nolint:errcheck
}

func handleErrorType(err error) (int, interface{}) {
	var (
		code    = http.StatusInternalServerError
		iterErr = err
	)

	for iterErr != nil {
		switch hErr := iterErr.(type) {
		case *validator.ValidationError:
			return http.StatusBadRequest, ValidationError{
				Code:            http.StatusBadRequest,
				ValidationError: hErr,
			}
		case HandleableError:
			return hErr.StatusCode(), BasicError{
				Code:  hErr.StatusCode(),
				Error: hErr.Error(),
			}
		}
		iterErr = errors.Unwrap(iterErr)
		continue
	}
	return code, BasicError{
		Code:  code,
		Error: err.Error(),
	}
}

func hasStackTrace(err error) ([]string, bool) {
	var stack []string
	stacked, ok := err.(*errors.Error)
	if !ok {
		return nil, false
	}
	for _, d := range stacked.StackTrace() {
		stack = append(stack, fmt.Sprintf("%s:%d", d.RelFileName(), d.FileLine()))
	}
	return stack, true
}
