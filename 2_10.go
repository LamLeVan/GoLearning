package main

//Price
// const IN_NORMS = 1000
// const OUT_NORMS = 2000
//
// //electronics in norms
// const GLUE = 60
// const PROTECTION = 90
// const HOUSE_HOLD = 200
// const MANUFACTURE = 450
//
type Day int

const (
	MONDAY Day = 1 + iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
