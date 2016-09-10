package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/julienschmidt/httprouter"
	"github.com/patito/bidu/model"
	"github.com/stretchr/testify/require"
)

func TestCreateSite(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error: %s ", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO site")
	mock.ExpectExec("INSERT INTO site").
		WithArgs("4D", "4d", "address", "comments").
		WillReturnResult(sqlmock.NewResult(0, 1))

	m := model.New(db)
	h := New(m)

	body := strings.NewReader(`{"name": "4D","slug": "4d","comments": "comments","physicaladdress": "address"}`)

	resp := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/sites", body)
	if err != nil {
		t.Fatal(err)
	}

	router := httprouter.New()
	router.Handle("POST", "/api/sites", h.CreateSite)
	router.ServeHTTP(resp, req)
	require.Equal(t, 201, resp.Code)
}
