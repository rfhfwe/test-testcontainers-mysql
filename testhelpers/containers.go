package testhelpers

import (
	"context"
	"path/filepath"

	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

type PostgresContainer struct {
	*mysql.MySQLContainer
	ConnectionString string
}

func CreatePostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := mysql.Run(ctx,
		"mysql:8.0",
		mysql.WithDatabase("test"),
		mysql.WithUsername("root"),
		mysql.WithPassword("123456"),

		mysql.WithScripts(filepath.Join("/Users/tal/GolandProjects/test_container/sql", "db.sql")),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		MySQLContainer:   pgContainer,
		ConnectionString: connStr,
	}, nil
}
