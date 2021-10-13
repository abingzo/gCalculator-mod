package test

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestCmdStd(t *testing.T) {
	cmd := exec.Command("./eval")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = strings.NewReader("1+2")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 2)
}
