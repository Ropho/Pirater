package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	user "github.com/Ropho/Pirater/internal/model/user"
)

func (serv *Server) authenticateUser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := serv.Store.Cookie().Get(r, serv.Config.Env.SessionName)
		if err != nil {
			serv.error(w, r, http.StatusInternalServerError, "SESSION ERROR")
			serv.Logger.Errorf("session get error: [%w]", err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "UNABLE TO AUTH")
			serv.Logger.Error("Authentication: GET USER_ID ERROR")
			return
		}

		u, err := serv.Store.User().FindById(id.(int))
		if err != nil {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "UNABLE TO AUTH")
			serv.Logger.Errorf("Authentication: FIND USER BY ID ERROR: [%w]", err)
			return
		}

		serv.Logger.Info("AUTHENTICATE USER GOOD")

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (serv *Server) authorizeAdmin(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(ctxKeyUser).(*user.User)
		if !ok {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "not found user")
			serv.Logger.Error("not found user")
			return
		}

		if u.Right != user.Admin {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "UNABLE TO AUTHORIZE ADMIN")
			serv.Logger.Error("not admin: ", u)
			return
		}
		serv.Logger.Info("Authorize Admin GOOD")

		next.ServeHTTP(w, r)
	})
}

func (serv *Server) setRequestId(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestId, id)))
	})
}

func (serv *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := serv.Logger.WithFields(log.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestId),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		logger.Infof("completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start))

		fmt.Println("")
	})
}
