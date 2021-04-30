package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type machine struct {
	ip                 string
	nmap_done          bool
	gobuster_instances int
	gobuster_done      bool
}

func launch_scan(machine machine) {
	fmt.Println(exec.Command("mkdir", machine.ip).Run())
	fmt.Println(exec.Command("nmap", "-sC", "-sV", "-O", "-p-", "-oA", machine.ip+"/nmap ", machine.ip).Output())
	machine.gobuster_done = true
	fmt.Println(exec.Command("grep", "'http'", "nmap.gmap").Run())

}
func main() {
	var current_scans [100]machine
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i = 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//READ
		m := machine{ip: scanner.Text(), nmap_done: false, gobuster_instances: 0, gobuster_done: false}
		current_scans[i] = m
		fmt.Println(m)
		i = i + 1
		go launch_scan(m)
	}
	for {
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
