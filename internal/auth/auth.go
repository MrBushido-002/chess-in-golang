package auth

import(
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
    "net/http"
    "errors"
    "strings"
)



func MakeJWT(PlayerID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
    claims := jwt.RegisteredClaims{
        Issuer:    "chess-in-golang",
        IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
        ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
        Subject:   PlayerID.String(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(tokenSecret))
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
    token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(tokenSecret), nil
    })
	if err != nil {
		return uuid.Nil, err
	}

    if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
        id, err := uuid.Parse(claims.Subject)
        return id, err
    }
    return uuid.Nil, err
}

func GetBearerToken(headers http.Header) (string, error) {
    authHeader := headers.Get("Authorization")
    if authHeader == "" {
        return "", errors.New("authorization header missing")
    }

    parts := strings.Fields(authHeader)
    if len(parts) == 2 && parts[0] == "Bearer" {
        return parts[1], nil
    }
    return "", errors.New("invalid authorization header format")
}

func AuthenticateRequest(r *http.Request, secret string) (uuid.UUID, error) {

	token, err := GetBearerToken(r.Header)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := ValidateJWT(token, secret)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
