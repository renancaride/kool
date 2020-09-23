package cmd

import (
	"io"
	"kool-dev/kool/cmd/shell"
)

// KoolService interface holds the contract for a
// general service which implements some bigger chunk
// of logic usually linked to a command.
type KoolService interface {
	Execute([]string) error

	shell.Exiter
	shell.OutputWriter
}

// DefaultKoolService holds handlers and functions shared by all
// services, meant to be used on commands when executing the services.
type DefaultKoolService struct {
	exiter shell.Exiter
	out    shell.OutputWriter
}

func newDefaultKoolService() *DefaultKoolService {
	return &DefaultKoolService{
		shell.NewExiter(),
		shell.NewOutputWriter(),
	}
}

// Exit proxies the call the given Exiter
func (k *DefaultKoolService) Exit(code int) {
	k.exiter.Exit(code)
}

// SetWriter proxies the call to the given OutputWriter
func (k *DefaultKoolService) SetWriter(w io.Writer) {
	k.out.SetWriter(w)
}

// Error proxies the call to the given OutputWriter
func (k *DefaultKoolService) Error(err error) {
	k.out.Error(err)
}

// Warning proxies the call to the given OutputWriter
func (k *DefaultKoolService) Warning(out ...interface{}) {
	k.out.Warning(out...)
}
