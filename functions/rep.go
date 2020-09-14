package functions

import (
	"strings"
	"os"
	"fmt"
)

var typ string = ""

func Exec_rep(com []string) {
	var rep repo_command
	for _, e := range com{
		spplited_command := strings.Split(e, Equalizer)
		sp := strings.TrimLeft(spplited_command[0], " ")
		switch (strings.ToLower(sp)){
		case "-name":
			rep.Name = spplited_command[1]
		case "-path":
			if ContainsQuotes(spplited_command[1]){
				spplited_command[1] = DeleteQuotes(spplited_command[1])
			}
			ph, _ := ReturnName(spplited_command[1])
			if _, err := os.Stat(ph); os.IsNotExist(err) {
				os.MkdirAll(ph, os.ModePerm)
			}
			rep.Path = spplited_command[1]
		case "-id":
			rep.Id = spplited_command[1]
		case "-route":
			rep.Route = spplited_command[1]
		}
	}
	if(rep.Path!="" && rep.Name !="" && rep.Id!=""){
		Report(rep)
	}else{
		fmt.Println("Too few arguments")
	}
}

func Report(r repo_command){
	verify, mounted_partition := VerifyMountedPartition(r.Id)
	if(verify){
		switch strings.ToLower(r.Name) {
		case "mbr":
			mb:=ReadMBR(mounted_partition.Path)
			createMbrReport(mb, r.Path)
		case "disk":
			mb:=ReadMBR(mounted_partition.Path)
			createDiskReport(mb,mounted_partition ,r.Path)
		case "sb":
			sb := ReadSB(mounted_partition.Path, mounted_partition.Init)
			createSbReport(sb,mounted_partition ,r.Path)
		case "bm_arbdir":
		case "bm_detdir":
		case "bm_inode":
		case "bm_block":
		case "bitacora":
		case "directorio":
			sb := ReadSB(mounted_partition.Path, mounted_partition.Init)
			avd := ReadAVD(mounted_partition.Path,sb.Inp_directory_tree)
			if(avd.Id!=0){
				createTreeReport(avd,mounted_partition,r.Path)
			}else{
				fmt.Println("Error creating report")
			}
		case "tree_file":
		case "tree_directorio":
		case "tree_complete":
		case "ls":
		default:
			fmt.Println("Incorrect params")
		}
	}else{
		fmt.Println("Specificated partition doesnt exists")
	}
}