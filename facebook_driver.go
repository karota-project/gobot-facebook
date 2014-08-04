package facebook

import (
  "github.com/hybridgroup/gobot"
)

type FacebookDriver struct {
  gobot.Driver
}

type FacebookInterface interface {
}

func NewFacebookDriver(a *FacebookAdaptor, name string) *FacebookDriver {
  return &FacebookDriver{
    Driver: *gobot.NewDriver(
      name,
      "facebook.FacebookDriver",
      a,
    ),
  }
}

func (f *FacebookDriver) adaptor() *FacebookAdaptor {
  return f.Driver.Adaptor().(*FacebookAdaptor)
}

func (f *FacebookDriver) Start() bool { return true }
func (f *FacebookDriver) Halt() bool { return true }
