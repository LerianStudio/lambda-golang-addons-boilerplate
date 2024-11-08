package main

import (
	"context"
	"fmt"
	"lambda-golang-addons-boilerplate/api"
	"lambda-golang-addons-boilerplate/config"
	"lambda-golang-addons-boilerplate/repository"
	"lambda-golang-addons-boilerplate/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"go.uber.org/zap"
)

var app *fiber.App
var logger *zap.Logger

func init() {
	// Initialize logger
	logger, _ = zap.NewProduction()
	defer func() {
		err := logger.Sync()
		if err != nil {
			fmt.Println("Failed to sync logger", zap.Error(err))
		}
	}()
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	// Initialize repositories
	command, err := repository.NewPostgresCommand(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME))
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	query, err := repository.NewPostgresQuery(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB_REPLICA_USER, cfg.DB_REPLICA_PASSWORD, cfg.DB_REPLICA_HOST, cfg.DB_REPLICA_PORT, cfg.DB_REPLICA_NAME))
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	cache := repository.NewRedis(fmt.Sprintf("%s:%s", cfg.REDIS_HOST, cfg.REDIS_PORT))

	// Initialize service
	svc := service.NewService(command, query, cache)

	// Initialize Fiber app
	app = fiber.New()

	// Initialize and register routes
	handler := api.NewHandler(svc, logger)
	handler.RegisterRoutes(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Info("Handler invoked", zap.Any("request", req))

	// Convert APIGatewayProxyRequest to fasthttp.Request
	fasthttpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(fasthttpReq)
	fasthttpReq.Header.SetMethod(req.HTTPMethod)
	fasthttpReq.SetRequestURI(req.Path)
	for k, v := range req.Headers {
		fasthttpReq.Header.Set(k, v)
	}
	fasthttpReq.SetBody([]byte(req.Body))

	// Create fasthttp.RequestCtx
	requestCtx := &fasthttp.RequestCtx{}
	requestCtx.Init(fasthttpReq, nil, nil)

	// Handle the request with Fiber app
	app.Handler()(requestCtx)

	// Convert fasthttp.Response to APIGatewayProxyResponse
	resp := events.APIGatewayProxyResponse{
		StatusCode: requestCtx.Response.StatusCode(),
		Headers:    make(map[string]string),
		Body:       string(requestCtx.Response.Body()),
	}
	requestCtx.Response.Header.VisitAll(func(key, value []byte) {
		resp.Headers[string(key)] = string(value)
	})

	return resp, nil
}

func main() {
	if isLambda() {
		lambda.Start(Handler)
	}

	logger.Info("Starting server on port 3000")
	if err := app.Listen(":3000"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func isLambda() bool {
	// Logic to determine if running in Lambda environment
	return false
}
