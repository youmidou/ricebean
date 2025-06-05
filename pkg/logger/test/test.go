package test

import (
	tests "github.com/sirupsen/logrus/hooks/test"
	"io"
	"ricebean/pkg/logger/interfaces"
	lwrapper "ricebean/pkg/logger/logrus"
)

// NewNullLogger creates a discarding logger and installs the test hook.
func NewNullLogger() (interfaces.Logger, *tests.Hook) {
	logger, hook := tests.NewNullLogger()
	logger.Out = io.Discard
	return lwrapper.NewWithFieldLogger(logger), hook
}
