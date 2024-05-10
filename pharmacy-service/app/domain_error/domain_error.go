package domain_error

import (
	"encoding/json"
	"fmt"
	dc "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error/domain_code"
	"github.com/gofiber/fiber/v2/log"
)

var (
	OK                        = NewDomainError(dc.OK, "OK")
	ErrCheckIfUserExists      = NewDomainError(dc.Internal, "error checking if user exists")
	ErrUserAlreadyExists      = NewDomainError(dc.BadRequest, "user already exists")
	ErrSignUp                 = NewDomainError(dc.Internal, "error signing up")
	ErrCreateUser             = NewDomainError(dc.Internal, "error creating user")
	ErrUserNotFound           = NewDomainError(dc.NotFound, "user not found")
	ErrWrongPassword          = NewDomainError(dc.BadRequest, "wrong password")
	ErrSignIn                 = NewDomainError(dc.Internal, "error signing in")
	ErrGetPassword            = NewDomainError(dc.Internal, "error getting password")
	ErrGetUserID              = NewDomainError(dc.Internal, "error getting user id")
	ErrGetUser                = NewDomainError(dc.Internal, "error getting user")
	ErrFetchMedicinalProducts = NewDomainError(dc.Internal, "error fetching medicinal products")
	ErrGetSession             = NewDomainError(dc.Internal, "error getting session")
	ErrUnauthorized           = NewDomainError(dc.Unauthorized, "unauthorized")
	ErrSaveSession            = NewDomainError(dc.Internal, "error saving session")
	ErrCreateSession          = NewDomainError(dc.Internal, "error creating session")
	ErrParseRequestBody       = NewDomainError(dc.BadRequest, "error parsing request body")
	ErrRequestBodyInvalid     = NewDomainError(dc.BadRequest, "request body validation failed")
	ErrInvalidUserID          = NewDomainError(dc.BadRequest, "invalid user id")
	ErrForbidden              = NewDomainError(dc.Forbidden, "forbidden")
	ErrDeleteSession          = NewDomainError(dc.Internal, "error deleting session")
)

type DomainError struct {
	code    dc.DomainCode
	message string
	params  map[string]any
}

func (de *DomainError) Error() string {
	if de.params != nil {
		jsonParams, err := json.Marshal(de.params)
		if err != nil {
			log.Warnf("error processing params: %s", err.Error())
			return de.message
		}

		return fmt.Sprintf("%s: %s", jsonParams, de.message)
	}

	return de.message
}

func (de *DomainError) Code() dc.DomainCode {
	return de.code
}

func (de *DomainError) Message() string {
	return de.message
}

func (de *DomainError) Params() map[string]any {
	return de.params
}

func NewDomainError(code dc.DomainCode, message string) *DomainError {
	return &DomainError{
		code:    code,
		message: message,
	}
}

func (de *DomainError) WithParams(params ...any) *DomainError {
	if params == nil {
		return de
	}

	if len(params)%2 != 0 {
		log.Warnf("params length = %d: params length should be even number", len(params))
		return de
	}

	paramsMap := make(map[string]any)

	for k := 0; k < len(params)-1; k += 2 {
		paramName, ok := params[k].(string)

		if !ok {
			log.Warnf("paramName = %+v: param name should be of type string", params[k])
			return de
		}

		paramsMap[paramName] = params[k+1]
	}

	de.params = paramsMap
	return de
}
