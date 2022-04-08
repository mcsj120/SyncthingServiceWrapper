package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	cmd := exec.Command("D:\\Applications\\syncthing-windows-amd64-v1.18.1\\syncthing.exe")
	KillRunningProcess("syncthing.exe")
	err := cmd.Run()
	if err != nil {
		logger.Error(err)
	}
}
func (p *program) Stop(s service.Service) error {
	logger.Info(KillRunningProcess("syncthing.exe"))
	// Stop should not block. Return with a few seconds.
	return nil
}

//Sample windows command: tasklist /FI "IMAGENAME eq syncthing.exe"
func KillRunningProcess(name string) bool {
	cmd := exec.Command("TASKLIST", "/FI", fmt.Sprintf("IMAGENAME eq %s", name))
	result, err := cmd.Output()
	lines := strings.Split(string(result), "\n")
	logger.Info(lines)
	if len(lines) >= 5 {
		itr := 3
		for itr < len(lines) {
			process_arr := strings.Split(lines[itr], " ")
			process_itr := 1
			for process_itr < len(process_arr) {
				_, err := strconv.Atoi(process_arr[process_itr])
				if err == nil {
					cmd = exec.Command("TASKKILL", "/F", "/PID", process_arr[process_itr])
					result, err = cmd.Output()
					break
				}
				process_itr = process_itr + 1
			}
			itr += 2
		}
	}
	if err != nil {
		return false
	}
	return !bytes.Contains(result, []byte("No tasks are running"))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	svcFlag := flag.String("service", "", "Control this binary as a service")
	flag.Parse()
	svcConfig := &service.Config{
		Name:        "SyncthingWrapper",
		DisplayName: "Syncthing Wrapper",
		Description: "This is used to run syncthing in the background",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		logger.Info("Running: ", *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
	} else {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
}
