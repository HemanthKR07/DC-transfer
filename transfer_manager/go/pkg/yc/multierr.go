package yc

import (
	"fmt"
	"strings"
)

type multerr struct {
	errs []error
}

func (e *multerr) Errors() []error {
	return e.errs
}

func (e *multerr) Error() string {
	lines := make([]string, len(e.errs))
	for k, v := range e.errs {
		lines[k] = v.Error()
	}
	return strings.Join(lines, "\n")
}

func Errors(err error) []error {
	if err == nil {
		return nil
	}
	switch err := err.(type) {
	case interface {
		Errors() []error
	}:
		// go.uber.org/multierr
		return err.Errors()
	case interface {
		WrappedErrors() []error
	}:
		// github.com/hashicorp/go-multierror
		return err.WrappedErrors()
	default:
	}
	return []error{err}
}

// nolint:descriptive_errors
func Append(lhs, rhs error) error {
	if lhs == nil {
		//nolint:descriptiveerrors
		return rhs
	} else if rhs == nil {
		//nolint:descriptiveerrors
		return lhs
	}
	var result []error
	result = append(result, Errors(lhs)...)
	result = append(result, Errors(rhs)...)
	//nolint:descriptiveerrors
	return &multerr{result}
}

func CombineGoroutines(funcs ...func() error) error {
	errChan := make(chan error, len(funcs))
	for _, f := range funcs {
		go func(f func() error) {
			var err error
			defer func() {
				if r := recover(); r != nil {
					errChan <- fmt.Errorf("panic recovered: %v", r)
				} else {
					errChan <- err
				}
			}()
			err = f()
		}(f)
	}
	var errs error
	for i := 0; i < cap(errChan); i++ {
		errs = Append(errs, <-errChan)
	}
	return errs
}
