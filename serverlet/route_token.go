package main

import (
	"encoding/json"
	"fmt"
	"github.com/PichuChen/go-bbs"
	"github.com/PichuChen/go-bbs/crypt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

func routeToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postToken(w, r)
		return
	}

}

func postToken(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	userec, err := findUserecById(username)
	if err != nil {
		m := map[string]string{
			"error":             "grant_error",
			"error_description": err.Error(),
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Write(b)
		return

	}

	log.Println("found user:", *userec)
	err = verifyPassword(userec, password)
	if err != nil {

		// TODO: add delay, warning, notify user

		m := map[string]string{
			"error":             "grant_error",
			"error_description": err.Error(),
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Write(b)
		return
	}

	// Generate Access Token
	token := newAccessTokenWithUsername(username)
	m := map[string]string{
		"access_token": token,
		"token_type":   "bearer",
	}

	b, _ := json.MarshalIndent(m, "", "  ")

	w.Write(b)

}

func findUserecById(userid string) (*bbs.Userec, error) {

	for _, it := range userRecs {
		if userid == it.UserId {
			return it, nil
		}
	}
	return nil, fmt.Errorf("user record not found")

}

func verifyPassword(userec *bbs.Userec, password string) error {
	log.Println("password", userec.Password)
	res, err := crypt.Fcrypt([]byte(password), []byte(userec.Password[:2]))
	str := strings.Trim(string(res), "\x00")
	log.Println("res", str, err, []byte(str), []byte(userec.Password))

	if str != userec.Password {
		return fmt.Errorf("password incorrect")
	}
	return nil

}

func getTokenFromRequest(r *http.Request) string {
	a := r.Header.Get("Authorization")
	s := strings.Split(a, " ")
	if len(s) < 2 {
		log.Println("getTokenFromRequest error: len(s) < 2, got", len(s))
		return ""
	}
	return s[1]
}

func newAccessTokenWithUsername(username string) string {
	claims := &jwt.StandardClaims{
		ExpiresAt: 30 * 60, // TODO: Setting me in config
		// Issuer:    "test",
		Subject: username,
	}

	// TODO: Setting me in config
	// openssl ecparam -name prime256v1 -genkey -noout -out pkey
	privateKey := `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIABVEwM0EuOpmOoe803/vYswLUtsaR71xfGzk06TGBy/oAoGCCqGSM49
AwEHoUQDQgAEV8qJS5x98i0eM+UUiV5qB2JZhT67Ojl6/rZ4xKcHM/NLpUJP3wDp
eFQfMaMiAKQHoGs6rk5z1l/tUUVjJWrw0A==
-----END EC PRIVATE KEY-----`

	// openssl ec -in pkey -pubout
	// 	publicKey := `-----BEGIN PUBLIC KEY-----
	// MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV8qJS5x98i0eM+UUiV5qB2JZhT67
	// Ojl6/rZ4xKcHM/NLpUJP3wDpeFQfMaMiAKQHoGs6rk5z1l/tUUVjJWrw0A==
	// -----END PUBLIC KEY-----`

	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		log.Println("parse private key failed:", err)
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Println("sign token failed:", err)
		return ""
	}
	return ss
}
