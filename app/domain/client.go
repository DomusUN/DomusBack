package domain

type HistoryServices struct {
	Date  string `json:"date" bson:"date"`
	Score int    `json:"score" bson:"score"`
}

type ClientMetadata struct {
	Historyservices []HistoryServices `json:"arr_hservices" bson:"arr_hservices"`
}
