package auth

import(
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

//Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

//GenerateToken creates jwt token, returns a token and error
func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(2000 * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

//HandleToken decodes the jwt token for authentication
func HandleToken(tkn string) (string, error) {
	claims := &Claims{}

	tk, err := jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !tk.Valid {
		return "", err
	}

	return claims.UserID, nil
}