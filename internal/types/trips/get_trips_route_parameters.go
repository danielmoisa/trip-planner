// Code generated by go-swagger; DO NOT EDIT.

package trips

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetTripsRouteParams creates a new GetTripsRouteParams object
// with the default values initialized.
func NewGetTripsRouteParams() GetTripsRouteParams {

	var (
		// initialize parameters with default values

		pageDefault    = int64(1)
		perPageDefault = int64(25)
	)

	return GetTripsRouteParams{
		Page: &pageDefault,

		PerPage: &perPageDefault,
	}
}

// GetTripsRouteParams contains all the bound params for the get trips route operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTripsRoute
type GetTripsRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*page
	  Minimum: 1
	  In: query
	  Default: 1
	*/
	Page *int64 `query:"page"`
	/*per_page
	  Maximum: 80
	  Minimum: 1
	  In: query
	  Default: 25
	*/
	PerPage *int64 `query:"per_page"`
	/*Returns all trips matching the supplied name, if you need to add spaces just add an underscore (_).
	  In: query
	*/
	TripName *string `query:"trip_name"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTripsRouteParams() beforehand.
func (o *GetTripsRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPerPage, qhkPerPage, _ := qs.GetOK("per_page")
	if err := o.bindPerPage(qPerPage, qhkPerPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qTripName, qhkTripName, _ := qs.GetOK("trip_name")
	if err := o.bindTripName(qTripName, qhkTripName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTripsRouteParams) Validate(formats strfmt.Registry) error {
	var res []error

	// page
	// Required: false
	// AllowEmptyValue: false

	if err := o.validatePage(formats); err != nil {
		res = append(res, err)
	}

	// per_page
	// Required: false
	// AllowEmptyValue: false

	if err := o.validatePerPage(formats); err != nil {
		res = append(res, err)
	}

	// trip_name
	// Required: false
	// AllowEmptyValue: false

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetTripsRouteParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetTripsRouteParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	if err := o.validatePage(formats); err != nil {
		return err
	}

	return nil
}

// validatePage carries on validations for parameter Page
func (o *GetTripsRouteParams) validatePage(formats strfmt.Registry) error {

	// Required: false
	if o.Page == nil {
		return nil
	}

	if err := validate.MinimumInt("page", "query", *o.Page, 1, false); err != nil {
		return err
	}

	return nil
}

// bindPerPage binds and validates parameter PerPage from query.
func (o *GetTripsRouteParams) bindPerPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetTripsRouteParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("per_page", "query", "int64", raw)
	}
	o.PerPage = &value

	if err := o.validatePerPage(formats); err != nil {
		return err
	}

	return nil
}

// validatePerPage carries on validations for parameter PerPage
func (o *GetTripsRouteParams) validatePerPage(formats strfmt.Registry) error {

	// Required: false
	if o.PerPage == nil {
		return nil
	}

	if err := validate.MinimumInt("per_page", "query", *o.PerPage, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("per_page", "query", *o.PerPage, 80, false); err != nil {
		return err
	}

	return nil
}

// bindTripName binds and validates parameter TripName from query.
func (o *GetTripsRouteParams) bindTripName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TripName = &raw

	return nil
}
