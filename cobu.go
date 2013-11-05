package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "net/http"
)

var (
    pPath string
    port  string
    addr  string
)

func init() {
    pPath = os.Getenv("PPATH")
    port  = os.Getenv("PORT")
    if pPath == "" {
        fmt.Println("PPATH environment variable not set.")
        os.Exit(1)
    }
    if port == "" {
        port = "12345"
    }
    addr = fmt.Sprintf(":%s", port)
}

func updateCodebase() {
    cmd     := exec.Command("git", "pull")
    cmd.Dir  = pPath
    err     := cmd.Start()
    if err != nil {
        log.Fatal("Error while running git: ", err)
    }
    log.Printf("Updating code base on %s", pPath)
    err = cmd.Wait()
    if err != nil {
        log.Fatal("Error while updating code base: ", err)
    } else {
        out, _ := cmd.Output()
        log.Printf("Code base updated\n%s", out)
    }
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
    log.Println("Handling deployment request")
    go updateCodebase()
    fmt.Fprint(w, "ok")
}

func main() {
    http.HandleFunc("/deploy", handleRequest)
    log.Printf("About to listen on %s\n", addr)
    err := http.ListenAndServe(addr, nil)
    if err != nil {
        log.Fatal("Error while starting server: ", err)
    }
}
