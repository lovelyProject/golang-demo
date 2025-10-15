package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.Path,
		}).Info("request received")
		next.ServeHTTP(w, r)
	})
}
