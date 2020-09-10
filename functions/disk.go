package functions

import(
	"strings"
	"fmt"
	"os"
	"strconv"
)

var Equalizer string = "->"


func Exec_mrdisk(com []string) {
	splitted_command := strings.Split(com[1], Equalizer)
	if splitted_command[0] == "-path" {
		file_name := splitted_command[1]
		deleteFile(file_name)
	} else {
		fmt.Println(splitted_command[0], "command unknow")
	}
}

func deleteFile(path string){
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := os.Remove(path)
		if err != nil {
			fmt.Println(err)
			return
		}	
		fmt.Println("Removed successfully!")
	}else{
		fmt.Println("Error: File doesnt exists!")
	}
}

func Exec_mkdisk(com []string)  {
	var new_disk Mkdisk_command
	for _, element := range com {
		spplited_command := strings.Split(element, Equalizer)
		trimmed := strings.TrimLeft(spplited_command[0], " ")
		switch strings.ToLower(trimmed) {
		case "-size":
			i, err := strconv.Atoi(spplited_command[1])
			if i > 0 && err==nil{
				new_disk.Size = i
			} else {
				fmt.Println("Size must be positive! ")
			}
		case "-path":
			if ContainsQuotes(spplited_command[1]){
				spplited_command[1] = DeleteQuotes(spplited_command[1])
			}
			if _, err := os.Stat(spplited_command[1]); os.IsNotExist(err) {
				os.MkdirAll(spplited_command[1], os.ModePerm)
			}
			new_disk.Path = spplited_command[1]
		case "-name":
			if strings.HasSuffix(spplited_command[1], ".dsk") {
				new_disk.Name = spplited_command[1]
			} else {
				fmt.Println("Error! Name must have .dsk extension")
				return 
			}
		case "-unit":
			new_disk.Unit = spplited_command[1]
		default:
			if spplited_command[0] != "mkdisk" && spplited_command[0]!=""{
				fmt.Println(spplited_command[0], "command unknow")
			}
		}
	}
	if(new_disk.Path!="" && new_disk.Size != 0 && new_disk.Name!=""){
		WriteFile(new_disk.Path+new_disk.Name, Calc_filesize(new_disk.Unit, new_disk.Size,false)) 
	}else{
		fmt.Println("Too few arguments")
	}
}

func GetPartitionByName(m mbr, name string) Partition{
	for _,e := range m.Partitions{
		nm := string(e.Name[:])
		if(CompareBytes(name, nm)){
			return e
		}
	}
	return Partition{}
}