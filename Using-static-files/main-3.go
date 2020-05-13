package main

import (
    "fmt"
    "io/ioutil"
	"sync"
)

// Method that reads from static text files

func readFile(fileNo string, ch chan<- string, wg *sync.WaitGroup) {
   
	defer wg.Done()
	
	data, err := ioutil.ReadFile("file-"+fileNo+".txt")
	if err != nil {
		fmt.Println("File reading error", err)
    	return
	  }

	// Saving the data to channel
    ch <- string(data)
}

// Reads data concurrently from files and stores it in an array

func main() {
    ch := make(chan string)
    var results []string
	var FileNum string
	var FileNums = []string{"a","b","c"}
    var wg sync.WaitGroup
	
	for _, FileNum = range FileNums {
        wg.Add(1)
        go readFile(FileNum, ch, &wg)
    }

    // close the channel in the background
    go func() {
        wg.Wait()
        close(ch)
	}()
	
    // reading from channel 
    for res := range ch {
        results = append(results, res)
    }

	fmt.Println(results)

    
}
