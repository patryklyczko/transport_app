package structures

type Driver struct {
	ID       string   `json:"id" bson:"id"`
	Name     string   `json:"name" bson:"name"`
	Position Position `json:"position" bson:"position"`
	Orders   []Order  `json:"orders" bson:"orders"`
	Capacity float32  `json:"capacity" bson:"capacity"`
}

type DriverRequest struct {
	Name     string   `json:"name" bson:"name"`
	Position Position `json:"position" bson:"position"`
	Orders   []Order  `json:"orders" bson:"orders"`
	Capacity float32  `json:"capacity" bson:"capacity"`
}
