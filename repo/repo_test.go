package repo

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"test_container/models"
	"test_container/testhelpers"
)

type CustomerRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repo        *CustomerRepo
	ctx         context.Context
}

func (suite *CustomerRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = pgContainer

	// 传入一个 gorm 对象
	connStr := pgContainer.ConnectionString

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	repository := NewCustomerRepo(db)

	suite.repo = repository
}

func (suite *CustomerRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

// TestCreateCustomer 测试插入数据
func (suite *CustomerRepoTestSuite) TestCreateCustomer() {
	t := suite.T()

	assert.NoError(t, suite.repo.Create(suite.ctx, &models.Customer{
		Name:  "swy",
		Email: "123@qq.com",
		Age:   22,
	}))
	assert.NoError(t, suite.repo.Create(suite.ctx, &models.Customer{
		Name:  "abc",
		Email: "123@qq.com",
		Age:   23,
	}))
	assert.NoError(t, suite.repo.Create(suite.ctx, &models.Customer{
		Name:  "def",
		Email: "456@qq.com",
		Age:   24,
	}))
	assert.NoError(t, suite.repo.Create(suite.ctx, &models.Customer{
		Name:  "ghi",
		Email: "789@qq.com",
		Age:   25,
	}))
}

// TestGetCustomerById 测试根据 id 查询
func (suite *CustomerRepoTestSuite) TestGetCustomerById() {
	t := suite.T()

	tests := []struct {
		id      int64
		want    *models.Customer
		wantErr bool
	}{
		{1, &models.Customer{1, "swy", 22, "123@qq.com", ""}, false},
		{2, &models.Customer{2, "abc", 23, "123@qq.com", ""}, false},
		{3, &models.Customer{3, "def", 24, "456@qq.com", ""}, false},
		{4, &models.Customer{4, "ghi", 25, "789@qq.com", ""}, false},
		{5, &models.Customer{5, "aaa", 26, "123@qq.com", ""}, true},
	}

	for _, tt := range tests {
		got, err := suite.repo.FindOne(suite.ctx, tt.id)

		if err != nil {
			if tt.wantErr {
				assert.Equal(t, err, gorm.ErrRecordNotFound)
				continue
			} else {
				assert.NoError(t, err)
			}
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Get() got = %v, want %v", got, tt.want)
		}
	}
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
