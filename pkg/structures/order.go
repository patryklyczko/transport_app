package structures

import "time"

type UID struct {
	ID string `json:"id" bson:"id"`
}

type Position struct {
	Lat float32 `json:"lat" bson:"lat"`
	Lon float32 `json:"lon" bson:"lon"`
}

type Order struct {
	ID           string        `json:"id" bson:"id"`
	PositionTake Position      `json:"position_take" bson:"position_take"`
	PositionSend Position      `json:"position_send" bson:"position_send"`
	TimeAdd      time.Time     `json:"time_add" bson:"time_add"`
	TimeEnd      time.Time     `json:"time_end" bson:"time_end"`
	Gain         int64         `json:"gain" bson:"gain"`
	Weight       float32       `json:"weight" bson:"weight"`
	Split        bool          `json:"split" bson:"split"`
	Taken        bool          `json:"taken" bson:"taken"`
	TimePack     time.Duration `json:"time_pack" bson:"time_pack"`
	TimeFinish   time.Time     `json:"time_finish" bson:"time_finish"`
}

type OrderRequest struct {
	PositionTake Position      `json:"position_take" bson:"position_take"`
	PositionSend Position      `json:"position_send" bson:"position_send"`
	TimeAdd      time.Time     `json:"time_add" bson:"time_add"`
	TimeEnd      time.Time     `json:"time_end" bson:"time_end"`
	Gain         int64         `json:"gain" bson:"gain"`
	Weight       float32       `json:"weight" bson:"weight"`
	Split        bool          `json:"split" bson:"split"`
	TimePack     time.Duration `json:"time_pack" bson:"time_pack"`
}

type OrderAlgorithm struct {
	ID           string    `json:"id" bson:"id"`
	Gain         int64     `json:"gain" bson:"gain"`
	TimeFinish   time.Time `json:"time_finish" bson:"time_finish"`
	PositionTake Position  `json:"position_take" bson:"position_take"`
	PositionSend Position  `json:"position_send" bson:"position_send"`
	Weight       float32   `json:"weight" bson:"weight"`
	Taken        bool      `json:"taken" bson:"taken"`
	Priority     float32   `json:"priority" bson:"priority"`
}

type OrderPagination struct {
	OrderCount int     `json:"order_count"`
	Orders     []Order `json:"orders"`
}
