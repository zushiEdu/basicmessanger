package types

type Token struct {
	ID        int    `json:"id,string"`
	Signature string `json:"signature,string"`
	Expiry    string `json:"expiry,string"`
}

type TokenRequest struct {
	ID int `json:"id,string"`
}
