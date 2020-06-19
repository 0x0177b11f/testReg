package testReg

//easyjson:json
type Data struct {
	Rank int    `json:"rank"`
	Body string `json:"body"`
}

//easyjson:json
type DataArray []Data
