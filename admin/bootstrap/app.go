package bootstrap

import (
	"github.com/herhe-com/framework/foundation"
	"project.io/shop/admin/config"
)

func Boot() {

	application := foundation.Application{}

	//Bootstrap the application.
	application.Boot()

	//Bootstrap the other service.
	config.Boot()
}
