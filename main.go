package main

import (
	"sync"
	"fmt"
	"log"
	"vamshi/Utils"
	"vamshi/Api"
	"vamshi/Models"
	"github.com/aws/aws-sdk-go/aws/session"

	// "reflect"
)
var wg sync.WaitGroup
var mx sync.Mutex

func main(){
	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
	vol_ids, _ := Utils.AWS_resource_reader_from_txt("/home/abacus/cloud_trial_dump_project/Dumps/volume_ids.txt")
	log.Println("Successfully read the file")

	dict := make(map[string]Models.EventJson)
	// fmt.Println("length of vols is ",len(vol_ids)) is 6
	for  index, value := range vol_ids{
		fmt.Println(index)

		wg.Add(1)
		go func(index int ,value string){ // passing index and value as arguments is necessary if not index and value remain global variables and we end up having only the results of index 5 because index self update itself from 0 to 5 and go routine takes the global index value which is 5 
			defer wg.Done()
			event , snap_id := Api.OriginalVolEventNameTimeResource(sess,value)
			// you can also do this way .. below one 


			event.EventResource.(map[string]interface{})[snap_id] = Api.CreatedSnapEventNameTimeResource(sess,snap_id) // this is called assert typing any questions go to https://github.com/vamshikrishna2001/GO-TRYOUTS/blob/main/Compulsory-scenarios-of-assert-types
			// h := make(map[string]interface{})
			// h[snap_id] = Api.CreatedSnapEventNameTimeResource(sess,snap_id)
			// event.EventResource = h
			
			mx.Lock()
			dict[value] = event
			mx.Unlock()
		}(index, value)

		
	}
	wg.Wait()

	Utils.Json_writer("/home/abacus/cloud_trial_dump_project/Dumps/vol_to_snap.json",dict)
	
	// c := Api.CreatedSnapEventNameTimeResource(sess,"snap-0f803e97f882751f1")
	// fmt.Println(c)

}

