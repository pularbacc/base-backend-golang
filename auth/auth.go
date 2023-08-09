package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"pular.server/crypto"
	"pular.server/database"
	"pular.server/route/process"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auth struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type Identity struct {
	Account_id int64 `json:"account_id"`
}

func Login(db *sql.DB, loginReq LoginReq) (Auth, error) {
	var auth Auth

	passwordHash, err := database.GetPassByEmail(db, loginReq.Email)
	if err != nil {
		return auth, err
	}

	checkPass := crypto.CheckHash(loginReq.Password, passwordHash)
	if !checkPass {
		return auth, errors.New("incorect password")
	}

	account, err := database.GetAccountInfoByEmail(db, loginReq.Email)
	if err != nil {
		return auth, err
	}

	idToken, err := database.CreateToken(db, account.Id)
	if err != nil {
		return auth, err
	}

	token, err := crypto.GenerateToken(jwt.StandardClaims{
		Id:        strconv.FormatInt(idToken, 10),
		ExpiresAt: time.Now().Add(time.Minute * 6000).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   "TOKEN",
	})
	if err != nil {
		return auth, err
	}

	idRefreshToken, err := database.CreateRefreshToken(db, database.Token{
		Token_id:   idToken,
		Account_id: account.Id,
	})
	if err != nil {
		return auth, err
	}

	refreshToken, err := crypto.GenerateToken(jwt.StandardClaims{
		Id:        strconv.FormatInt(idRefreshToken, 10),
		ExpiresAt: time.Now().Add(time.Minute * 6000).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   "REFRESH_TOKEN",
	})
	if err != nil {
		return auth, err
	}

	auth = Auth{
		Access_token:  token,
		Refresh_token: refreshToken,
	}

	return auth, nil
}

func IdentityToken(db *sql.DB, token string) (Identity, error) {
	var identity Identity

	claims, err := crypto.ParseToken(token)
	if err != nil {
		return identity, err
	}

	idToken, err := strconv.ParseInt(claims.Id, 10, 64)
	if err != nil {
		return identity, err
	}

	idAccount, err := database.GetAccountByToken(db, idToken)
	if err != nil {
		return identity, err
	}

	identity = Identity{
		Account_id: idAccount,
	}

	return identity, nil
}

// middleware validation for request

func ValidationRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("on validation auth")

		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			process.WriteError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			process.WriteError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		db := r.Context().Value("db").(*sql.DB)

		indentity, err := IdentityToken(db, token)

		if err != nil {
			process.WriteError(w, http.StatusUnauthorized, "Token not valid")
			return
		}

		// ctx := context.WithValue(r.Context(), "identity", &indentity)
		// ctx = context.WithValue(r.Context(), "token", token)

		// fn(w, r.WithContext(ctx))

		ctx := context.WithValue(r.Context(), "identity", &indentity)
		ctx = context.WithValue(ctx, "Authorization", authHeader)

		fn(w, r.WithContext(ctx))
	}
}
