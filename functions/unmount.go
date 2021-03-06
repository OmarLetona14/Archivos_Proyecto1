package functions

import (
	"strings"
	"strconv"
	"regexp"
	"fmt"
)


func Exec_unmount(com [] string){
	var unmount_command Unmount_command
	for _,element := range com{
		spplited_command := strings.Split(element, Equalizer)
		trimmed:= strings.TrimLeft(spplited_command[0], " ")
		rex :=`-id\d+`
		mtch,_ := regexp.MatchString(rex,trimmed)
		if mtch{
			getNumbers := strings.Split(trimmed, "-id")
			index,err := strconv.Atoi(getNumbers[1])
			if(err==nil){
				unmount_command.List_id[index] = spplited_command[1]
			}else{
				fmt.Println("An error ocurred when trying to unmount partition")
			}
		}else{
			if(strings.HasPrefix(trimmed,"#")){
				fmt.Println(trimmed)
			}
			if(trimmed!=" " || trimmed!="" ||trimmed!="\n"||
			spplited_command[0]!="\t"){
				fmt.Println("unrecognize command")
			}
		}
	}
	if(!Unmount(unmount_command.List_id)){
		fmt.Println("Especificated id doesnt exists in mounted partitions")
	}
}

func Unmount(ids[20] string)bool{
	mounted:=false
	for z:=0; z<len(ids);z++{
		if(ids[z]!=""){
			for i,element := range Partitions_m{
				if(element.Identifier!=""){
					if(CompareBytes(ids[z],element.Identifier)){
						Partitions_m[i] = Mounted_partition{}
						fmt.Println("Partition", element.Identifier, "unmounted sucessfully")
						mounted=true 
					}
				}
			}
		}
	}
	return mounted
}