// Code generated by go-swagger; DO NOT EDIT.

package shopper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/mercadofarma/services/models"
)

// NewSignUpShopperParams creates a new SignUpShopperParams object
//
// There are no default values defined in the spec.
func NewSignUpShopperParams() SignUpShopperParams {

	return SignUpShopperParams{}
}

// SignUpShopperParams contains all the bound params for the sign up shopper operation
// typically these are obtained from a http.Request
//
// swagger:parameters signUpShopper
type SignUpShopperParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*SignUpShopperRequest
	  Required: true
	  In: body
	*/
	SignUpShopperRequest *models.SignUpShopperRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSignUpShopperParams() beforehand.
func (o *SignUpShopperParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SignUpShopperRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("signUpShopperRequest", "body", ""))
			} else {
				res = append(res, errors.NewParseError("signUpShopperRequest", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.SignUpShopperRequest = &body
			}
		}
	} else {
		res = append(res, errors.Required("signUpShopperRequest", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
