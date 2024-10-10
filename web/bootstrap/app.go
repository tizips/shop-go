package bootstrap

import (

	//Containers and other services must be started immediately
	"github.com/herhe-com/framework/foundation"

	//Delayed startup of other services init
	"project.io/shop/web/config"
)

func Boot() {

	application := foundation.Application{}

	//Bootstrap the application.
	application.Boot()

	//Bootstrap the other service.
	config.Boot()
}
