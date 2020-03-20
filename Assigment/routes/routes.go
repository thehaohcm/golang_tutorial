package routes

import (
	friend "Assigment/controllers/friend"
	subscribe "Assigment/controllers/subscribe"
	user "Assigment/controllers/user"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "Assigment/docs"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))

		r.Use(jwtauth.Authenticator)

		r.Post("/addFriend", friend.AddFriend)
		r.Get("/showAllUsers", user.ShowAllUser)
	})

	router.Group(func(r chi.Router) {
		// router.Post("/addFriend", friend.AddFriend)
		r.Get("/listFriend", friend.ShowListFriend)
		r.Post("/commonFriend", friend.ShowCommonFriends)
		r.Post("/subscribe", subscribe.Subscribe)
		r.Post("/blockSubscribe", subscribe.BlockSubscribe)
		r.Post("/listRecipients", subscribe.ShowListRecipients)

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition"
		))
	})
	// router.Get("/showAllUsers", user.ShowAllUser)
	// // router.Post("/addFriend", friend.AddFriend)
	// router.Get("/listFriend", friend.ShowListFriend)
	// router.Post("/commonFriend", friend.ShowCommonFriends)
	// router.Post("/subscribe", subscribe.Subscribe)
	// router.Post("/blockSubscribe", subscribe.BlockSubscribe)
	// router.Post("/listRecipients", subscribe.ShowListRecipients)

	// router.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition"
	// ))

	return router
}
