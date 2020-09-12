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
	trimmed := strings.TrimLeft(splitted_command[0], " ")
	if strings.ContainsAny(trimmed,"-path") {
		trimmed_path := splitted_command[1]
		if ContainsQuotes(trimmed_path){
			trimmed_path = DeleteQuotes(trimmed_path)
		}
		file_name := trimmed_path
		deleteFile(file_name)
	} else {
		if(strings.HasPrefix(trimmed,"#")){
			fmt.Println(trimmed)
		}else{
			fmt.Println(trimmed, "command unknow")
		}
		
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
				err := os.MkdirAll(spplited_command[1], 0777)
				if(err!=nil){
					fmt.Println("Error creating route")
					fmt.Println(err)
				}
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
			if(strings.HasPrefix(trimmed,"#")){
				fmt.Println(trimmed)
			}
			if strings.ToLower(trimmed) != "mkdisk" && trimmed!=""{
				fmt.Println(trimmed, "command unknow")
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

func PrintMBR(m mbr){
	if(m.Size != 0){
		fmt.Println("Tamanio", m.Size)
		fmt.Println("Fecha de creacion", GetString(m.Time[:]))
		fmt.Println("Disk signature", m.Disk_signature)
		for _,e := range m.Partitions{
			if(e.Status!='0' && e.Type!='0'){
				fmt.Println("----------------------------------------")
				fmt.Println("Name", GetString(e.Name[:]))
				fmt.Println("Status", string(e.Status))
				fmt.Println("Type", string(e.Type))
				fmt.Println("Fit", GetString(e.Fit[:]))
				fmt.Println("Start", strconv.Itoa(int(e.Start)))
				fmt.Println("Size", strconv.Itoa(int(e.Size)))
			}
		}	
	}else{
		fmt.Println("Empty disk")
	}
}