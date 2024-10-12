package utility

import (
	"fmt"
	"net/http"
    "encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"strings"
)

func CreateToken(username string) (string, error) {
    // Create the claims
	secretKey := []byte("yo yo yo yo jessy pink man") // Change this to your secret key
    claims := jwt.MapClaims{
        "sub": username,
        "iss": "academia",
        "aud": "student",
        "exp": time.Now().Add(time.Hour).Unix(),
        "iat": time.Now().Unix(),
    }

    // Create a new token object
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with a secret key
    signedToken, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}

//veryfy token
// -- > learn type assertion in go
func VerifyToken(r *http.Request) (string, error) {
    secretKey := []byte("yo yo yo yo jessy pink man") // Change this to your secret key
    authHeader := r.Header.Get("Authorization")
    
    if authHeader == "" {
        return "", fmt.Errorf("authorization header is missing")
    }

    // Split the header to get the token
    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return "", fmt.Errorf("invalid authorization format")
    }

    tokenString := parts[1]
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate the algorithm
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        return "", err
    }

    // Check if the token is valid
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Serialize claims to JSON
        claimsJSON, err := json.Marshal(claims)
        if err != nil {
            return "", err
        }
        return string(claimsJSON), nil
    }

    return "", fmt.Errorf("invalid token")
}