package middlewares

import (
	"App/internal/models"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func TransactionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tx := models.InitGorm.Begin()

		contextWithTx := r.Context()
		contextWithTx = context.WithValue(contextWithTx, "tx", tx)

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r.WithContext(contextWithTx))

		if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			tx.Commit()
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}
