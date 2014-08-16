package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-facebook"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	facebookAdaptor := facebook.NewFacebookAdaptor("facebook-a01")
	facebookDriver := facebook.NewFacebookDriver(facebookAdaptor, "facebook-d01")

	master.AddRobot(
		gobot.NewRobot(
			"facebook",
			[]gobot.Connection{facebookAdaptor},
			[]gobot.Device{facebookDriver},
			func() {
				fmt.Println("work")
			}))

	master.Start()
}
