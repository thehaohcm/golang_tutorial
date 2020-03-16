package routes

import (
	friend "Assigment/controllers/friend"
	subscribe "Assigment/controllers/subscribe"
	user "Assigment/controllers/user"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "Assigment/docs"
)

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/showAllUsers", user.ShowAllUser)
	router.Post("/addFriend", friend.AddFriend)
	router.Get("/listFriend", friend.ShowListFriend)
	router.Post("/commonFriend", friend.ShowCommonFriends)
	router.Post("/subscribe", subscribe.Subscribe)
	router.Post("/blockSubscribe", subscribe.BlockSubscribe)
	router.Post("/listRecipients", subscribe.ShowListRecipients)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition"
	))

	return router
}
