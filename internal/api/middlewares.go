package api

import (
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"text-converter/internal/cfg"
)

func requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := cfg.GetConfig()
		config.RequestId = uuid.NewV4().String()
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := cfg.GetConfig()
		cfg.NewLogger()
		cfg.BindFields(logrus.Fields{
			"requestUri": r.RequestURI,
			"requestMethod": r.Method,
			"requestId": config.RequestId,
		})
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := cfg.GetConfig()
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		if pair[0] != config.Username || pair[1] != config.Password {
			cfg.Logger.WithFields(logrus.Fields{
				"remoteUser": pair[0],
			}).Warning("not_authorized")
			http.Error(w, "Not authorized", 401)
			return
		}

		cfg.BindFields(logrus.Fields{"remoteUser": pair[0]})

		next.ServeHTTP(w, r)
	})
}