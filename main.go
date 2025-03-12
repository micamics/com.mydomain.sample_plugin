package main

import (
	"net/http"

	sdkapi "sdk/api"
)

func main() {}

func Init(api sdkapi.IPluginApi) {
	// Your plugin code here
	httpAPI := api.Http()

	// Register navigation menu items
	httpAPI.Navs().AdminNavsFactory(func(r *http.Request) []sdkapi.AdminNavItemOpt {
		return []sdkapi.AdminNavItemOpt{
			{
				Category:  sdkapi.NavCategorySystem,
				Label:     "Sample Plugin",
				RouteName: "settings.form",
			},
		}
	})
}
