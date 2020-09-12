package functions

import(
	"strings"
	"fmt"
)

func Exec_mkfs(com [] string){
	var mkfs_command Mkfs_command
	for _, element :=range com{
		spplited_command := strings.Split(element, Equalizer)
		trimmed:= strings.TrimLeft(spplited_command[0], " ")
		switch strings.ToLower(trimmed){
		case "-id":
			mkfs_command.Identifier = spplited_command[1]
		case "-type":
			mkfs_command.Type = spplited_command[1]
		case "-add":
			mkfs_command.Add = true
		case "unit":
			mkfs_command.Unit = spplited_command[1][0]
		default:
			if(strings.HasPrefix(trimmed,"#")){
				fmt.Println(element)
			}else{
				if trimmed!="mkfs" && trimmed!=""{
					fmt.Println("Unrecognize command")
				} 
			}
		}
	}
	if(mkfs_command.Identifier!=""){
		verify, mounted_partition := VerifyMountedPartition(mkfs_command.Identifier) 
		if verify {
			FormatPartition(mkfs_command, mounted_partition)
			mounted_partition.Formatted = true
			fmt.Println("********** PARTITION FORMATTED CORRECTLY **********")
		}else{
			fmt.Println("The specified partition is not found or does not exist")
		}	
	}else{
		fmt.Println("Not enough arguments")
	}
}

func FormatPartition(mkfs Mkfs_command, m Mounted_partition){
	new_format := Super_Boot{}
	Format(&new_format, m.Dsk ,int64(m.Size), int64(m.Init))
	disk_path := m.Dsk.Path + m.Dsk.Name 
	WriteSuperB(disk_path,new_format,m.Init, int64(m.Dsk.Size))
}

