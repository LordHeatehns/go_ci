package routers

func (a *apiRegistry) HealthCheck() {
	a.app.Get("/health", a.hand.Health)
	a.app.Get("/ready", a.hand.Ready)
}
