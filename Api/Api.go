package Api

import (
	"time"
	"log"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudtrail"
	"vamshi/Models"
    // "reflect"
)


func Api_function(){
	fmt.Println("hello form api1")
}


func OriginalVolEventNameTimeResource(sess *session.Session,resource_id string) (Models.EventJson, string) {

    // Initialize AWS session
    // sess := session.Must(session.NewSessionWithOptions(session.Options{
    //     SharedConfigState: session.SharedConfigEnable,
    // }))

    // Create CloudTrail service client
    svc := cloudtrail.New(sess)

    // Get resource name from command line arguments
    resourceName := resource_id

    // Prepare input parameters for LookupEvents API call
    params := &cloudtrail.LookupEventsInput{
        LookupAttributes: []*cloudtrail.LookupAttribute{
            {
                AttributeKey:   aws.String("ResourceName"),
                AttributeValue: aws.String(resourceName),
            },
        },
    }

    // Retrieve events associated with the resource name
    resp, err := svc.LookupEvents(params)
    if err != nil {
		log.Fatalf("service client create cheyyalekapoyam ayya in vol")
    }

    var Event Models.EventJson;
    fmt.Println("Events associated with resource:", resourceName)
    // fmt.Println(reflect.TypeOf(resp.Events))
    // fmt.Println(resp.Events)
    var snap_id string;
    for _, event := range resp.Events {
        fmt.Println("Event Name:", *event.EventName)
        fmt.Println("Event Time:", *event.EventTime)
        fmt.Println("Event Source:", *event.EventSource)
        fmt.Println("-----------------------------")
        snap_id = *event.Resources[0].ResourceName
        snapshot_map := make(map[string]interface{})
        snapshot_map[snap_id] = nil
        Event = Models.EventJson{EventName:*event.EventName,EventTime:*event.EventTime,EventResource:snapshot_map}
        // fmt.Println(event)
        // fmt.Println(reflect.TypeOf(event))
    }
	return Event,snap_id
}


func CreatedSnapEventNameTimeResource(sess *session.Session,resource_id string) (interface{}) {

    // Initialize AWS session
    // sess := session.Must(session.NewSessionWithOptions(session.Options{
    //     SharedConfigState: session.SharedConfigEnable,
    // }))

    // Create CloudTrail service client
    svc := cloudtrail.New(sess)

    // Get resource name from command line arguments
    resourceName := resource_id

    // Prepare input parameters for LookupEvents API call
    params := &cloudtrail.LookupEventsInput{
        LookupAttributes: []*cloudtrail.LookupAttribute{
            {
                AttributeKey:   aws.String("ResourceName"),
                AttributeValue: aws.String(resourceName),
            },
        },
    }

    // Retrieve events associated with the resource name
    resp, err := svc.LookupEvents(params)
    if err != nil {
		log.Fatalf("service client create cheyyalekapoyam ayya in snapshot")
    }


    created_snapshot_map := make(map[string]time.Time)
    for _, event := range resp.Events {

        if *event.EventName == "CopySnapshot"{
            created_snapshot_map[*event.EventName] = *event.EventTime
        }

        if *event.EventName == "DeleteSnapshot"{
            created_snapshot_map[*event.EventName] = *event.EventTime
        }
    }
    return created_snapshot_map
}
