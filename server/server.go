package server

import (
	"fmt"
	"net/http"

	"github.com/Kolynes/PotGo/context"
)

type Server struct {
	host         string
	port         int
	CertFile     string
	KeyFile      string
	baseInstance *http.Server
	context      *context.Context
}

func (server *Server) Initiate() {
	server.baseInstance = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.host, server.port),
		Handler: server,
	}
	print("Server running...")
	if server.CertFile != "" && server.KeyFile != "" {
		server.baseInstance.ListenAndServeTLS(server.CertFile, server.KeyFile)
	} else {
		server.baseInstance.ListenAndServe()
	}
}

func (server *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response := server.context.GetMiddlewareContext().Handle(request)
	writer.WriteHeader(response.Status)
	response.Headers.Write(writer)
	writer.Write(response.Body)
}

func NewServer(context *context.Context) *Server {
	return &Server{
		host:     (*context.Env)["HOST"].(string),
		port:     (*context.Env)["PORT"].(int),
		CertFile: (*context.Env)["CERT_FILE"].(string),
		KeyFile:  (*context.Env)["KEY_FILE"].(string),
		context:  context,
	}
}
