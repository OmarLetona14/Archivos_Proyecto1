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
var line string = ""
var mline bool =false

func main(){
	fmt.Println("Welcome to the console! (Press x to finish)")
	reader := bufio.NewReader(os.Stdin)
	finish_app := false
	for !finish_app {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = functions.Get_text(input)
		if strings.ToLower(input) != "x"{
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
	if(!strings.HasSuffix(i,"\\*")){
		m_command += functions.Get_text(i)
		recognize_command(functions.Splitter(m_command))
		m_command = ""
	}else{
		m_command += strings.TrimRight(i, "\\*")
	}
}

func recognize_command(commands []string) {
	var trimmed string = ""
	if(strings.ContainsAny(commands[0], " ")){
		leftTrim := strings.TrimRight(commands[0], " ")
		trimmed = strings.TrimLeft(leftTrim, " ")
	}else{
		trimmed = commands[0]
	}
	switch strings.ToLower(trimmed) {
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
		functions.Exec_mkfs(commands)
	case "rep":
		functions.Exec_rep(commands)
	case "disk":
		//PrintMBR(ReadMBR())
	case "mount":
		if(len(commands)>=2){
			functions.Exec_mount(commands)
		}else{
			fmt.Println("MOUNTED PARTITIONS")
			fmt.Println("-----------------------------------")
			functions.PrintMount()
		}
	case "mkdir":
		functions.Exec_mkdir(commands)
	default:
		if(strings.ToLower(commands[0])!=""){
			fmt.Println("Not supported command! ")
		}
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
			trimmed := functions.GetString([]byte(scanner.Text()))
			if !strings.HasPrefix(trimmed, "#"){
				if !strings.HasSuffix(trimmed, "\\*"){
					if !mline {
						fmt.Println("Executing ", trimmed, "... ")
						execute_console(strings.TrimRight(trimmed, " "))
					}else{
						line += trimmed
						fmt.Println("Executing ", line, "... ")
						execute_console(strings.TrimRight(line, " "))
						mline = false
					}
				}else{
					deleted := strings.TrimRight(trimmed, "\\*")
					line += deleted
					mline = true
				}
			}else{
				fmt.Println(scanner.Text())
			}
		}else{
			fmt.Println("Empty file")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}