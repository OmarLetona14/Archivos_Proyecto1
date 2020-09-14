package functions

import(
	"strings"
	"fmt"
	"strconv"
)

var partition Mounted_partition


func Exec_mkfile(com [] string){
	var mkfile_command Mkfile_command
	for _, element := range com{
		spplited_command := strings.Split(element, Equalizer)
		trimmed := strings.TrimLeft(spplited_command[0], " ")
		switch strings.ToLower(trimmed){
		case "-id":
			mkfile_command.Id = spplited_command[1]
		case "-path":
			if(ContainsQuotes(spplited_command[1])){
				spplited_command[1] = DeleteQuotes(spplited_command[1])
			}
			mkfile_command.Path = spplited_command[1]
		case "-p":
			mkfile_command.P = true
		case "-size":
			conv, err := strconv.Atoi(spplited_command[1])
			if err==nil {
				if(conv>=0){
					mkfile_command.Size = int64(conv)
				}else{
					fmt.Println("Size must be positive")
				}
			}else{
				fmt.Println("Error on size parameter!")
			}
		case "-cont":
			mkfile_command.Cont = spplited_command[1]
		default:
			if(strings.HasPrefix(trimmed,"#")){
				fmt.Println(element)
			}
			if(strings.ToLower(trimmed)!="mkfile" && trimmed!="" && trimmed!=" "){
				fmt.Println("Command unknow",trimmed)
			}
		}
	}
	//sb := ReadSB(partition.Path, partition.Init)
	ver, part := VerifyMountedPartition(mkfile_command.Id)
	if(ver){
		partition = part
		CreateFile(mkfile_command,part)
	}else{
		fmt.Println("Partition doesnt exists")
	}
}

func CreateFile(commad Mkfile_command, p Mounted_partition){
	sb_b := ReadSB(p.Path, p.Init)
	CreateSystemFile(commad,commad.P, sb_b, p)
}
