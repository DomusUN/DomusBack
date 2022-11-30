package domain

type ServicesWorker struct {
	IdServices int
	Score      int
}

type WorkerMetadata struct {
	Description string           `json:"description" bson:"description"`
	NumberId    string           `json:"number_id" bson:"number_id"`
	Score       int              `json:"score" bson:"score"`
	ArrServices []ServicesWorker `json:"arr_services" bson:"arr_services"`
}
