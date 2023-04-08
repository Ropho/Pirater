package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (serv *Server) authenticateUser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := serv.SessionStore.Get(r, serv.Config.Env.SessionName)
		if err != nil {
			serv.error(w, r, http.StatusInternalServerError, "SESSION ERROR")
			logrus.Error("session get error: ", err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "UNABLE TO AUTH")
			logrus.Error("GET USER_ID ERROR: ", err)
			return
		}

		// logrus.Info(id)

		u, err := serv.Store.User().FindById(id.(int))
		if err != nil {
			serv.error(w, r, http.StatusNetworkAuthenticationRequired, "UNABLE TO AUTH")
			logrus.Error("FIND USER BY ID ERROR: ", err)
			return
		}

		logrus.Info("AUTHENTICATE USER GOOD")

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *Server) setRequestId(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestId, id)))
	})
}

func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.Logger.WithFields(logrus.Fields{
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
