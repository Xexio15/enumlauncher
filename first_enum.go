package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
)

type machine struct {
    ip string
    nmap_done bool
    gobuster_instances int
    gobuster_done bool
}

func launch_scan(machine machine){
    exec.Command("mkdir " + machine.ip)
    exec.Command("nmap -sC -sV -O -p- -oA " + machine.ip + "/nmap " + machine.ip)
    machine.gobuster_done = true
    exec.Command("grep 'http' nmap.gmap")


}
func main() {
    var current_scans [100]string :=ine, 0)
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    //READ
        m := machine{ip:scanner.Text(), nmap_done: false, gobuster_instances: 0, gobuster_done: false}
        current_scans = append(m)
        go launch_scan(m)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
