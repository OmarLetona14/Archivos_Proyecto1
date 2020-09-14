package functions

import(
	"unsafe"
	"fmt"
	"strings"
)

var c_p Mounted_partition
var s_boot Super_Boot

func SeparatePath_Name(path_name string)(path, name string){
	separate := strings.Split(path_name,"/")
	for i,e := range separate{
		if(e!=""){
			if(i!=(len(separate)-1)){
				path += e
			}else{
				name = e
			}
		}
	}
	return 
}

func CreateSystemFile(command Mkfile_command, p_param bool, sb Super_Boot, partiti Mounted_partition){
	c_p := partiti
	if(c_p.Path!=""){
		path, name := SeparatePath_Name(command.Path)
		s_boot := ReadSB(c_p.Path, c_p.Init)
		root_avd := ReadAVD(c_p.Path, s_boot.Inp_directory_tree)
		current_avd:=GetDirectory(p_param, path, root_avd, c_p)
		current_detail := ReadDD(c_p.Path, current_avd.Directory_detail)
		index,overflowed := GetFreeDDIndex(current_detail)
		if(index==0 && overflowed){
			d := CreateDD(c_p)
			current_detail.Next_pointer = d.Id
			current_detail = d
		}
		current_detail.Table[index] = CreateNodeRegister(name, command)
		ModifyDD(c_p.Path, current_detail, current_detail.Id, int64(c_p.Dsk.Size))
	}else{
		fmt.Println("Partition invalid")
	}
	
}

func GetFreeDDIndex(dd_struct dd)(int,bool){
	for i, e := range dd_struct.Table{
		if(e.File_name[0]=='0'){
			return i, false
		}else if(i==4){
			return 0,true
		}
	}
	return 0,false
}

func CreateDD(p Mounted_partition) dd{
	new_dd := dd{}
	dd_size := unsafe.Sizeof(new_dd)
	sb_get := ReadSB(p.Path, p.Init)
	init :=CountBitsDD(sb_get)*int64(dd_size) + sb_get.Inp_directory_detail
	new_dd.Id = init
	WriteDD(p.Path, new_dd, new_dd.Id, int64(p.Dsk.Size))
	dd_created := ReadDD(p.Path, init)
	return dd_created
}

func CreateNodeRegister(name string,c Mkfile_command)inode_register{
	nodeR := inode_register{}
	copy(nodeR.File_name[:],name)
	nodeR.Creation_date = GetCurrentTime()
	nodeR.Modification_date = GetCurrentTime()
	//nodeR.Inode_pointer = CreateInode(c.Size, c.Cont, c).Inode_number
	return nodeR
}

func CreateInode(s int64, content string, comm Mkfile_command)inode{
	n_new := inode{}
	node_size := unsafe.Sizeof(n_new)
	init := CountBitsInode(s_boot)*int64(node_size) + s_boot.Inp_inode_table + s
	n_new.Inode_number = init
	n_new.File_size = s
	n_new.Block_count = 0
	GenerateContent(&n_new, content, comm)
	WriteInode(c_p.Path, n_new, n_new.Inode_number, int64(c_p.Dsk.Size))
	c_inode := ReadInode(c_p.Path, n_new.Inode_number,)
	return c_inode
}

func GenerateContent(n* inode , cont string, com Mkfile_command){
	var bytes_cont[25] byte
	var content_size int64= int64(len(cont))
	if(content_size<=25){
		copy(bytes_cont[:], cont)
		x, overflowed:= GetBlockFreeIndex(n)
		if(overflowed){
			n.Indirect_pointer = CreateInode(com.Size, cont, com).Inode_number
		}else{
			n.Block_array[x] = CreateBlock(bytes_cont)
			n.Block_count += 1
		}
	}else{
		var conter int64 = 0
		var acc string = ""
		for conter<content_size{
			var cont_bytes [25] byte
			var res string = ""
			conter, cont_bytes,res = GetCont(cont, content_size, conter)
			if(cont_bytes[0]!='0'){
				x, overflowed:=GetBlockFreeIndex(n)
				if(!overflowed){
					acc+=res
					n.Block_array[x] = CreateBlock(cont_bytes)
					n.Block_count += 1
				}else{
					n.Indirect_pointer = CreateInode(com.Size, strings.TrimLeft(cont, res),com).Inode_number
				}
			}else{
				return
			}
		}
	}
}

func GetCont(con string, content_size int64, string_pointer int64)(int64, [25] byte, string){
	counter := 0
	accum:=""
	var byts [25] byte
	for content_size>string_pointer{
		if(counter != 25){
			if(con[string_pointer]!='0'){
				byts[counter] = con[string_pointer]
				accum += string(con[string_pointer])
				counter+=1
				string_pointer+=1
			}
		}else{
			counter = 0
			return string_pointer, byts, accum
		}
	}
	return string_pointer, byts, ""
}

func CreateBlock(cont [25] byte)block{
	b := block{}
	block_size := unsafe.Sizeof(b)
	init := CountBitsBlock(s_boot)*int64(block_size) + s_boot.Inp_block
	b.Id = init
	b.Data = cont
	WriteBlock(c_p.Path,b, b.Id, int64(c_p.Dsk.Size))
	r_block := ReadBlock(c_p.Path, init)
	return r_block
}

func GetBlockFreeIndex(node * inode)(int,bool){
	for i, e := range node.Block_array{
		if(e.Data[0]=='0'){
			return i, false
		}else if(i==3){
			return 0,true
		}
	}
	return 0,false
}

