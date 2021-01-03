package models

type Uploadtar struct {
	Timeframe int64
	Target int
	Target_reached bool
	Serve_historical_blocks bool
	Bytes_left_in_cycle int
	Time_left_in_cycle int


}
