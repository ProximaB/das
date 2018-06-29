// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

type RoundOrder struct {
	Name string
	Rank int
}

var (
	RoundOrderFinal     = RoundOrder{"Final", 1}
	RoundOrderSemiFinal = RoundOrder{"Semi-Final", 2}
	RoundOrderQuarter   = RoundOrder{"Quarter Final", 3}
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

// Round defines the
type Round struct {
	ID              int
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
