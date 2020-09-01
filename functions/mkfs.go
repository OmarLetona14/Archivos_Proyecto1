package functions

import(
	"strings"
	"fmt"
)

func Exec_mkfs(com [] string){
	var mkfs_command Mkfs_command
	for _, element :=range com{
		spplited_command := strings.Split(element, Equalizer)
		switch strings.ToLower(spplited_command[0]){
		case "-id":
			mkfs_command.Identifier = spplited_command[1]
		case "-type":
			mkfs_command.Type = spplited_command[1]
		case "-add":
			mkfs_command.Add = true
		case "unit":
			mkfs_command.Unit = spplited_command[1][0]
		default:
			fmt.Println("Unrecognize command")
		}
	}
	if(mkfs_command.Identifier!=""){
		
	}else{
		fmt.Println("Not enough arguments")
	}

}

func FormatPartition(mkfs Mkfs_command){

}