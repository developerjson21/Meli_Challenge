package domain

import "time"

type Entity struct {
	ID               int       `json:"ID" binding:"required"`
	RepositoryUrl    string    `json:"REPOSITORY" binding:"required"`
	File             string    `json:"FILE" binding:"required"`
	InformationType  string    `json:"INFORMATION_TYPE" binding:"required"`
	LinesNumber      string    `json:"NUMBER_OF_LINES" NUMBER_OF_LINES:"required"`
	DetectionDate    time.Time `json:"DETECTION_DATE" binding:"required"`
	JobId            string    `json:"JOB_ID" binding:"required"`
	RepositoryOwner  string    `json:"REPOSITORY_OWNER" binding:"required"`
	AmountDetections int       `json:"AMOUNT_DETECTIONS" binding:"required"`
}