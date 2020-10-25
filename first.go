package main

import (
        "os"
        "fmt"
        cn "github.com/artex2000/first/console"
       )

func main() {
    con, err := cn.NewConsole()
    if err != nil {
        fmt.Printf("Error %v getting console\n", err)
        os.Exit(1)
    }

    con.Close()
}

