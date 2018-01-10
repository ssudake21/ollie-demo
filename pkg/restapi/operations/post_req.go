// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostReqHandlerFunc turns a function with the right signature into a post req handler
type PostReqHandlerFunc func(PostReqParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostReqHandlerFunc) Handle(params PostReqParams) middleware.Responder {
	return fn(params)
}

// PostReqHandler interface for that can handle valid post req params
type PostReqHandler interface {
	Handle(PostReqParams) middleware.Responder
}

// NewPostReq creates a new http.Handler for the post req operation
func NewPostReq(ctx *middleware.Context, handler PostReqHandler) *PostReq {
	return &PostReq{Context: ctx, Handler: handler}
}

/*PostReq swagger:route POST / postReq

Alexa service request

*/
type PostReq struct {
	Context *middleware.Context
	Handler PostReqHandler
}

func (o *PostReq) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostReqParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
