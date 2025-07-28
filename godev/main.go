package main

import (
	"os/exec"
    "os"

)
func main() {
 
    Jarname := os.Getenv("Jar") /*
    Ram := os.Getenv("Ram")
    Xmx := os.Getenv("Xmx")
    Xms := os.Getenv("Xms")
*/
    exec.Command("java", "-Xmx4G", "-Xms4G", "-jar", "/minecraft/", Jarname )
 //java -Xmx2G -jar fabric-server-mc.1.21.7-loader.0.16.14-launcher.1.0.3.jar nogui
}

