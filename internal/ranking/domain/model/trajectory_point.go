package model

import "time"

type TrajectoryPoint struct {
	SnapshotDate     time.Time
	Position, Points uint16
}
