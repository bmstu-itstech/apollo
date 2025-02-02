// Package httpport provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package httpport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /departaments)
	GetDepartaments(w http.ResponseWriter, r *http.Request)

	// (GET /departaments/{id})
	GetDepartament(w http.ResponseWriter, r *http.Request, id int)

	// (GET /disciplines)
	GetDisciplines(w http.ResponseWriter, r *http.Request)

	// (GET /disciplines/{id})
	GetDiscipline(w http.ResponseWriter, r *http.Request, id int)

	// (GET /materials)
	GetMaterials(w http.ResponseWriter, r *http.Request, params GetMaterialsParams)

	// (GET /materials/{uuid})
	GetMaterial(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID)

	// (PUT /materials/{uuid})
	PutMaterial(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /departaments)
func (_ Unimplemented) GetDepartaments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /departaments/{id})
func (_ Unimplemented) GetDepartament(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /disciplines)
func (_ Unimplemented) GetDisciplines(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /disciplines/{id})
func (_ Unimplemented) GetDiscipline(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /materials)
func (_ Unimplemented) GetMaterials(w http.ResponseWriter, r *http.Request, params GetMaterialsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /materials/{uuid})
func (_ Unimplemented) GetMaterial(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PUT /materials/{uuid})
func (_ Unimplemented) PutMaterial(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetDepartaments operation middleware
func (siw *ServerInterfaceWrapper) GetDepartaments(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDepartaments(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetDepartament operation middleware
func (siw *ServerInterfaceWrapper) GetDepartament(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDepartament(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetDisciplines operation middleware
func (siw *ServerInterfaceWrapper) GetDisciplines(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDisciplines(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetDiscipline operation middleware
func (siw *ServerInterfaceWrapper) GetDiscipline(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDiscipline(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMaterials operation middleware
func (siw *ServerInterfaceWrapper) GetMaterials(w http.ResponseWriter, r *http.Request) {

	var err error

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMaterialsParams

	// ------------- Optional query parameter "discipline_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "discipline_id", r.URL.Query(), &params.DisciplineId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "discipline_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMaterials(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMaterial operation middleware
func (siw *ServerInterfaceWrapper) GetMaterial(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", chi.URLParam(r, "uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "uuid", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMaterial(w, r, uuid)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PutMaterial operation middleware
func (siw *ServerInterfaceWrapper) PutMaterial(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", chi.URLParam(r, "uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "uuid", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutMaterial(w, r, uuid)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/departaments", wrapper.GetDepartaments)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/departaments/{id}", wrapper.GetDepartament)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/disciplines", wrapper.GetDisciplines)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/disciplines/{id}", wrapper.GetDiscipline)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/materials", wrapper.GetMaterials)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/materials/{uuid}", wrapper.GetMaterial)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/materials/{uuid}", wrapper.PutMaterial)
	})

	return r
}
