package facebook

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestFacebookAdaptor() *FacebookAdaptor {
	return NewFacebookAdaptor("myAdaptor")
}

func TestFacebookAdaptorConnect(t *testing.T) {
	a := initTestFacebookAdaptor()
	gobot.Expect(t, a.Connect(), true)
}

func TestFacebookAdaptorFinalize(t *testing.T) {
	a := initTestFacebookAdaptor()
	gobot.Expect(t, a.Finalize(), true)
}
