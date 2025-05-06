package main

import (
	"fmt"

	// migration "github.com/awaisniaz/user-management/pkg/db/migration"
	connection "github.com/awaisniaz/user-management/pkg/db"
)

func main() {
	fmt.Printf("Server is running on PORT: 3000")
	// migration.Up()
	connection.InitConnection()

}
