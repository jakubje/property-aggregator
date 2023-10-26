package api

import (
	"context"
	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (server *Server) getProperties(ctx *gin.Context) {
	cursor, err := server.mongoCollection.Find(ctx, bson.D{{}})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal().Msg(err.Error())
	}

	ctx.JSON(http.StatusOK, results)
}
