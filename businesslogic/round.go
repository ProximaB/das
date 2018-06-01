package businesslogic

import "time"

type RoundOrder struct {
	Name string
	Rank int
}

var (
	ROUND_ORDER_FINAL   = RoundOrder{"Final", 1}
	ROUND_ORDER_SEMI    = RoundOrder{"Semi-Final", 2}
	ROUND_ORDER_QUARTER = RoundOrder{"Quarter Final", 3}
	ROUND_ORDER_1_8     = RoundOrder{"1/8 Final", 4}
	ROUND_ORDER_1_16    = RoundOrder{"1/16 Final", 5}
	ROUND_ORDER_1_32    = RoundOrder{"1/32 Final", 6}
	ROUND_ORDER_1_64    = RoundOrder{"1/64 Final", 7}
	ROUND_ORDER_1_128   = RoundOrder{"1/128 Final", 8}
	ROUND_ORDER_1_256   = RoundOrder{"1/256 Final", 9}
	ROUND_ORDER_1_512   = RoundOrder{"1/512 Final", 10}
	ROUND_ORDER_1_1024  = RoundOrder{"1/1024 Final", 11}
	ROUND_ORDER_1_2048  = RoundOrder{"1/2048 Final", 12}
)

type Round struct {
	RoundID         int
	EventID         int
	Order           RoundOrder
	Entries         []EventEntry
	StartTime       time.Time
	EndTime         time.Time
	DateTimeCreated time.Time
	CreateUserID    int
	DateTimeUpdated time.Time
	UpdateUserID    int
}
