package auth

import (
	"customerCrud/model/general"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func GenerateToken(userID uuid.UUID) (string, error) {
	var viperI = viper.New()
	viperI.AddConfigPath(".")
	viperI.SetConfigName("config") // name of config file (without extension)
	viperI.SetConfigType("env")
	if err := viperI.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return "", err
		} else {
			// Config file was found but another error was produced
			fmt.Println("token", err)
			return "", err
		}
	}

	// Config file found and successfully parsed

	secretKey := viperI.GetString("MY_JWT_TOKEN")

	claims := &general.Claim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println(secretKey)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("signed", err)
		return "", err
	}

	return tokenString, nil
}

func CreateToken(userID uuid.UUID, tokenDetails *general.TokenDetails) error {

	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.TokenUUID = uuid.NewV4()

	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUUID = userID

	err := createAccessToken(userID, tokenDetails)
	if err != nil {
		return err
	}

	err = createRefreshToken(userID, tokenDetails)
	if err != nil {
		return err
	}

	return nil
}

func createAccessToken(userID uuid.UUID, tokenDetails *general.TokenDetails) error {

	var err error
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["access_uuid"] = tokenDetails.TokenUUID
	accessTokenClaims["user_id"] = userID
	accessTokenClaims["exp"] = tokenDetails.AtExpires

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	tokenDetails.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return err
	}
	return nil
}

//Creating Refresh Token
func createRefreshToken(userID uuid.UUID, tokenDetails *general.TokenDetails) error {

	var err error
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["refresh_uuid"] = tokenDetails.RefreshUUID
	refreshTokenClaims["user_id"] = userID
	refreshTokenClaims["exp"] = tokenDetails.RtExpires

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return err
	}

	return nil
}
