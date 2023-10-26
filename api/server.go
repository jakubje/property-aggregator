package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jakubje/property-aggregator/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config          util.Config
	router          *gin.Engine
	ctx             gin.Context
	mongoCollection *mongo.Collection
}

func NewServer(config util.Config, mongoCollection *mongo.Collection) (*Server, error) {

	server := &Server{
		config:          config,
		mongoCollection: mongoCollection,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/properties", server.getProperties)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
