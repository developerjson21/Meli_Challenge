package dto

import "time"

type RepositoryDTO struct {
	RepositoryUrl   string
	File         string
	InformationType string
	LinesNumber   string
	DetectionDate  time.Time
	JobId		string
	RepositoryOwner string
	AmountDetections int
}

