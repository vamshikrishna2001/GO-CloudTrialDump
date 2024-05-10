package Models

import (
	"time"
)

type EventJson struct{
	EventName string
	EventTime time.Time
	EventResource interface{}
}


type ChannelCombiner struct{
	Vol_id string
	Snap_id string 
	Event EventJson

}