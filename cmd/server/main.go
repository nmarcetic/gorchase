package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	gorchase "github.com/nmarcetic/gorchase"
	"github.com/nmarcetic/gorchase/pkg/logger"
	"github.com/nmarcetic/gorchase/server/accounts"
	accountsAPI "github.com/nmarcetic/gorchase/server/accounts/api"
	accountsStorage "github.com/nmarcetic/gorchase/server/accounts/store/mysql"
	database "github.com/nmarcetic/gorchase/server/mysql"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	listenHost     string
	listenPort     string
	helmConfigPath string
)

func main() {

	var (
		envDBHost = "GORCHASE_DB_HOST"
		envDBName = "GORCHASE_DB_NAME"
		envDBUser = "GORCHASE_DB_USER"
		envDBPass = "GORCHASE_DB_PASS"
		envDBPort = "GORCHASE_DB_PORT"
		defDBHost = "localhost"
		defDBName = "gorchase"
		defDBUser = "gorchase"
		defDBPass = "gorchase"
		defDBPort = "5432"
	)

	logger, err := logger.New()
	if err != nil {
		fmt.Println("Faild to initialize logger")
		os.Exit(1)
	}

	dbConfig := &database.Config{
		Host:     gorchase.Env(envDBHost, defDBHost),
		Password: gorchase.Env(envDBPass, defDBPass),
		Name:     gorchase.Env(envDBName, defDBName),
		User:     gorchase.Env(envDBUser, defDBUser),
		Port:     gorchase.Env(envDBPort, defDBPort),
	}

	// Init DB connection
	db, err := database.Init(*dbConfig)
	if err != nil {
		logger.Sugar().Panicf("Faild to connect to Database, with error: %s", err)
	}
	// Init accounts svc
	userRepository := accountsStorage.NewUserRepository(db)
	userSvc := accounts.NewUsersService(userRepository, *logger)

	pflag.CommandLine.StringVar(&listenHost, "addr", "0.0.0.0", "server listen addr")
	pflag.CommandLine.StringVar(&listenPort, "port", "8080", "server listen port")
	pflag.CommandLine.StringVar(&helmConfigPath, "config", "../charts/config.yaml", "Helm provider config path")

	// router
	router := gin.Default()
	if err != nil {
		fmt.Printf("Faild to initialize logger with error: %s", err)
	}
	// Gin uses zap logger
	glogger, _ := zap.NewProduction()
	defer glogger.Sync()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(glogger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(glogger, true))
	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	/********** Public Endpoints **********/
	publicV1 := router.Group("/api/v1")
	// Health check route
	publicV1.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Healthy"})
	})

	/********** Private Endpoints **********/
	// Register accounts routes
	accountsAPI.InitRouter(publicV1, userSvc)
	// Use accounts AuthN middleware fn to protect private routes
	router.Use(accountsAPI.AuthMiddleware(userSvc).MiddlewareFunc())

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", listenHost, listenPort),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Sugar().Panicf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
