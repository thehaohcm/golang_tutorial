package unit_testing

import (
	"golang_project/repositories"
    "testing"
    "fmt"
    // db "github.com/mattn/go-sqlite3"
    "github.com/stretchr/testify/assert"
    "os"
    "database/sql"
)

var (
	repo repositories.FriendConnectionRepository = repositories.New()
)

func TestMain(m *testing.M) {
    // os.Exit skips defer calls
    // so we need to call another function
    code, err := run(m)
    if err != nil {
        fmt.Println(err)
    }
    os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
    // pseudo-code, some implementation excluded:
    //
    // 1. create test.db if it does not exist
    // 2. run our DDL statements to create the required tables if they do not exist
    // 3. run our tests
    // 4. truncate the test db tables

    db, err := sql.Open("sqlite3", "file: golang_project.db")
    if err != nil {
        return -1, fmt.Errorf("could not connect to database: %w", err)
    }

    // truncates all test data after the tests are run
    defer func() {
        for _, t := range []string{ "friends", "subscribers", "users" } {
            _, _ = db.Exec(fmt.Sprintf("DELETE FROM %s", t))
        }

        db.Close()
    }()

    return m.Run(), nil
}

func TestInsertBook(t *testing.T) {
    // store := &BookStore{
    //     db: db,
    // }

    // b, err := repo.FindFriendsByEmail(context.TODO(), &Book{
    //     Title: "The Go Programming Language",
    //     AuthorID: 1,
    //     ISBN: "978-0134190440",
    //     Subject: "computers",
    // })

	result := repo.CreateFriendConnection([]string{
		"abc@def.com",
        "abc1@def.com",
	})

    // using https://github.com/stretchr/testify library for brevity
    // require.NoError(t, err)

    // assert.Equal(t, "The Go Programming Language", b.Name)
    // assert.Equal(t, 1, b.AuthorID)
    // assert.Equal(t, "978-0134190440", b.ISBN)s
    // assert.Equal(t, "computers", b.Subject)
    // assert.NotZero(t, b.ID)
    assert.Equal(t, true, result)
}