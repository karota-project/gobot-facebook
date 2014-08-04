package facebook

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestFacebookDriver() *FacebookDriver {
	return NewFacebookDriver(NewFacebookAdaptor("myAdaptor"), "myDriver")
}

func TestFacebookDriverStart(t *testing.T) {
	d := initTestFacebookDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestFacebookDriverHalt(t *testing.T) {
	d := initTestFacebookDriver()
	gobot.Expect(t, d.Halt(), true)
}
