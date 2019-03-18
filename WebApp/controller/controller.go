// Pac// kage controller loads the routes for each of the controllers.
package controller

import (
	"github.com/huzhaer/qianxun/controller/login"
	// "github.com/CloudyKit/jet/loaders/httpfs"
	// "github.com/CloudyKit/jet/examples/asset_packaging/assets/templates"
	// "github.com/huzhaer/qianxun/controller/analyse"
	"github.com/huzhaer/qianxun/controller/static"
	"github.com/huzhaer/qianxun/controller/index"
)



func LoadRoutes() {
	static.Load()
	index.Load()
	login.Load()
}
