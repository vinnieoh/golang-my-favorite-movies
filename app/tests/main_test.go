package tests

import (
    "log"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

    if err := godotenv.Load("./dotenv_files/.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    exitCode := m.Run()

    os.Exit(exitCode)
}

func TestEndpoint(t *testing.T) {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "Hello, World!", w.Body.String())
}