package main
import (
    "fmt"
    "os"
    "os/exec"
)





func main() {
    //Declaring launch argument string slice
    var args []string


    Xmx := os.Getenv("Xmx")
    if Xmx == "" {
        Xmx = "4G"
    }

    jarName := os.Getenv("jarName")
    if jarName == "" {
        jarName = "/mc/server.jar"
    } else {
        jarName="/mc/"+jarName
        fmt.Println(jarName)
    }


    args = append(args,"-Xmx" + Xmx) 

    args = append(args, "-jar", jarName)

   args = append(args, "nogui")


    fmt.Println(args)
    //exec.Command("java", "-Xmx4G", "-Xms4G", "-jar", Jarname )
    cmd := exec.Command("java", args...)
    fmt.Println("Exec,")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    fmt.Println("run")

    cmd.Run()

}
