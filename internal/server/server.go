package server

import (
	"C0lliNN/auth-server/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	auth auth.Auth
}

func NewServer(auth auth.Auth) Server {
	return Server{auth: auth}
}

func (s Server) Start() error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	r.POST("/token", func(context *gin.Context) {
		var req auth.ObtainTokenRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.Error(err)
			return
		}

		response, err := s.auth.ObtainToken(context, req)
		if err != nil {
			context.Error(err)
			return
		}

		context.JSON(http.StatusOK, response)
	})

	server := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
	}

	return server.ListenAndServe()
}
