// Package handlers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// User login
	// (POST /api/auth/login)
	PostApiAuthLogin(c *gin.Context)
	// User logout
	// (POST /api/auth/logout)
	PostApiAuthLogout(c *gin.Context)
	// Get user profile
	// (GET /api/auth/profile)
	GetApiAuthProfile(c *gin.Context)
	// Update user profile
	// (PUT /api/auth/profile)
	PutApiAuthProfile(c *gin.Context)
	// Refresh authentication token
	// (POST /api/auth/refresh)
	PostApiAuthRefresh(c *gin.Context)
	// Register a new user
	// (POST /api/auth/register)
	PostApiAuthRegister(c *gin.Context)
	// Initiate password reset
	// (POST /api/auth/reset)
	PostApiAuthReset(c *gin.Context)
	// Reset password
	// (PUT /api/auth/reset/{token})
	PutApiAuthResetToken(c *gin.Context, token string)
	// Verify user
	// (POST /api/auth/verify)
	PostApiAuthVerify(c *gin.Context)
	// Verify email
	// (GET /api/auth/verify/{token})
	GetApiAuthVerifyToken(c *gin.Context, token string)
	// Get all bills
	// (GET /api/bills)
	GetApiBills(c *gin.Context)
	// Create a new bill
	// (POST /api/bills)
	PostApiBills(c *gin.Context)
	// Delete a specific bill
	// (DELETE /api/bills/{id})
	DeleteApiBillsId(c *gin.Context, id string)
	// Get a specific bill
	// (GET /api/bills/{id})
	GetApiBillsId(c *gin.Context, id string)
	// Update a specific bill
	// (PUT /api/bills/{id})
	PutApiBillsId(c *gin.Context, id string)
	// Pay a specific bill
	// (POST /api/bills/{id}/pay)
	PostApiBillsIdPay(c *gin.Context, id string)
	// Get all feedback
	// (GET /api/feedback)
	GetApiFeedback(c *gin.Context)
	// Rate a merchant
	// (POST /api/feedback)
	PostApiFeedback(c *gin.Context)
	// Delete a feedback
	// (DELETE /api/feedback/{id})
	DeleteApiFeedbackId(c *gin.Context, id string)
	// Get a single feedback
	// (GET /api/feedback/{id})
	GetApiFeedbackId(c *gin.Context, id string)
	// Get all goals
	// (GET /api/goals)
	GetApiGoals(c *gin.Context)
	// Create a new goal
	// (POST /api/goals)
	PostApiGoals(c *gin.Context)
	// Delete a specific goal
	// (DELETE /api/goals/{id})
	DeleteApiGoalsId(c *gin.Context, id string)
	// Get a specific goal
	// (GET /api/goals/{id})
	GetApiGoalsId(c *gin.Context, id string)
	// Update a specific goal
	// (PUT /api/goals/{id})
	PutApiGoalsId(c *gin.Context, id string)
	// Get all integrations
	// (GET /api/integrations)
	GetApiIntegrations(c *gin.Context)
	// Connect with an accounting platform
	// (POST /api/integrations)
	PostApiIntegrations(c *gin.Context)
	// Disconnect from an accounting platform
	// (DELETE /api/integrations/{id})
	DeleteApiIntegrationsId(c *gin.Context, id string)
	// Get a specific integration
	// (GET /api/integrations/{id})
	GetApiIntegrationsId(c *gin.Context, id string)
	// Get all notifications
	// (GET /api/notifications)
	GetApiNotifications(c *gin.Context)
	// Send a new notification
	// (POST /api/notifications)
	PostApiNotifications(c *gin.Context)
	// Delete a specific notification
	// (DELETE /api/notifications/{id})
	DeleteApiNotificationsId(c *gin.Context, id string)
	// Get a specific notification
	// (GET /api/notifications/{id})
	GetApiNotificationsId(c *gin.Context, id string)
	// Get all reports
	// (GET /api/reports)
	GetApiReports(c *gin.Context)
	// Generate a new report
	// (POST /api/reports)
	PostApiReports(c *gin.Context)
	// Delete a specific report
	// (DELETE /api/reports/{id})
	DeleteApiReportsId(c *gin.Context, id string)
	// Get a specific report
	// (GET /api/reports/{id})
	GetApiReportsId(c *gin.Context, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostApiAuthLogin operation middleware
func (siw *ServerInterfaceWrapper) PostApiAuthLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiAuthLogin(c)
}

