package main

import (
	"fmt"
	"MIA/Project/functions"
	"strings"
	"os"
	"bufio"
	"log"
)

var equalizer string = "->"
var m_command string = ""

func main(){
	fmt.Println("Welcome to the console! (Press x to finish)")
	reader := bufio.NewReader(os.Stdin)
	finish_app := false
	for !finish_app {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = functions.Get_text(input)
		if input != "x"{
			if !strings.HasPrefix(input, "#") { 
				execute_console(input)
			}
		} else {
			fmt.Println("Finishing the app...")
			finish_app = true
		}
	}
}

func execute_console(i string) {
	if(!strings.HasSuffix(i,"/*")){
		m_command += functions.Get_text(i)
		recognize_command(functions.Splitter(m_command))
		m_command = ""
	}else{
		m_command += strings.TrimRight(i, "/*")
	}
}



func recognize_command(commands []string) {
	switch strings.ToLower(commands[0]) {
	case "mkdisk":
		functions.Exec_mkdisk(commands)
	case "exec":
		sub_command := strings.Split(commands[1], equalizer)
		if strings.ToLower(sub_command[0]) == "-path" {
			ReadFile(sub_command[1])		
		} else {
			fmt.Println("Not supported command! ")
			fmt.Println("Maybe you meant -path")
		}
	case "rmdisk":
		functions.Exec_mrdisk(commands)
	case "fdisk":
		functions.Exec_fdisk(commands)
	case "pause":
		fmt.Print("\nPress any key to continue... ")
		reader := bufio.NewReader(os.Stdin)
		x, _ := reader.ReadString('\n')
		x += ""
	case "unmount":
		functions.Exec_unmount(commands)
	case "mkfs":

	case "mount":
		if(len(commands)>=2){
			functions.Exec_mount(commands)
		}else{
			fmt.Println("MOUNTED PARTITIONS")
			fmt.Println("-----------------------------------")
			functions.PrintMount()
		}
	default:
		fmt.Println("Not supported command! ")
	}
}

func ReadFile(file_name string) {
	m_command = ""
	f, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if(scanner.Text()!= " "){
			if !strings.HasPrefix(scanner.Text(), "#"){
				fmt.Println("Executing ", scanner.Text(), "... ")
				execute_console(strings.TrimRight(scanner.Text(), " "))
			}else{
				fmt.Println(scanner.Text())
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}