package main

import (
	"github.com/Kerensky585/walletapi"
)

// TODO this  would not be required in a library module!!!! used for testing here to init endpoints, open the DB connection using settings.
func main() {

	if walletapi.InitEndPoints() {

		if walletapi.DbController("appsettings.json") {
			walletapi.Router.Run("localhost:8080")
		}
	}
}
