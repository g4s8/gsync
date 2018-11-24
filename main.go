package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const (
	envRepo = "GSYNC_REPO"
	envCred = "GSYNC_CREDENTIALS"
	envDest = "GSYNC_DESTINATION"
)

func getOrFail(name string) string {
	res := os.Getenv(name)
	if res == "" {
		log.Fatal(fmt.Sprintf("%s required", name))
	}
	return res
}

func gsync(repo string, dest string, cred string) ([]byte, error) {
	return exec.Command("/scripts/gsync.sh", repo, dest, cred).Output()
}

func main() {
	repo := getOrFail(envRepo)
	cred := getOrFail(envCred)
	dest := getOrFail(envDest)
	gsync(repo, dest, cred)
	http.HandleFunc("/sync", func(w http.ResponseWriter, r *http.Request) {
		out, err := gsync(repo, dest, cred)
		fmt.Printf("gsync.sh:\n%s", out)
		if err != nil {
			fmt.Fprintf(w, "Error occured: %s", err)
		} else {
			fmt.Fprintf(w, "Thanks, we will process your request")
		}
	})
	log.Fatal(http.ListenAndServe(":8042", nil))
}
