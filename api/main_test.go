package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/SoufianeRep/tscit/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T) *Server {
	db := db.GetDB()

	fmt.Println(db)
	server, err := NewServer(db)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
