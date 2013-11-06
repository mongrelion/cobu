package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "regexp"
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
    // TODO: Don't die if updating codebase update fails.
    if err != nil {
        log.Fatal("Error while updating code base: ", err)
    } else {
        out, _ := cmd.Output()
        log.Printf("Code base updated\n%s", out)
    }
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
    if requestComesFromGithub(req) {
        log.Println("Handling deployment request")
        go updateCodebase()
        fmt.Fprint(res, "ok")
    } else {
        http.NotFound(res, req)
    }
}

func requestComesFromGithub(req *http.Request) bool {
    r := regexp.MustCompile(`192\.30\.252\.\d{1,3}`)
    return r.MatchString(req.Header.Get("X-Remote-IP"))
}

func main() {
    http.HandleFunc("/deploy", handleRequest)
    log.Printf("About to listen on %s\n", addr)
    err := http.ListenAndServe(addr, nil)
    if err != nil {
        log.Fatal("Error while starting server: ", err)
    }
}
