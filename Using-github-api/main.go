/*

- Fetches info about the user from github API ( These requests happen concurrently )
- Stores the data in form of slices and structs
- Exposes the data using http endpoints

*/



package main

import (
    "fmt"
    "io/ioutil"
    "sync"
    "net/http"
    "encoding/json"
    "strings"
)

// Struct 

type jsonUser struct {
    Name string `json:"login"`
    Id int `json:"id"`
    PublicRepos int `json:"public_repos"`
  }

  var results []jsonUser
  var names []string
  var ids, repos []int

// utility method

func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

// method that calls Github API
func getUser(username string, ch chan<- jsonUser, wg *sync.WaitGroup) {
   
	defer wg.Done()
	
    res, err := http.Get("https://api.github.com/users/"+username)
    if err != nil {
        fmt.Println("Error", err)
        return
    }

    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
   
    if err != nil {
        fmt.Println("Error", err)
        return
    }

    if res.StatusCode != 200 {
        fmt.Println("Unexpected Status Code : ", err)
        return
    }

    user := jsonUser{}

    err = json.Unmarshal(body, &user)
   
    if err != nil {
        fmt.Println("Error", err)
        return
    }
    
    // Storing the user struct in channel
    ch <- user
}


func main() {
    ch := make(chan jsonUser)

	var username string
	var usernames = []string{"emmenko","gaearon","bebraw"}
    var wg sync.WaitGroup
	
	for _, username = range usernames {
        wg.Add(1)
        go getUser(username, ch, &wg)
    }

    // closing the channel 

    go func() {
        wg.Wait()
        close(ch)
	}()
	
    // reading from channel 
    for res := range ch {
        results = append(results, res)
        names = append(names, res.Name)
        ids = append(ids, res.Id)
        repos = append(repos, res.PublicRepos)
    }

    // Endpoints

    http.HandleFunc("/", index)
    http.HandleFunc("/names", displayNames)
    http.HandleFunc("/ids", displayIds)
    http.HandleFunc("/public-repos", displayPublicRepos)

    http.ListenAndServe(":3000",nil)

}

func index(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, `
    Welcome !
        Goto /names to get the names
        Goto /ids to get the ids
        Goto /public-repos to get the number of public repos
    `)
}

func displayNames(w http.ResponseWriter, r *http.Request){

    namesString := strings.Join(names, ",")
    fmt.Fprintf(w, namesString)
}


func displayIds(w http.ResponseWriter, r *http.Request){

    idStrings := arrayToString(ids,",")
    fmt.Fprintf(w, idStrings)

}

func displayPublicRepos(w http.ResponseWriter, r *http.Request){

    repoStrings := arrayToString(repos,",")
    fmt.Fprintf(w, repoStrings)

}
