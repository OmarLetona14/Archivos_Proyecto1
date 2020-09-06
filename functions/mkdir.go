package functions

import(
	"strings"
	"fmt"
)

var current_partition Mounted_partition

func Exec_mkdir(com [] string){
	var mkdir_command Mkdir_command
	for _, e := range com{
		spplited_command := strings.Split(e, Equalizer)
		switch strings.ToLower(spplited_command[0]){
		case "-id":
			mkdir_command.Id = spplited_command[1]
		case "-path":
			mkdir_command.Path = spplited_command[1]
		case "-p":
			mkdir_command.P = true
		}
	}
	if(GetPartition(mkdir_command.Id)){
		pth := mkdir_command.Path
		if ContainsQuotes(mkdir_command.Path) {
			pth = DeleteQuotes(mkdir_command.Path)
		}
		err := addDirectory(pth, mkdir_command.P)
		if(err==nil){
			fmt.Println("********** DIRECTORY  " , pth, " CREATED CORRECTLY **********")
		}else{
			fmt.Println("********** AN ERROR OCCURRED WHEN TRYING TO CREATE DIRECTORY **********")
		}
	}else{
		fmt.Println("Specificated partition doesnt exist")
	}

}

func GetPartition(identifier string)bool{
	for _, e := range Partitions_m{
		if(CompareBytes(identifier,  e.Identifier)){
			current_partition = e
			return true
		}
	}
	return false
}