package app

import "github.com/armuh16/otpcallassesment/assesment"

func mapUrls() {
	// router.GET("/assesment1", users.GetUser)
	// router.GET("/assesment2", users.SearchUser)
	router.GET("/assesment3", assesment.JSON)
}
