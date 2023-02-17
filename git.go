package main

import "io/ioutil"

func main() {
    ioutil.WriteFile("proof", []byte("proof"), 0644)
}
