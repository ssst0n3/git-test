package main

import "os/exec"

func main() {
    cmd := exec.Command("calc.exe")
    cmd.Run()
}
