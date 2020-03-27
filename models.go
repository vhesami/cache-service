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
	UserId    string    `json:"user_id"`
	Recency   int       `json:"hours"`
	Size      int       `json:"size"`
	Type      TokenType `json:"type"`
	Delimiter string    `json:"delimiter"`
}
type Token struct {
	UserId     string `json:"user_id"`
	LastUpdate string `json:"last_update"`
	Token      string `json:"token"`
	IsKeyword  bool   `json:"is_keyword"`
	Occurs     *int64 `json:"occurs"`
}

type TokenType int

const (
	STANDARD TokenType = 1
	KEYWORD  TokenType = 2
	BOTH     TokenType = 3
)

func (r *RetrieveRequest) Compact() {
	if r.Size == 0 {
		r.Size = 10
	}
	if r.Recency == 0 {
		r.Recency = 1
	}
	if r.Delimiter == "" {
		r.Delimiter = ","
	}
	if r.Type == 0 {
		r.Type = STANDARD
	}
}
