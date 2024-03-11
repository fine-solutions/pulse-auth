package postgres

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"io"
	"os"
	"pulse-auth/cmd/pulse/config"
	"pulse-auth/internal/model"
	"pulse-auth/internal/utils"
	"testing"
)

var db *Storage

func TestMain(m *testing.M) {
	t := &testing.T{}
	setup(t)
	code := m.Run()
	teardown(t)
	os.Exit(code)
}

func teardown(t *testing.T) {
	dropTables(t)
}

func setup(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		log.Error("can't init logger")
	}

	cfg := config.StorageConfig{
		Username: "admin",
		Password: "admin123",
		Port:     5430,
		Name:     "test_db",
	}
	s, err := NewStorage(log, cfg)
	if err != nil {
		t.Fatalf("failed to create s: %v", err)
	}

	err = s.db.Ping()
	if err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}

	db = s
	createTables(t)
}

// TODO: сделать автоматические миграции через goose
func createTables(t *testing.T) {
	//fmt.Println(os.Getwd())
	sqlFile, err := os.Open("../../../migrations/20240306085507_init.sql")
	if err != nil {
		t.Fatal("Error opening SQL file:", err)
	}
	defer sqlFile.Close()
	queryBytes, err := io.ReadAll(sqlFile)
	if err != nil {
		t.Fatal("Error reading SQL file:", err)
	}
	query := string(queryBytes)

	_, err = db.db.Exec(query)
	if err != nil {
		t.Fatal("Error executing SQL query:", err)
	}
}

func dropTables(t *testing.T) {
	q := `DROP TABLE token_table; DROP TABLE user_table;`

	_, err := db.db.Exec(q)
	if err != nil {
		t.Fatal("Error executing SQL query:", err)
	}
}

func TestUser(t *testing.T) {
	username := "Nikita"
	password := "secretPassword"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Error("hash password:", err)
	}
	id := utils.GenerateUUID()
	u := &model.UserRegister{
		ID:             id,
		Username:       username,
		HashedPassword: hashedPassword,
	}

	userFromDb, err := db.CreateUser(context.Background(), u)
	if err != nil {
		t.Error("can't create user", err)
	}

	userLogin := &model.UserLogin{
		Username:       username,
		HashedPassword: hashedPassword,
	}

	userFromLogin, err := db.LoginUser(context.Background(), userLogin)
	if err != nil {
		t.Error("can't user login", err)
	}

	userById, err := db.GetUserByID(context.Background(), model.UserID(id))
	if err != nil {
		t.Error("can't get uesr by id", err)
	}

	assert.Equal(t, userFromDb, userFromLogin)
	assert.Equal(t, userFromDb, userById)
}