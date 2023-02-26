package repositories

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDatabaseConnection(t *testing.T) {
	var connString = os.Getenv("POSTGRES_CON_STRING")
	conn := NewDatabaseConnection(connString)
	rs, err := conn.Query("select * from prices")
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}
