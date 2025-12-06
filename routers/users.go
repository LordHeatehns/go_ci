package routers

func (a *apiRegistry) UsersAPi() {
	api := a.app.Group("/api/v1" + "/users")
	api.Post("/get/users", a.hand.GetUsers)
	api.Post("/get/users/test", a.hand.GetUsersTest)

}
