package data

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventId          int // PK
	EventName        string
	EventDescription string
	EventLocation    string
	EventDate        time.Time
}

type Diagram struct {
	gorm.Model
	DiagramId          int // PK
	EventId            int // FK
	DiagramName        string
	DiagramDescription string
}

type Test struct {
	gorm.Model
	TestId    int // PK
	DiagramId int // FK
}
