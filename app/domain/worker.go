package domain

type ServicesWorker struct {
	IdServices int     `json:"service_id" bson:"service_id"`
	Score      float32 `json:"score" bson:"score"`
	Price      int32   `json:"price"  bson:"price"`
}

type WorkerMetadata struct {
	NumberId    string           `json:"number_id" bson:"number_id"`
	Description string           `json:"description" bson:"description"`
	Score       float32          `json:"score" bson:"score"`
	Places      []string         `json:"places" bson:"places"`
	ArrServices []ServicesWorker `json:"arr_services" bson:"arr_services"`
}
