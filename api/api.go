package api

// COIN BALANCE PARAMS
type CoinBalanceParams struct {
	Username string
}

// COIN BALANCE RESPONSE (when succesful response)
type CoinBalanceResponse struct {
	Code int

	//ACCOUNT BALANCE
	Balance int64
}

// ERROR RESPONSE
type Error struct {
	//Error code
	Code int
	//Error message
	Message string
}
