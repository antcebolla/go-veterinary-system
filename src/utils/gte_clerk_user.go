package utils

import (
	"errors"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-gonic/gin"
)

func GetClerkInfo(ctx *gin.Context) (*clerk.User, *clerk.SessionClaims, error) {
	userNonTyped, ex := ctx.Get("user")
	if !ex {
		return nil, nil, errors.New("no user found in context, middleware failed to set it, something is wrong")
	}
	user, ok := userNonTyped.(*clerk.User)
	if !ok {
		return nil, nil, errors.New("user in context is not of type *clerk.User, something is wrong")
	}

	claimsNonTyped, ex := ctx.Get("claims")
	if !ex {
		return nil, nil, errors.New("no claims found in context, middleware failed to set them, something is wrong")
	}
	claims, ok := claimsNonTyped.(*clerk.SessionClaims)
	if !ok {
		return nil, nil, errors.New("claims in context is not of type *clerk.SessionClaims, something is wrong")
	}

	return user, claims, nil
}
