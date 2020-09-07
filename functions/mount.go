package functions

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func Exec_mount(com [] string){
	var current_mounted_disk Mounted_disk
	var new_mount Mount_command
	var unmounted_partition Partition
	for _,element := range com {
		spplited_command := strings.Split(element, Equalizer)
		switch  strings.ToLower(spplited_command[0]) {
		case "-path":
			if ContainsQuotes(spplited_command[1]){
				spplited_command[1] = DeleteQuotes(spplited_command[1])
			}
			if _, err := os.Stat(spplited_command[1]); !os.IsNotExist(err) {
					new_mount.Path = spplited_command[1]
			}else{
				fmt.Println("Especificated disk doesnt exist")
				return
			}
		case "-name":
			//Primero verificamos si la particion no esta montada
			v, _ := VerifyMountedPartitionByName(spplited_command[1])
			if(!v){
				//Obtenemos el MBR del disco especificado
				dsik := ReadMBR(new_mount.Path)
				unmounted_partition =GetPartitionByName(dsik,spplited_command[1])
				if(!verifyMountedDisk(GetPath(new_mount.Path))){
					mount_disk(GetPath(new_mount.Path))
				}
				current_mounted_disk = GetDiskByPath(new_mount.Path)
				for _,element := range dsik.Partitions {
					name_dsk := strings.TrimRight(string(element.Name[:])," ") 
					if(CompareBytes(spplited_command[1],name_dsk)){
						new_mount.Name = spplited_command[1]
					}
				}
				if(new_mount.Name == ""){
					fmt.Println("Partition doesnt exists in disk")
				}
			}else{
				fmt.Println("This partition is already mounted")
			}
		}
	}
	if(new_mount.Path != "" && new_mount.Name != ""){
		var mounted Mounted_partition
		if(!ContainsQuotes( new_mount.Path)){
			mounted.Path = new_mount.Path
		}else{
			mounted.Path = DeleteQuotes(new_mount.Path)
		}
		mounted.Name = new_mount.Name
		mounted.Identifier = GetMountIdentifier(new_mount.Path)
		if(current_mounted_disk.Identifier!=""){
			mounted.Dsk = current_mounted_disk
		}
		mounted.Init = unmounted_partition.Start
		mounted.Size = unmounted_partition.Size
		fmt.Println("********** PARTITION ", mounted.Identifier, "MOUNTED **********")
		Partitions_m[Partitions_size] = mounted
		Partitions_size += 1
		
	}else{
		fmt.Println("Too few arguments")
	}
}

func PrintMount(){
	for _,element := range Partitions_m{
		if(element.Identifier!=""){
			fmt.Println("IDENTIFIER:", element.Identifier)
			fmt.Println("DISK:", element.Path)
			fmt.Println("PARTITION", element.Name)
			fmt.Println("----------------------------------")
		}
	}
}

func GetMountIdentifier(path string)string{
	for i, element := range Disks_m{
		absolute_path := element.Path + element.Name
		if(absolute_path==path){
			id := "vd" + element.Identifier + strconv.Itoa(element.Count_mounted)
			element.Count_mounted += 1
			Disks_m[i] = element
			return id
		}
	}
	return ""
}

func mount_disk(path string, name string){
	m := ReadMBR(path + name)
	var dsk Mounted_disk 
	dsk.Identifier = GetIdentifier(Disks_size)
	dsk.Size = int(m.Size)
	dsk.Path = path
	dsk.Name = name
	Disks_m[Disks_size] = dsk
	Disks_size += 1
}

func verifyMountedDisk(path string, name string)bool{
	for _,element := range Disks_m{
		if(element.Name!=""){
			path_abs := element.Path + element.Name
			if(CompareBytes(path+name, path_abs)){
				return true
			}
		}
	}
	return false
}

func GetIdentifier(elements int) string {
	return string(97+elements)
}

func VerifyMountedPartitionByName(name string)(bool, Mounted_partition){
	for _,element := range Partitions_m{
		if(CompareBytes(name, element.Name)){
			fmt.Println("Partition",name, "Is already mounted", element.Name)
			return true, element
		}
	}
	return false, Mounted_partition{}
}


func VerifyMountedPartition(id string)(bool, Mounted_partition){
	for _,element := range Partitions_m{
		if(CompareBytes(id, element.Identifier)){
			return true, element
		}
	}
	return false, Mounted_partition{}
}

func GetDiskByPath(path string) Mounted_disk{
	for _, element := range Disks_m{
		if(CompareBytes(path, element.Path)){
			return element
		}
	}
	return Mounted_disk{}
}