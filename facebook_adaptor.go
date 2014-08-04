package facebook

import (
	"github.com/hybridgroup/gobot"
)

type FacebookAdaptor struct {
	gobot.Adaptor
}

func NewFacebookAdaptor(name string) *FacebookAdaptor {
	return &FacebookAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"facebook.FacebookAdaptor",
		),
	}
}

func (f *FacebookAdaptor) Connect() bool {
	return true
}

func (f *FacebookAdaptor) Finalize() bool {
	return true
}
