package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/danilocordeirodev/social-go/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	// comments, err := app.store.Comments.GetByPostID(r.Context(), post.ID)
	// if err != nil {
	// 	app.internalServerError(w, r, err)
	// 	return
	// }

	// post.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) usersContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "userID")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
		ctx := r.Context()

		post, err := app.store.Users.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(postCtx).(*store.User)
	return user
}
