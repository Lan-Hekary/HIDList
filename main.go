package main

import "fmt"
import "runtime"
import "github.com/karalabe/hid"

func main() {
    var age int
    
    fmt.Println("OS: %s", runtime.GOOS)
    fmt.Println("Architecture: %s", runtime.GOARCH)
    var usbs = hid.Enumerate(0,0)

    for _, usb := range usbs {
        fmt.Println(usb.Path)
    }
    fmt.Println("What is your age?")
    fmt.Scan(&age)
    fmt.Println("Your  age ", age)
}