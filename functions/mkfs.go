package functions

import(
	"strings"
	"fmt"
)

func Exec_mkfs(com [] string){
	var mkfs_command Mkfs_command
	for _, element :=range com{
		spplited_command := strings.Split(element, Equalizer)
		switch strings.ToLower(spplited_command[0]){
		case "-id":
			mkfs_command.Identifier = spplited_command[1]
		case "-type":
			mkfs_command.Type = spplited_command[1]
		case "-add":
			mkfs_command.Add = true
		case "unit":
			mkfs_command.Unit = spplited_command[1][0]
		default:
			fmt.Println("Unrecognize command")
		}
	}
	if(mkfs_command.Identifier!=""){
		verify, mounted_partition := VerifyMountedPartition(mkfs_command.Identifier) 
		if verify {
			FormatPartition(mkfs_command, mounted_partition)
		}else{
			fmt.Println("The specified partition is not found or does not exist")
		}
		
	}else{
		fmt.Println("Not enough arguments")
	}

}

func FormatPartition(mkfs Mkfs_command, m Mounted_partition){
	var new_format = Super_Boot{}
	Format(&new_format, m.Dsk ,int64(m.Size), int64(m.Init))
	disk_path := m.Dsk.Path + m.Dsk.Name 
	WriteSuperB(disk_path,new_format,m.Init)
	//s := ReadSB(disk_path, m.Init)
	//printSB(&s, int64(m.Init))
	//printSB(&new_format, int64(m.Init))
}

