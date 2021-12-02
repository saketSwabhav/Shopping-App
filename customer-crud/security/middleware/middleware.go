package middleware

import (
	"customerCrud/model/general"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Token -> ", r.Header)

		var viperI = viper.New()
		viperI.AddConfigPath(".")
		viperI.SetConfigName("config") // name of config file (without extension)
		viperI.SetConfigType("env")
		if err := viperI.ReadInConfig(); err != nil {
			if err, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				fmt.Println(err)
				return
			} else {
				// Config file was found but another error was produced
				fmt.Println(err)
				return
			}
		}

		// Config file found and successfully parsed

		jwtKey := viperI.GetString("MY_JWT_TOKEN")
		// var jwtKey = os.Getenv("ACCESS_SECRET")
		// var jwtKey = fmt.Println("-------- jwtKey ->", jwtKey)

		if r.Header.Get("Token") != "" {

			claims := &general.Claim{}

			token, err := jwt.ParseWithClaims(r.Header.Get("Token"), claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte(jwtKey), nil
			})

			fmt.Println("-------- token.Valid ->", token.Valid)

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					http.Error(w, "User Not Authorized", http.StatusUnauthorized)
					return
				}
				http.Error(w, "User Not Authorized", http.StatusBadRequest)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {
			http.Error(w, "User Not Authorized", http.StatusUnauthorized)
		}
	})
}
