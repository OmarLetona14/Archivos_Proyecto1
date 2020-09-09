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
	if(getPartition(mkdir_command.Id)){
		pth := mkdir_command.Path
		sb := ReadSB(current_partition.Path, current_partition.Init)
		if ContainsQuotes(mkdir_command.Path) {
			pth = DeleteQuotes(mkdir_command.Path)
		}
		r := ReadAVD(current_partition.Path, sb.Inp_directory_tree)
		err,result := AddDirectory(pth, mkdir_command.P, r,current_partition)
		if(err==nil){
			createTreeReport(result,current_partition)
			fmt.Println(Content)
		}else{
			fmt.Println("********** AN ERROR OCCURRED WHEN TRYING TO CREATE DIRECTORY **********")
		}
	}else{
		fmt.Println("Specificated partition doesnt exist")
	}

}

func getPartition(identifier string)bool{
	for _, e := range Partitions_m{
		if(CompareBytes(identifier,  e.Identifier)){
			current_partition = e
			return true
		}
	}
	return false
}
