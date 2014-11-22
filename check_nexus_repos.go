package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
import  "flag"
import "os"

type HandledJson interface {
    Handler(url string, verbose bool) int
}

type Repositories struct {

    Data []Repository

}

type Repository struct {
    Exposed bool
    Id string
    Name string
    RepoType string
}

func (list Repositories) Handler (url string, verbose bool) (items int) {

    for i, repo := range list.Data {
        // items += 1
        if verbose {

            fmt.Fprintf(os.Stderr, "%d:name:%s:Id:%s:Exposed:%t:\n", i, repo.Name, repo.Id, repo.Exposed)

        }

        // an error here means nexus is down or servien bad json
        //so raise an unknown
        rv, _ := get_content(url + "/" + repo.Id + "/status", new(RepositoryState), verbose)
        // if err != nil {
        //     continue
        // }

        items += rv

    }

    return

}

type RepositoryState struct {

    Data _RepositoryState

}

type _RepositoryState struct {
    ProxyMode string
    Format string
    RepoType string
    LocalStatus string
    Id string
    RemoteStatus string
}

func (state RepositoryState) Handler (url string, verbose bool) (items int) {

    s := state.Data
    if verbose {

        fmt.Fprintf(os.Stderr, "repo:ProxyMode:%s:Id:%s:RemoteStatus:%s:LocalStatus:%s:RepoType:%s:Format:%s:\n", s.ProxyMode, s.Id, s.RemoteStatus, s.LocalStatus, s.RepoType, s.Format)
    
    }

    //repo:ProxyMode:BLOCKED_AUTO:Id:NPM:RemoteStatus:UNAVAILABLE:LocalStatus:IN_SERVICE:RepoType:proxy:Format:maven2:
    if s.ProxyMode == "BLOCKED_AUTO" {
        return 1
    }

    return 0

}

func get_content(url string, data HandledJson, verbose bool) (items int, err error) {

    defer func() {
        if err := recover(); err != nil {
            return
        }
    }()

    if verbose {
        fmt.Fprintf(os.Stderr, "fetching:url:%s:\n", url)
    }

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return
    }

    req.Header.Add("Accept", "application/json")
    res, err := client.Do(req)
    if err != nil {
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return
    }

    err = json.Unmarshal(body, data)
    if err != nil {

        if verbose {

            switch v := err.(type) {
                case *json.SyntaxError:
                    fmt.Fprintf(os.Stderr, string(body[v.Offset-40:v.Offset]))
            }
            
        }
        return
    }

    items = data.Handler(url, verbose)

    return
}

func main() {

    status := "OK"
    rv := 0
    name := "JenkinsQueue"

    // c := make(chan int)

    verbose := flag.Bool("v", false, "verbose output")
    warn := flag.Int("w", 10, "warning level for job queue depth")
    crit := flag.Int("c", 20, "critical level for job queue depth")
    host := flag.String("h", "http://gec-maven-nexus.walmart.com/nexus/service/local", "base url for nexus api like http://gec-maven-nexus.walmart.com/nexus/service/local")

    url := *host + "/repositories"

    flag.Parse()

    if len(flag.Args()) > 0 {

        flag.Usage()
        os.Exit(3)

    }

    defer func() {
        if err := recover(); err != nil {
            fmt.Println(name + " Unknown: ", err)
            os.Exit(3)
        }
    }()

    if *verbose {

        fmt.Printf("checking repos on:%s:warning:%d:critical:%d\n", url, *warn, *crit)
        
    }

    jobs, err := get_content(url, new(Repositories), *verbose)
    if err != nil {

        fmt.Printf("%s Unknown: %T %s %#v\n", name, err, err, err)

        os.Exit(3)

    }

    if jobs >= *crit {
        status = "Critical"
        rv = 1
    } else if jobs >= *warn {
        status = "Warning"
        rv = 2
    } 

    fmt.Printf("%s %s: %d\n", name, status, jobs)
    os.Exit(rv)

}
