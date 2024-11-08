package user

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	supa "github.com/nedpals/supabase-go"
)

type UserHandler interface {
	GetUserData(token string) (*supa.User, error)
}

type UserHandlerSupa struct {
	client *supa.Client
}

func NewSupabaseUserHandler(client *supa.Client) *UserHandlerSupa {
    return &UserHandlerSupa{client: client}
}

func (s *UserHandlerSupa) GetUserData(cookie string) (*supa.User, error) {

	token := cookie
	if strings.HasPrefix(token, "base64-") {
		token = strings.TrimPrefix(token, "base64-")
	}

	// Usuń białe znaki i sprawdź poprawność długości Base64
	token = strings.TrimSpace(token)
	if len(token)%4 != 0 {
		padding := 4 - (len(token) % 4)
		token += strings.Repeat("=", padding)
	}

	// Zdekoduj Base64
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}
    // Przekonwertuj na string i przekształć w mapę JSON
    var tokenData map[string]interface{}
    if err := json.Unmarshal(decoded, &tokenData); err != nil {
		fmt.Println(err)
        return nil, errors.New("Failed to unmarshal JSON token")
    }

    // Wyciągnij access_token z dekodowanego ciasteczka
    accessToken, ok := tokenData["access_token"].(string)
    if !ok || accessToken == "" {
		fmt.Println(ok)
        return nil, errors.New("Access token not found")
    }
	ctx := context.Background()
	user, err := s.client.Auth.User(ctx, accessToken)
	if err != nil {
		fmt.Println(ok)
		return nil, errors.New("Invalid token")
	}

	return user, nil
}