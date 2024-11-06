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
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	// Initialize repositories
	postgresConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
	command, err := repository.NewPostgresCommand(postgresConnStr)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	query, err := repository.NewPostgresQuery(postgresConnStr)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	cache := repository.NewRedis(fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort))

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
	fasthttpReq.Header.SetMethod(req.HTTPMethod)
	fasthttpReq.SetRequestURI(req.Path)
	for k, v := range req.Headers {
		fasthttpReq.Header.Set(k, v)
	}
	fasthttpReq.SetBody([]byte(req.Body))

	// Create fasthttp.Response
	fasthttpResp := fasthttp.AcquireResponse()

	// Handle the request with Fiber app
	requestCtx := &fasthttp.RequestCtx{
		Request:  *fasthttpReq,
		Response: *fasthttpResp,
	}
	app.Handler()(requestCtx)

	// Convert fasthttp.Response to APIGatewayProxyResponse
	resp := events.APIGatewayProxyResponse{
		StatusCode: fasthttpResp.StatusCode(),
		Headers:    make(map[string]string),
		Body:       string(fasthttpResp.Body()),
	}
	fasthttpResp.Header.VisitAll(func(key, value []byte) {
		resp.Headers[string(key)] = string(value)
	})

	logger.Info("Handler response", zap.Any("response", resp))

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
