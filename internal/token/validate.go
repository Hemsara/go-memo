package internal

import (
	"calendar_automation/pkg/initializers"
	"context"
	"fmt"
	"os"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ctx = context.Background()

func Validate(str string) (isValid bool, token *jwt.Token, err error, claims jwt.MapClaims) {
	token, err = jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing algorithm was used")
		}
		var secretKey string

		secretKey = os.Getenv("ADMIN_SECRET")

		return []byte(secretKey), nil
	})

	if err != nil {
		return false, nil, err, nil
	}

	clm, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, nil, fmt.Errorf("invalid token"), nil
	}
	sub, ok := clm["sub"].(string)
	if !ok {
		return false, nil, fmt.Errorf("subject claim is not a string"), nil
	}
	exist := ValidateFromRedis(sub)

	if !exist {
		return false, nil, fmt.Errorf("token access revoked"), nil
	}

	return true, token, nil, clm
}

func ValidateFromRedis(uuid string) (isValid bool) {
	_, err := initializers.RD.Get(ctx, uuid).Result()

	if err != nil {
		return false
	}

	return true

}

func CreateToken(
	email string,
	duration time.Duration,
) (id uuid.UUID, token string, err error) {
	now := time.Now().UTC()

	id, err = uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, "", err
	}

	claims := make(jwt.MapClaims)

	claims["sub"] = id.String()
	claims["exp"] = now.Add(duration).Unix()
	claims["email"] = email

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("ADMIN_SECRET")))
	if err != nil {
		return uuid.UUID{}, "", err
	}
	if err := SaveTokenToRedis(id.String(), token, duration); err != nil {
		return uuid.UUID{}, "", err
	}
	return id, token, nil
}

func SaveTokenToRedis(key, token string, expiration time.Duration) error {
	client := initializers.RD
	return client.Set(ctx, key, token, expiration).Err()
}
