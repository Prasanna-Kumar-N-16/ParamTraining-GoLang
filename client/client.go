package main

import (
	"context"
	"fmt"
	"log"
	"rpc/proto"
	"strconv"
	"time"

	"github.com/c-bata/go-prompt"
	"google.golang.org/grpc"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Add", Description: "Addition is done"},
		{Text: "Subtract", Description: "Subtraction is done"},
		{Text: "Multiply", Description: "Multiplication is done"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordAfterCursor(), true)
}
func ncompleter(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(nil, d.GetWordAfterCursor(), true)
}
func main() {
	t := prompt.Input(">>>", completer)
	m := prompt.Input(">", ncompleter)
	a, err := strconv.ParseInt(string(m[0]), 10, 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseInt(string(m[2]), 10, 64)
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewAddServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if t == "Add" {
		r, err := client.Add(ctx, &proto.Request{A: a, B: b})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println(r.GetResult())
	}
	if t == "Subtract" {
		re, err := client.Subtract(ctx, &proto.Request{A: a, B: b})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println(re.GetResult())
	}
	if t == "Multiply" {
		rl, err := client.Multiply(ctx, &proto.Request{A: a, B: b})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println(rl.GetResult())
	}
	/*g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid a"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid b"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	g.GET("/sub/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid a"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid b"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Subtract(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid a"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid b"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}*/
}
