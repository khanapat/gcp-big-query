package database

import (
	"context"
	"database/sql"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

func MssqlConn() *sql.DB {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&parseTime=true&charset=utf8", viper.GetString("mssql.username"), viper.GetString("mssql.password"), viper.GetString("mssql.host"), viper.GetString("mssql.schema"))
	zap.L().Debug(connString)
	dbType := fmt.Sprint(viper.GetString("mssql.type"))
	conn, err := sql.Open(dbType, connString)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("Cannot connect to mssql - %s", err))
	}
	zap.L().Debug("mssql OK!")

	return conn
}

func BigQueryConn() *bigquery.Client {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, viper.GetString("bigquery.projectID"), option.WithCredentialsJSON([]byte(viper.GetString("bigquery.credentials"))))
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("Cannot connect to BigQuery - %s", err))
	}
	return client
}
