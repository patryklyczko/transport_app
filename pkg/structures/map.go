package structures

type MapRequest struct {
	Path string `json:"path" bson:"path"`
}

type NodePositions struct {
	Parent   int64    `json:"parent" bson:"parent"`
	Position Position `json:"position" bson:"position"`
}

type NodesRelations struct {
	Parent   int64   `json:"parent" bson:"parent"`
	Children []int64 `json:"children" bson:"children"`
	MaxSpeed string  `json:"max_speed" bson:"max_speed"`
}
