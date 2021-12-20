package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) route() http.Handler {
	middleware.DefaultLogger = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			app.logger.Debugw(r.URL.RequestURI(),
				"addr", r.RemoteAddr,
				"protocol", r.Proto,
				"method", r.Method,
			)
			next.ServeHTTP(w, r)
		})
	}

	r := chi.NewMux()

	r.Use(middleware.Logger, middleware.Recoverer, middleware.StripSlashes)
	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", app.loginHandler)
		r.Post("/refresh", app.refreshTokenHandler)
		r.Post("/logout", app.logoutUserHandler)
	})

	r.With(app.isHelp).Route("/", func(r chi.Router) {
		r.With(app.auth).Route("/user", func(r chi.Router) {
			r.With(app.userCtx).Route("/", func(r chi.Router) {
				r.Get("/", app.getUserProfileHandler)
				r.Put("/", app.updateUserProfileHandler)

				r.Route("/room", func(r chi.Router) {
					r.Get("/opened", app.getOpenedRoomsListHandler)
					r.Get("/closed", app.getClosedRoomsListHandler)

					r.Post("/", app.createRoomHandler)
					r.Post("/join", app.joinToRoomHandler)

					r.With(app.roomIdCtx).Route("/{room_id}", func(r chi.Router) {
						r.Put("/close", app.setCloseRoomHandler)
						r.Delete("/", app.deleteRoomHandler)

						r.With(app.roomNotClosed).Route("/", func(r chi.Router) {
							r.With(app.isRoomParticipants).Route("/", func(r chi.Router) {
								r.Post("/code", app.getRoomCodeHandler)
								r.With(app.isRoomOwner).With(app.participantIdCtx).Delete("/participant/{participant_id}", app.deleteRoomParticipantHandler)
								r.Delete("/leave_room", app.leaveRoomHandler)

								r.With(app.isRoomParticipants).Route("/purchase", func(r chi.Router) {
									r.Post("/", app.addPurchaseHandler)

									r.With(app.purchaseIdCtx).Route("/{purchase_id}", func(r chi.Router) {
										r.With(app.isPurchaseOwner).Route("/", func(r chi.Router) {
											r.Put("/", app.updatePurchaseHandler)
											r.Delete("/", app.deletePurchaseHandler)
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})

	return r
}
