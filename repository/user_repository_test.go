package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	ur := repository.NewUserRepository(sqlxDB)

	mockUser := &domain.User{
		ID:       "1",
		Name:     "TestUser",
		Password: "password",
		Email:    "testuser@gmail.com",
	}

	t.Run("success", func(t *testing.T) {
		query := `INSERT INTO users $begin:math:text$id, name, password, email$end:math:text$ VALUES $begin:math:text$\\?, \\?, \\?, \\?$end:math:text$`
		mock.ExpectExec(query).
			WithArgs(mockUser.ID, mockUser.Name, mockUser.Password, mockUser.Email).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := ur.Create(context.Background(), mockUser)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		query := `INSERT INTO users $begin:math:text$id, name, password, email$end:math:text$ VALUES $begin:math:text$\\?, \\?, \\?, \\?$end:math:text$`
		mock.ExpectExec(query).
			WithArgs(mockUser.ID, mockUser.Name, mockUser.Password, mockUser.Email).
			WillReturnError(sql.ErrConnDone)

		err := ur.Create(context.Background(), mockUser)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	ur := repository.NewUserRepository(sqlxDB)

	mockUser := domain.User{
		ID:       "1",
		Name:     "TestUser",
		Password: "password",
		Email:    "testuser@gmail.com",
	}

	t.Run("success", func(t *testing.T) {
		query := `SELECT id, name, password, email FROM users WHERE email = \?`
		rows := sqlmock.NewRows([]string{"id", "name", "password", "email"}).
			AddRow(mockUser.ID, mockUser.Name, mockUser.Password, mockUser.Email)

		mock.ExpectQuery(query).WithArgs(mockUser.Email).WillReturnRows(rows)

		user, err := ur.GetByEmail(context.Background(), mockUser.Email)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		query := `SELECT id, name, password, email FROM users WHERE email = \?`
		mock.ExpectQuery(query).WithArgs(mockUser.Email).WillReturnError(sql.ErrConnDone)

		_, err := ur.GetByEmail(context.Background(), mockUser.Email)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	ur := repository.NewUserRepository(sqlxDB)

	mockUser := domain.User{
		ID:       "1",
		Name:     "TestUser",
		Password: "password",
		Email:    "testuser@gmail.com",
	}

	t.Run("success", func(t *testing.T) {
		query := `SELECT id, name, password, email FROM users WHERE id = \?`
		rows := sqlmock.NewRows([]string{"id", "name", "password", "email"}).
			AddRow(mockUser.ID, mockUser.Name, mockUser.Password, mockUser.Email)

		mock.ExpectQuery(query).WithArgs(mockUser.ID).WillReturnRows(rows)

		user, err := ur.GetByID(context.Background(), mockUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		query := `SELECT id, name, password, email FROM users WHERE id = \?`
		mock.ExpectQuery(query).WithArgs(mockUser.ID).WillReturnError(sql.ErrConnDone)

		_, err := ur.GetByID(context.Background(), mockUser.ID)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}