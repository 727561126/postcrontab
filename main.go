package main

import (
	"bytes"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"os/exec"
)

var c = cron.New()

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")

}

func addJob(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form["jobtime"])
	fmt.Println(r.Form["jobvalue"])
	spec := "*/1 * * * * * "
	c.AddFunc(spec, func() {
		log.Println("cron running")
//		exec_shell("/bin/bash", "-c", "ls -rlt")
		exec_shell("php", "-v", "-c")
	})

}

func exec_shell(s, t, x string) {
	cmd := exec.Command(s,t,x)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func startJob(w http.ResponseWriter, r *http.Request) {
	c.Start()

}
func stopJob(w http.ResponseWriter, r *http.Request) {
	c.Stop()
}

func main() {

	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/addjob", addJob)
	http.HandleFunc("/startjob", startJob)
	http.HandleFunc("/stopjob", stopJob)
	er := http.ListenAndServe(":9000", nil)
	if er != nil {
		log.Fatal("ListenAndServe:", er)
	}

}
