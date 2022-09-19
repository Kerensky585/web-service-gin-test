# web-service-gin-test
For testing the wallet api GO code
Checkout and then in the project folder run with:

...> go run .

This will start the endpoints and connect to the local file DB, see the `appsettings.json` file for the settings.

Format / available end points are listed in [GIN-Debug], but in case, using Curl:

`...$> curl http://localhost:8080/api/v1/wallets/1/balance`


`...$> curl http://localhost:8080/api/v1/wallets/1/credit` --include --header "Content-Type: application/json" --request "POST" --data "{\"wid\": \"1\" ,\"amount\": \"100.00\"}"


`...$> curl http://localhost:8080/api/v1/wallets/1/debit` --include --header "Content-Type: application/json" --request "POST" --data "{\"wid\": \"1\" ,\"amount\": \"100.00\"}"

were `.../wid/...` is the integer wallet ID, i.e. `.../1/balance...`, `.../1/debit...`, or `..,/1/credit...`