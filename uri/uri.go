package uri

import "github.com/dong-tran/goinaction/controller"

func CreateControllers() []controller.Controller {
	var controllers []controller.Controller
	controllers = append(controllers, controller.CompanyController{})
	controllers = append(controllers, controller.ShopController{})
	controllers = append(controllers, controller.DummyController{})
	return controllers
}
