package model

import "time"

// TWork is a running Work
type TWork struct {
	Job      *TJob
	Created  time.Time
	Start    time.Time
	End      time.Time
	Replicas int
	Result   *TResult
}

// TWorkInstance ....
type TWorkInstance struct {
	Job     *TJob
	Work    *TWork
	Created time.Time
	Start   time.Time
	End     time.Time
	Node    string
	Result  *TResult
}

// TResult ...
type TResult struct {
}
