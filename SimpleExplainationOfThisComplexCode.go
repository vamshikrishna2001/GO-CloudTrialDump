// // Online Go compiler to run Golang program online
// // Print "Try programiz.pro" message

// package main


// import (
// 	"fmt"
// 	"time"
// 	"sync"
	
// )
// var wg sync.WaitGroup
// var mx sync.Mutex

// func volume(vol string) chan string{
//     ch := make(chan string)
    
//     go func(){
//         time.Sleep(5*time.Second)
//         d := fmt.Sprintf("%s_%d", vol, 18)
//         ch <- d
//         close(ch)
//     }()
//     return ch

    
// }

// func snapshot(ch chan string) string{
//     time.Sleep(3*time.Second)
//     d := <-ch 
//     e := fmt.Sprintf("%s_%d", d, 18)
//     return e
    
// }

// func main() {
//     fmt.Println("Try programiz.pro")
//     l := []string{"a","b","c","d","e"}
    
//     var result []string;
//     for _ , value := range l{
//         wg.Add(1)
//         go func(value string){
//             defer wg.Done()
            
//             d := volume(value)
//             e := snapshot(d)
            
//             mx.Lock()
//             result = append(result , e)
//             mx.Unlock()
//         }(value)
//     }
//     wg.Wait()
//     fmt.Println(result)
// }

// output is 
// Try programiz.pro
// 5.001312442s
// [d_18_18 b_18_18 c_18_18 e_18_18 a_18_18]