package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	log "notes/backend/internal/logger"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis(addr string, password string) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	// Test the connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	// Reset database
	err = rdb.FlushDB(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to reset Redis database: %v", err)
	}

	return rdb, nil
}

type Session struct {
	UserID string `json:"userId"`
}

func SetSession(w http.ResponseWriter, userID string) {
	sessionID := "sess:" + generateSessionID()
	sessionData := Session{UserID: userID}

	// Simpan sesi ke Redis
	data, err := json.Marshal(sessionData) // Tambahkan error handling
	if err != nil {
		log.LogInfof("could not marshal session data: %v", err)
		http.Error(w, "could not set session", http.StatusInternalServerError)
		return
	}

	err = rdb.Set(context.Background(), sessionID, data, 10*time.Minute).Err() // Tambahkan error handling
	if err != nil {
		log.LogInfof("could not set session in redis: %v", err)
		http.Error(w, "could not set session", http.StatusInternalServerError)
		return
	}

	// Set cookie di browser
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(10 * time.Minute),
	})
}

func GetSession(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, err // Kembalikan error jika cookie tidak ada
	}

	sessionID := cookie.Value
	sessionData, err := rdb.Get(context.Background(), sessionID).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("session does not exist") // Jika sesi tidak ada
	} else if err != nil {
		return nil, err // Kembalikan error jika terjadi kesalahan
	}

	var session Session
	if err := json.Unmarshal([]byte(sessionData), &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.LogInfo("Entering SessionMiddleware")
		// Read the body of the request using io package
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read request body", http.StatusBadRequest)
			return
		}

		// Restore the request body using io.NopCloser
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Check if the operation is login or register
		if isLoginOrRegisterOperation(string(body)) {
			log.LogInfo("Skipping session validation for login/register")
			ctx := context.WithValue(r.Context(), "responseWriter", w)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Check session for other operations
		session, err := GetSession(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "session", session)
		fmt.Println("Session ID:", session.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateSessionID() string {
	return uuid.New().String() // Menghasilkan UUID sebagai session ID
}

func isLoginOrRegisterOperation(body string) bool {
	// Simple check for mutation operation containing login or register
	return strings.Contains(body, "mutation") &&
		(strings.Contains(body, "Login") || strings.Contains(body, "Register"))
}
