package Models

import (
	"time"
)

type EventJson struct{
	EventName string
	EventTime time.Time
	EventResource interface{}
}
