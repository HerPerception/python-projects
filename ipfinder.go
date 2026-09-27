package main

import (
        "fmt"
        "net"
        "strings"
)

func main(){
        fmt.Println("Input a website to find the IP address:")
        var input string
        fmt.Scanln(&input)
        input = strings.TrimPrefix(input, "https://")
        addr, err := net.LookupHost(input)
        if err != nil {
                fmt.Println(err)
                fmt.Println("Problem finding IP address. Check URL and try again")
                return
        }
        for index, ip := range addr {
                fmt.Printf("IP address %d: %s\n", index+1, ip)
        }

}
