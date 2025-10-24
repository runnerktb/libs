package token

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/runnerktb/libs/universe"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	lib "github.com/runnerktb/libs/helper"
)

func ClaimToken(tokens []string) (response AuthorizationInfo, serr serror.SError) {
	//secretKey := []byte("um_super_apps")
	publicKeyFile := "public_key.pem"
	publicKeyBytes, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		secretKey := []byte("um_super_apps")

		if tokens == nil {
			return response, serror.New("Token tidak ditemukan")
		}

		tokenString, err := parseToken(tokens[0])
		if err != nil {
			return response, serror.NewFromError(err)
		}
		decode, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, serror.NewFromError(fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]))
			}
			return []byte(secretKey), nil
		})
		resToken := decode.Claims.(jwt.MapClaims)
		if err != nil {
			return response, serror.NewFromError(err)
		}

		tmpUserAccess := []string{}
		for _, v := range resToken["user_access"].([]interface{}) {
			tmpUserAccess = append(tmpUserAccess, v.(string))
		}
		isOrgAdmin := int(helper.StringToInt(helper.IntToString(int(resToken["isorgadmin"].(float64))), 0))
		isActive := int(helper.StringToInt(helper.IntToString(int(resToken["isactive"].(float64))), 0))
		response = AuthorizationInfo{
			UserID:         fmt.Sprintf("%v", resToken["user_id"]),
			Username:       fmt.Sprintf("%v", resToken["username"]),
			IsOrgAdmin:     &isOrgAdmin,
			IsActive:       isActive,
			OrganizationId: fmt.Sprintf("%v", resToken["organization_id"]),
			App:            fmt.Sprintf("%v", resToken["app"]),
			Email:          fmt.Sprintf("%v", resToken["email"]),
			Exp:            int(helper.StringToInt(helper.IntToString(int(resToken["exp"].(float64))), 0)),
			Name:           fmt.Sprintf("%v", resToken["name"]),
			RealmID:        fmt.Sprintf("%v", resToken["realm_id"]),
			UserAccess:     tmpUserAccess,
			Jti:            fmt.Sprintf("%v", resToken["jti"]),
			Sid:            fmt.Sprintf("%v", resToken["sid"]),
		}
		return response, nil
	} else {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
		if err != nil {
			return response, serror.NewFromError(err)
		}

		if tokens == nil {
			return response, serror.New("Token tidak ditemukan")
		}

		tokenString, err := parseToken(tokens[0])
		if err != nil {
			return response, serror.NewFromError(err)
		}
		decode, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
		resToken := decode.Claims.(jwt.MapClaims)
		if err != nil {
			return response, serror.NewFromError(err)
		}

		tmpUserAccess := []string{}
		for _, v := range resToken["user_access"].([]interface{}) {
			tmpUserAccess = append(tmpUserAccess, v.(string))
		}
		isOrgAdmin := int(helper.StringToInt(helper.IntToString(int(resToken["isorgadmin"].(float64))), 0))
		isActive := int(helper.StringToInt(helper.IntToString(int(resToken["isactive"].(float64))), 0))
		response = AuthorizationInfo{
			UserID:         fmt.Sprintf("%v", resToken["user_id"]),
			Username:       fmt.Sprintf("%v", resToken["username"]),
			IsOrgAdmin:     &isOrgAdmin,
			IsActive:       isActive,
			OrganizationId: fmt.Sprintf("%v", resToken["organization_id"]),
			App:            fmt.Sprintf("%v", resToken["app"]),
			Email:          fmt.Sprintf("%v", resToken["email"]),
			Exp:            int(helper.StringToInt(helper.IntToString(int(resToken["exp"].(float64))), 0)),
			Name:           fmt.Sprintf("%v", resToken["name"]),
			RealmID:        fmt.Sprintf("%v", resToken["realm_id"]),
			UserAccess:     tmpUserAccess,
			Jti:            fmt.Sprintf("%v", resToken["jti"]),
			Sid:            fmt.Sprintf("%v", resToken["sid"]),
		}
		return response, nil
	}
}

func parseToken(source string) (token string, err error) {

	separator := " "
	valueSection := 1
	expectedTokenLength := 2

	if source == "" {
		return token, errors.New("Token tidak ditemukan")
	}

	tokens := strings.Split(source, separator)
	if len(tokens) != expectedTokenLength {
		return token, errors.New("Token tidak valid")
	}

	token = tokens[valueSection]
	return token, nil
}

func HasAccess(module string, userAccess []string) (isHas bool, serr serror.SError) {
	acc := universe.ParseModule(module)
	if acc.ID == "" {
		serr = serror.New(fmt.Sprintf("module %s not found", module))
		return false, serr
	}

	if lib.CheckArrayString(acc.ID, userAccess) {
		return true, nil
	}

	serr = serror.New("Token module denied")
	return false, serr
}
