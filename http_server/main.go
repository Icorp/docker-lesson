package main

import (
	"context"
	"github.com/Icorp/docker-lesson/calc_proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/cal/add/:a/:b", func(c *gin.Context) {

		a, err := strconv.ParseInt(c.Param("a"), 10, 64)
		if err != nil {
			_ = c.Error(err)
			return
		}

		b, err := strconv.ParseInt(c.Param("b"), 10, 64)
		if err != nil {
			panic(err)
			return
		}

		conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			_ = c.Error(err)
			return
		}
		defer conn.Close()

		protoConn := calc_proto.NewAddServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		r, err := protoConn.Add(ctx, &calc_proto.CalcRequest{
			A: a,
			B: b,
		})
		if err != nil {
			_ = c.Error(err)
		}

		c.JSON(200, gin.H{
			"result": r.GetResult(),
		})
	})
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
