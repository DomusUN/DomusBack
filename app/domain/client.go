package domain

type HistoryServices struct {
	Date  string
	Score int
}

type ClientMetadata struct {
	Historyservices []HistoryServices `json:"arr_services" bson:"arr_services"`
}
