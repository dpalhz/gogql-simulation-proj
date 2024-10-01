package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "notes/backend/internal/logger"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Session struct {
	UserID string `json:"userId"`
}

func SetSession(ctx context.Context, userID string) {
	w, ok := ctx.Value("responseWriter").(http.ResponseWriter)
	if !ok {
		log.LogInfo("could not retrieve response writer from context")
		return
	}

	sessionID := "auth:" + generateSessionID()
	sessionData := Session{UserID: userID}

	data, err := json.Marshal(sessionData)
	if err != nil {
		log.LogInfof("could not marshal session data: %v", err)
		http.Error(w, "could not set session", http.StatusInternalServerError)
		return
	}

	err = rdb.Set(context.Background(), sessionID, data, 10*time.Minute).Err()
	if err != nil {
		log.LogInfof("could not set session in redis: %v", err)
		http.Error(w, "could not set session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(10 * time.Minute),
	})

	log.LogInfo("session set successfully")
}

func GetSession(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, err
	}

	sessionID := cookie.Value
	sessionData, err := rdb.Get(context.Background(), sessionID).Result() // Menggunakan rdb
	if err == redis.Nil {
		return nil, fmt.Errorf("session does not exist")
	} else if err != nil {
		return nil, err
	}

	var session Session
	if err := json.Unmarshal([]byte(sessionData), &session); err != nil {
		return nil, err
	}

	return &session, nil
}



var ErrUnauthorized = fmt.Errorf("unauthorized: user is not authenticated")


func IsAuthenticated(ctx context.Context) (bool, error) {
    auth, ok := ctx.Value("auth").(bool)
    if !ok || !auth {
        return false, ErrUnauthorized
    }
    return true, nil
}


func generateSessionID() string {
	return uuid.New().String() 
}


