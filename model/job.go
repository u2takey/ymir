package model

import "time"

// TJob is a test job config, saved as a job
type TJob struct {
	Name          string
	Description   string   // Annotation
	Script        string   // ConfigMap
	Replicas      int      // NodesSelected
	NodesSelected []string // NodesSelected
	Status        string   // JobStatus
	Created       time.Time
}
