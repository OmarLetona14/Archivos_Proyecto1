package functions

import(
	"strings"
	"fmt"
)

var current_p Mounted_partition
var current_sb Super_Boot
var p_command bool
var current_pointer *avd 

func AllDirectories(dir string, part Mounted_partition, p_com bool){
	current_p = part
	GetSB()
	p_command = p_com
	directories := strings.Split(dir, "/")
	for _, e := range directories{
		if(e!=""){
			root := ReadAVD(current_p.Path + current_p.Name, current_p.Init)
			if(root.Creation_date[0]==0){//Comprueba si existe la carpeta root
				//Se crea la carpeta root
				WriteRoot(current_p.Path +current_p.Name, createAVD("/", ""), current_sb.Inp_directory_tree)
				current_pointer = &root
			}else{
				createDirectory(e)	
			}
		}
	}
}

func createDirectory(directory string){
	comprobate, dir := searchSub(current_pointer, directory)//Comprueba si el subdirectorio existe dentro del directorio actual
	if(comprobate){
		//Si existe cambiamos de puntero
		current_pointer = dir
	}else{
		//Si no existe lo creamos
		if(p_command){
			//Se crea el nuevo directorio
			c := createAVD(directory,"")
			//Lo agregamos al listado de subdirectorio del directorio en donde nos encontramos
			i,space:= GetFreeIndex(current_pointer) //Comprobamos que haya espacio en el subdictorio, de no ser asi
			//Se crea un nuevo directorio 
			if(space){
				current_pointer.Sub_directory_pointers[i] = &c
				//actualizamos el puntero
				current_pointer = current_pointer.Sub_directory_pointers[i]
			}else{
				nxt := createAVD(string(current_pointer.Directory_name[:]), "")
				current_pointer.Avd_next = &nxt
				current_pointer.Avd_next.Sub_directory_pointers[0] = &c
				//Actualizamos el puntero
				current_pointer = current_pointer.Avd_next.Sub_directory_pointers[0]
			}
				
		}else{
			//Si el parametro p no esta espeficiado muestra un error
			fmt.Println("***** FATAL *****")
			fmt.Println("Error: Directory doesnt exists")
			return 
		}
	}
}

func GetFreeIndex(dir *avd)(int64, bool){
	for i,e := range dir.Sub_directory_pointers{
		if(e==nil){
			return int64(i), true
		}
	}
	return 0,false
}

func searchSub(r* avd, dir_name string)(bool, *avd){
	for _,e := range r.Sub_directory_pointers{
		if(e!=nil){
			e_name := string(e.Directory_name[:])
			if(CompareBytes(dir_name, e_name)){
				return true, e
			}
		}
	}
	if(r.Avd_next!=nil){
		searchSub(r.Avd_next, dir_name)
	}
	return false,nil
}

func createAVD(directory_name string, proper string) avd{
	v:=avd{}
	var pointers [6] *avd
	var ddetail dd
	v.Creation_date = GetCurrentTime()
	copy(v.Directory_name[:], directory_name)
	v.Sub_directory_pointers = pointers
	v.Directory_detail = ddetail
	copy(v.Proper[:], proper)
	return v
}

func GetSB(){
	current_sb = ReadSB(current_p.Path+ current_p.Name, current_p.Init)
}