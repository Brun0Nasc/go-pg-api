package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
		issuer:    os.Getenv("JWT_ISSUER"),
	}
}

func (s *jwtService) GenerateToken(id uint, nome string) (string, error) {
	// Define as claims (dados contidos no token)
	claims := jwt.MapClaims{
		"id":   id,
		"nome": nome,
		"exp":  time.Now().Add(time.Hour * 2).Unix(), // Token expira em 2 horas
		"iss":  s.issuer,
		"iat":  time.Now().Unix(),
	}

	// Cria um token com as claims e assina com a chave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *jwtService) ParseToken(tokenString string) (*jwt.Token, error) {
	// Faz o parsing do token JWT usando a chave secreta
	// e retorna um token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica o algoritmo de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Retorna a chave secreta
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Verifica se o token é válido
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	// Faz o cast das claims para o tipo jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to cast claims to MapClaims")
	}

	// Acessa a informação "id" do token JWT
	id, ok := claims["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("failed to cast id to float64")
	}

	fmt.Println(id) // imprime o ID do usuário do token JWT

	return token, nil
}

// // ValidateToken verifica se um token é válido ou não.
// // Retorna um booleano indicando se o token é válido ou não.
// func (s *jwtService) ValidateToken(tokenString string) (bool, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(s.secretKey), nil
// 	})

// 	if err != nil {
// 		return false, err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		userID, ok := claims["id"].(float64)
// 		if !ok {
// 			return false, fmt.Errorf("invalid token")
// 		}

// 		// Armazena o ID do usuário na variável "userID" do contexto
// 		fmt.Printf("user ID: %v\n", userID)
// 		return true, nil
// 	}

// 	return false, fmt.Errorf("invalid token")
// }
