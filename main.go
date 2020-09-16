package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"krungthai.com/khanapat/backend-web-big-query/analyze"
	"krungthai.com/khanapat/backend-web-big-query/database"
)

func init() {
	runtime.GOMAXPROCS(1)
	initViper()
	initLogConfig()
}

func initViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic("There is no such a config file in path ./config/config.yaml")
		} else {
			log.Panic("There is some problem about data in file")
		}
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func initLogConfig() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"

	config := zap.NewProductionConfig()
	var logLevel zapcore.Level
	switch viper.GetString("log.level") {
	case "info":
		logLevel = zapcore.InfoLevel
	case "debug":
		logLevel = zapcore.DebugLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		log.Fatal("There is no log level config")
	}
	config.Level = zap.NewAtomicLevelAt(logLevel)
	if viper.GetString("log.env") == "dev" {
		config.Encoding = "console"
	} else {
		config.Encoding = "json"
	}
	config.EncoderConfig = encoderConfig

	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
}

func main() {
	r := mux.NewRouter()
	rWithPrefix := r.PathPrefix("/bootcamp/data").Subrouter()

	db := database.MssqlConn()
	defer db.Close()

	dbInquiryMerchantByLatLong := analyze.NewInquiryMerchantByLatLongFn(db)

	analyzeHanlder := analyze.NewHandler(
		analyze.NewGetMerchantFn(dbInquiryMerchantByLatLong),
	)

	rWithPrefix.HandleFunc("/inquiry", analyzeHanlder.Inquiry).Methods("POST")

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", viper.GetString("app.port")),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	zap.L().Info(fmt.Sprintf("⇨ http server started on [::]:%s", viper.GetString("app.port")))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			zap.L().Fatal("cannot start server",
				zap.Error(err))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("app.waitTimeout"))
	defer cancel()

	srv.Shutdown(ctx)

	zap.L().Info("shutting down")
	os.Exit(0)
}
