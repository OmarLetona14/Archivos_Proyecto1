package functions

import(
	"strings"
	//"fmt"
	"errors"
	"unsafe"
)

var current_p Mounted_partition
var current_sb Super_Boot
var p_command bool
var current_pointer *avd_binary
var id_avd int
var root avd_binary

func AddDirectory(dir string, p_com bool, current_root avd_binary, p Mounted_partition)(e error, a avd_binary){
	root = current_root
	current_p = p
	GetSB()
	if(root.Creation_date[0]==0){
		new_avd := WriteBinaryAVD("/", "")
		root = new_avd
	}
	err := AllDirectories(dir, p_com)
	if err!=nil {
		return errors.New("*******************"),root
	}
	ModifySB(p.Path, current_sb, p.Init, int64(p.Dsk.Size))
	//s := ReadSB(p.Path, p.Init)
	//printSB(&s, int64(p.Init))
	return nil,root
}

func AllDirectories(dir string, p_com bool)(e error){
	p_command = p_com 
	directories := strings.Split(dir, "/")
	current_pointer = &root
	for _, e := range directories{
		if(e!=" " && e!=""){
			err := createDirectory(e)
			if(err!=nil){
				return errors.New("Couldnt create directory")
			}
		}
	}
	return nil
}

func createDirectory(directory string)(e error){
	comprobate, dir := searchSub(current_pointer, directory)//Comprueba si el subdirectorio existe dentro del directorio actual
	if(comprobate){
		//Si existe cambiamos de puntero
		current_pointer = &dir
	}else{
		//Si no existe lo creamos
		if(p_command){
			//Se crea el nuevo directorio
			c := WriteBinaryAVD(directory,"")
			current_sb.Last_mount_date = GetCurrentTime()
			ModifySB(current_p.Path,current_sb, current_p.Init, int64(current_p.Dsk.Size))
			//Lo agregamos al listado de subdirectorio del directorio en donde nos encontramos
			i,space:= GetFreeIndex(current_pointer) //Comprobamos que haya espacio en el subdictorio, de no ser asi
			//Se crea un nuevo directorio 
			if(space){
				current_pointer.Sub_directory_pointers[i] = c.Id
				//actualizamos el puntero
				result := ReadAVD(current_p.Path, c.Id)
				p_cont := *current_pointer
				ModifyAVD(current_p.Path,p_cont,current_pointer.Id, int64(current_p.Dsk.Size))
				current_pointer = &result
			}else{
				nxt := WriteBinaryAVD(string(current_pointer.Directory_name[:]), "")
				current_pointer.Avd_next = nxt.Id
				sub:=ReadAVD(current_p.Path, current_pointer.Avd_next)
				sub.Sub_directory_pointers[0] = c.Id
				//Actualizamos el puntero
				result := ReadAVD(current_p.Path, c.Id)
				p_cont := *current_pointer
				ModifyAVD(current_p.Path,p_cont,current_pointer.Id, int64(current_p.Dsk.Size))
				current_pointer = &result
			}
		}else{
			//Si el parametro p no esta espeficiado muestra un error
			return errors.New("Error: Directory doesnt exists")
		}
	}
	return nil
}



func GetFreeIndex(dir *avd_binary)(int64, bool){
	for i,e := range dir.Sub_directory_pointers{
		if(e==0){
			return int64(i), true
		}
	}
	return 0,false
}


func searchSub(r *avd_binary, dir_name string)(bool, avd_binary){
	for _,e := range r.Sub_directory_pointers{
		if(e!=0){
			current_avd := ReadAVD(current_p.Path, e)
			e_name := string(current_avd.Directory_name[:])
			if(CompareBytes(dir_name, e_name)){
				return true, current_avd
			}
		}
	}
	if(r.Avd_next!=0){
		avd_nxt := ReadAVD(current_p.Path, r.Avd_next)
		searchSub(&avd_nxt, dir_name)
	}
	return false,avd_binary{}
}

func WriteBinaryAVD(dir_name string, prop string)avd_binary{
	bin := avd_binary{}
	bin_size := unsafe.Sizeof(bin)
	init := CountBits(current_sb)*int64(bin_size) + current_sb.Inp_directory_tree
	bin.Id = init
	copy(bin.Directory_name[:], []byte(dir_name))
	bin.Creation_date = GetCurrentTime()
	bin.Directory_detail = 0
	bin.Avd_next = 0
	//fmt.Println("WRITING ", bin)
	WriteAVD(current_p.Path, bin,bin.Id, int64(current_p.Dsk.Size))
	avd_new := ReadAVD(current_p.Path, bin.Id)
	ModifyBitmap(current_p.Path, current_sb.Ffb_directory_tree, int64(current_p.Dsk.Size))
	current_sb.Ffb_directory_tree += 1
	current_sb.Free_virtual_tree_count -= 1
	current_sb.Last_mount_date = GetCurrentTime()
	ModifySB(current_p.Path, current_sb,current_p.Init, int64(current_p.Dsk.Size))
	//fmt.Println("READING", avd_new)
	return avd_new
}


func GetSB(){
	current_sb = ReadSB(current_p.Path, current_p.Init)
}