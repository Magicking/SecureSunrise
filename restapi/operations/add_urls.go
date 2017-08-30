// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddUrlsHandlerFunc turns a function with the right signature into a add urls handler
type AddUrlsHandlerFunc func(AddUrlsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddUrlsHandlerFunc) Handle(params AddUrlsParams) middleware.Responder {
	return fn(params)
}

// AddUrlsHandler interface for that can handle valid add urls params
type AddUrlsHandler interface {
	Handle(AddUrlsParams) middleware.Responder
}

// NewAddUrls creates a new http.Handler for the add urls operation
func NewAddUrls(ctx *middleware.Context, handler AddUrlsHandler) *AddUrls {
	return &AddUrls{Context: ctx, Handler: handler}
}

/*AddUrls swagger:route POST /control/add addUrls

addURLs

Add a video URL to feed


*/
type AddUrls struct {
	Context *middleware.Context
	Handler AddUrlsHandler
}

func (o *AddUrls) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddUrlsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
