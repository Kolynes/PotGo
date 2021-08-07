package server

import (
	"fmt"
	"net/http"

	"github.com/Kolynes/PotGo/context"
)

type Server struct {
	Host     string
	Port     int
	CertFile string
	KeyFile  string
	server   *http.Server
	Context  *context.Context
}

type IServer interface {
	http.Handler
	Initiate()
}

func (server *Server) Initiate(context *context.Context) {
	server.Context = context
	server.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.Host, server.Port),
		Handler: server,
	}
	print("Server running...")
	if server.CertFile != "" && server.KeyFile != "" {
		server.server.ListenAndServeTLS(server.CertFile, server.KeyFile)
	} else {
		server.server.ListenAndServe()
	}
}

func (server *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response, err := server.Context.Middleware[0].Handle(request)
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		writer.WriteHeader(response.Status)
		response.Headers.Write(writer)
		writer.Write(response.Body)
	}
}
