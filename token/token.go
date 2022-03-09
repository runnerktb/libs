package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"strings"
)

func ClaimToken(tokens []string) (response AuthorizationInfo, serr serror.SError) {
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

	if err != nil {
		return response, serror.NewFromError(err)
	}

	resToken := decode.Claims.(jwt.MapClaims)
	tmpUserAccess := []string{}
	for _, v := range resToken["user_access"].([]interface{}) {
		tmpUserAccess = append(tmpUserAccess, v.(string))
	}
	isOrgAdmin := int(helper.StringToInt(helper.IntToString(int(resToken["isorgadmin"].(float64))), 0))
	response = AuthorizationInfo{
		UserID:     fmt.Sprintf("%v", resToken["user_id"]),
		Username:   fmt.Sprintf("%v", resToken["username"]),
		IsOrgAdmin: &isOrgAdmin,
		//IsActive:       int(helper.StringToInt(helper.IntToString(int(resToken["isactive"].(float64))), 0)),
		OrganizationId: fmt.Sprintf("%v", resToken["organization_id"]),
		AppId:          fmt.Sprintf("%v", resToken["app"]),
		Exp:            int(helper.StringToInt(helper.IntToString(int(resToken["exp"].(float64))), 0)),
		UserAccess:     tmpUserAccess,
	}

	return response, nil
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
