package main

import (
	"time"
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


// passing index and value as arguments is necessary if not index and value remain global variables and we end up having only the results of index 5 because index self update itself from 0 to 5 and go routine takes the global index value which is 5 
func VolumeRetriever(index int ,value string,sess *session.Session) (<-chan Models.ChannelCombiner)  { 
	ch := make(chan Models.ChannelCombiner)

	go func(){
		defer close(ch)
		fmt.Println(index,value)
		event , snap_id := Api.OriginalVolEventNameTimeResource(sess,value)
		ch_combiner	:= new(Models.ChannelCombiner)
		// fmt.Println("jblkblwdsk",ch_combiner.snap_id)

		ch_combiner.Snap_id = snap_id
		ch_combiner.Event = event
		ch_combiner.Vol_id = value
		ch <- *ch_combiner
	}()


	return ch
}

func SnapshotRetriver(ch <-chan Models.ChannelCombiner,sess *session.Session , dict map[string]Models.EventJson){
	channel_combiner := <- ch
	snap_id := channel_combiner.Snap_id
	event := channel_combiner.Event
	value := channel_combiner.Vol_id

	event.EventResource.(map[string]interface{})[snap_id] = Api.CreatedSnapEventNameTimeResource(sess,snap_id) // this is called assert typing any questions go to https://github.com/vamshikrishna2001/GO-TRYOUTS/blob/main/Compulsory-scenarios-of-assert-types
	// h := make(map[string]interface{})
	// h[snap_id] = Api.CreatedSnapEventNameTimeResource(sess,snap_id)
	// event.EventResource = h
	
	mx.Lock()
	dict[value] = event
	mx.Unlock()
}

func main(){
	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
	vol_ids, _ := Utils.AWS_resource_reader_from_txt("/home/abacus/cloud_trial_dump_project/Dumps/volume_ids.txt")
	log.Println("Successfully read the file")

	dict := make(map[string]Models.EventJson)
	// fmt.Println("length of vols is ",len(vol_ids)) is 6

	start := time.Now()
	for  index, value := range vol_ids{  
		wg.Add(1)
		go func(index int , value string){  // you need to give the parameters as the argument or else by the time new go routine is created index value is getting updated 
			defer wg.Done()
			ch := VolumeRetriever(index ,value,sess)
			SnapshotRetriver(ch,sess,dict)

		}(index , value)
	

	}
	wg.Wait()



	end := time.Now()

	diff := end.Sub(start)
	fmt.Println("Total time taken ",diff)
	Utils.Json_writer("/home/abacus/cloud_trial_dump_project/Dumps/vol_to_snap.json",dict)
	
	// c := Api.CreatedSnapEventNameTimeResource(sess,"snap-0f803e97f882751f1")
	// fmt.Println(c)

}

