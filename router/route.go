package router

import (
	"github.com/julienschmidt/httprouter"
	"zvon/handler"
)

type Route struct {
	Name, Path string

	// Route will set its own Content-Type header. If it's false, it  is automatically set to application/json
	CustomContentType bool

	// Handlers for different methods
	GET    httprouter.Handle
	POST   httprouter.Handle
	PUT    httprouter.Handle
	DELETE httprouter.Handle
}

var routes = []Route{
	{
		Name:              "issues_hook",
		Path:              "/webhook/issues",
		CustomContentType: true,
		POST:              handler.IssuesWebhook,
	},
	{
		Name: "banging",
		Path: "/bang",
		GET:  handler.BangerHandler,
	},
}
