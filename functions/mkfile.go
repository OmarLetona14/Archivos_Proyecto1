package functions

import(
	"strings"
	"fmt"
	"strconv"
)

var partition Mounted_partition

func createFile(){


}

func Exec_mkfile(com [] string){
	var mkfile_command Mkfile_command
	for _, e := range com{
		spplited_command := strings.Split(e, Equalizer)
		switch strings.ToLower(spplited_command[0]){
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
		default:
			fmt.Println("Command unknow")
		}
	}
	if(mkfile_command.P){
		addDirectory()
	}
	
}

func CreateDir(){

}