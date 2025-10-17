package user

import (
	"os"
)

func AuthenticateAdmin(username, password string) bool {
	envUser := os.Getenv("ADMIN_USERNAME")
	envPass := os.Getenv("ADMIN_PASSWORD")

	return username == envUser && password == envPass
}