// PostApiAuthLogout operation middleware
func (siw *ServerInterfaceWrapper) PostApiAuthLogout(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiAuthLogout(c)
}

// GetApiAuthProfile operation middleware
func (siw *ServerInterfaceWrapper) GetApiAuthProfile(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiAuthProfile(c)
}

// PutApiAuthProfile operation middleware
func (siw *ServerInterfaceWrapper) PutApiAuthProfile(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutApiAuthProfile(c)
}

// PostApiAuthRefresh operation middleware
func (siw *ServerInterfaceWrapper) PostApiAuthRefresh(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiAuthRefresh(c)
}

// PostApiAuthRegister operation middleware
func (siw *ServerInterfaceWrapper) PostApiAuthRegister(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiAuthRegister(c)
}

// PostApiAuthReset operation middleware
func (siw *ServerInterfaceWrapper) PostApiAuthReset(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiAuthReset(c)
}

// PutApiAuthResetToken operation middleware
func (siw *ServerInterfaceWrapper) PutApiAuthResetToken(c *gin.Context) {

	var err error

	// ------------- Path parameter "token" -------------
	var token string

	err = runtime.BindStyledParameterWithOptions("simple", "token", c.Param("token"), &token, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter token: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutApiAuthResetToken(c, token)
}

// PostApiAuthVerify operation middleware
func (siw *ServerInterfaceWrapper) PostApiAuthVerify(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiAuthVerify(c)
}

// GetApiAuthVerifyToken operation middleware
func (siw *ServerInterfaceWrapper) GetApiAuthVerifyToken(c *gin.Context) {

	var err error

	// ------------- Path parameter "token" -------------
	var token string

	err = runtime.BindStyledParameterWithOptions("simple", "token", c.Param("token"), &token, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter token: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiAuthVerifyToken(c, token)
}

// GetApiBills operation middleware
func (siw *ServerInterfaceWrapper) GetApiBills(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiBills(c)
}

// PostApiBills operation middleware
func (siw *ServerInterfaceWrapper) PostApiBills(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiBills(c)
}

// DeleteApiBillsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiBillsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiBillsId(c, id)
}

// GetApiBillsId operation middleware
func (siw *ServerInterfaceWrapper) GetApiBillsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiBillsId(c, id)
}

// PutApiBillsId operation middleware
func (siw *ServerInterfaceWrapper) PutApiBillsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutApiBillsId(c, id)
}

// PostApiBillsIdPay operation middleware
func (siw *ServerInterfaceWrapper) PostApiBillsIdPay(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiBillsIdPay(c, id)
}

// GetApiFeedback operation middleware
func (siw *ServerInterfaceWrapper) GetApiFeedback(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiFeedback(c)
}

// PostApiFeedback operation middleware
func (siw *ServerInterfaceWrapper) PostApiFeedback(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiFeedback(c)
}

// DeleteApiFeedbackId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiFeedbackId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiFeedbackId(c, id)
}

// GetApiFeedbackId operation middleware
func (siw *ServerInterfaceWrapper) GetApiFeedbackId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiFeedbackId(c, id)
}

// GetApiGoals operation middleware
func (siw *ServerInterfaceWrapper) GetApiGoals(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiGoals(c)
}

// PostApiGoals operation middleware
func (siw *ServerInterfaceWrapper) PostApiGoals(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiGoals(c)
}

// DeleteApiGoalsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiGoalsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiGoalsId(c, id)
}

// GetApiGoalsId operation middleware
func (siw *ServerInterfaceWrapper) GetApiGoalsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiGoalsId(c, id)
}

// PutApiGoalsId operation middleware
func (siw *ServerInterfaceWrapper) PutApiGoalsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutApiGoalsId(c, id)
}

// GetApiIntegrations operation middleware
func (siw *ServerInterfaceWrapper) GetApiIntegrations(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiIntegrations(c)
}

// PostApiIntegrations operation middleware
func (siw *ServerInterfaceWrapper) PostApiIntegrations(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiIntegrations(c)
}

// DeleteApiIntegrationsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiIntegrationsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiIntegrationsId(c, id)
}

// GetApiIntegrationsId operation middleware
func (siw *ServerInterfaceWrapper) GetApiIntegrationsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiIntegrationsId(c, id)
}

// GetApiNotifications operation middleware
func (siw *ServerInterfaceWrapper) GetApiNotifications(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiNotifications(c)
}

// PostApiNotifications operation middleware
func (siw *ServerInterfaceWrapper) PostApiNotifications(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiNotifications(c)
}

// DeleteApiNotificationsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiNotificationsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiNotificationsId(c, id)
}

// GetApiNotificationsId operation middleware
func (siw *ServerInterfaceWrapper) GetApiNotificationsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiNotificationsId(c, id)
}

// GetApiReports operation middleware
func (siw *ServerInterfaceWrapper) GetApiReports(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiReports(c)
}

// PostApiReports operation middleware
func (siw *ServerInterfaceWrapper) PostApiReports(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiReports(c)
}

// DeleteApiReportsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiReportsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiReportsId(c, id)
}

// GetApiReportsId operation middleware
func (siw *ServerInterfaceWrapper) GetApiReportsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiReportsId(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/api/auth/login", wrapper.PostApiAuthLogin)
	router.POST(options.BaseURL+"/api/auth/logout", wrapper.PostApiAuthLogout)
	router.GET(options.BaseURL+"/api/auth/profile", wrapper.GetApiAuthProfile)
	router.PUT(options.BaseURL+"/api/auth/profile", wrapper.PutApiAuthProfile)
	router.POST(options.BaseURL+"/api/auth/refresh", wrapper.PostApiAuthRefresh)
	router.POST(options.BaseURL+"/api/auth/register", wrapper.PostApiAuthRegister)
	router.POST(options.BaseURL+"/api/auth/reset", wrapper.PostApiAuthReset)
	router.PUT(options.BaseURL+"/api/auth/reset/:token", wrapper.PutApiAuthResetToken)
	router.POST(options.BaseURL+"/api/auth/verify", wrapper.PostApiAuthVerify)
	router.GET(options.BaseURL+"/api/auth/verify/:token", wrapper.GetApiAuthVerifyToken)
	router.GET(options.BaseURL+"/api/bills", wrapper.GetApiBills)
	router.POST(options.BaseURL+"/api/bills", wrapper.PostApiBills)
	router.DELETE(options.BaseURL+"/api/bills/:id", wrapper.DeleteApiBillsId)
	router.GET(options.BaseURL+"/api/bills/:id", wrapper.GetApiBillsId)
	router.PUT(options.BaseURL+"/api/bills/:id", wrapper.PutApiBillsId)
	router.POST(options.BaseURL+"/api/bills/:id/pay", wrapper.PostApiBillsIdPay)
	router.GET(options.BaseURL+"/api/feedback", wrapper.GetApiFeedback)
	router.POST(options.BaseURL+"/api/feedback", wrapper.PostApiFeedback)
	router.DELETE(options.BaseURL+"/api/feedback/:id", wrapper.DeleteApiFeedbackId)
	router.GET(options.BaseURL+"/api/feedback/:id", wrapper.GetApiFeedbackId)
	router.GET(options.BaseURL+"/api/goals", wrapper.GetApiGoals)
	router.POST(options.BaseURL+"/api/goals", wrapper.PostApiGoals)
	router.DELETE(options.BaseURL+"/api/goals/:id", wrapper.DeleteApiGoalsId)
	router.GET(options.BaseURL+"/api/goals/:id", wrapper.GetApiGoalsId)
	router.PUT(options.BaseURL+"/api/goals/:id", wrapper.PutApiGoalsId)
	router.GET(options.BaseURL+"/api/integrations", wrapper.GetApiIntegrations)
	router.POST(options.BaseURL+"/api/integrations", wrapper.PostApiIntegrations)
	router.DELETE(options.BaseURL+"/api/integrations/:id", wrapper.DeleteApiIntegrationsId)
	router.GET(options.BaseURL+"/api/integrations/:id", wrapper.GetApiIntegrationsId)
	router.GET(options.BaseURL+"/api/notifications", wrapper.GetApiNotifications)
	router.POST(options.BaseURL+"/api/notifications", wrapper.PostApiNotifications)
	router.DELETE(options.BaseURL+"/api/notifications/:id", wrapper.DeleteApiNotificationsId)
	router.GET(options.BaseURL+"/api/notifications/:id", wrapper.GetApiNotificationsId)
	router.GET(options.BaseURL+"/api/reports", wrapper.GetApiReports)
	router.POST(options.BaseURL+"/api/reports", wrapper.PostApiReports)
	router.DELETE(options.BaseURL+"/api/reports/:id", wrapper.DeleteApiReportsId)
	router.GET(options.BaseURL+"/api/reports/:id", wrapper.GetApiReportsId)
}