package main

type StoreRequest struct {
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
type StoreResponse struct {
	Success          bool `json:"success"`
	StoredTokenCount int  `json:"stored_token_count"`
}
type RetrieveRequest struct {
	UserId  string `json:"user_id"`
	Recency int    `json:"hours"`
	Size    int    `json:"size"`
	Type    int    `json:"type"` // 0 --> Both of Theme, 1 --> Standard Only, 2 --> Keywords Only
}
type Token struct {
	UserId     string `json:"user_id"`
	LastUpdate string `json:"last_update"`
	Token      string `json:"token"`
	IsKeyword  bool   `json:"is_keyword"`
	Occurs     *int64 `json:"occurs"`
}
