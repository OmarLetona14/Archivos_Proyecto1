package functions

import(
	"fmt"
	"strings"
	"strconv"
	"unsafe"
	"os"
	//"bytes"
	"bufio"
)

func Exec_fdisk(com []string) {
	var new_partition Mfdisk_command
	for _, element := range com {
		spplited_command := strings.Split(element, Equalizer)
		trimmed := strings.TrimLeft(spplited_command[0], " ")
		switch strings.ToLower(trimmed) {
		case "-size":
			i, _ := strconv.Atoi(spplited_command[1])
			if i > 0 {
				new_partition.Size = int64(i)
			} else {
				fmt.Println("Partition size must be positive")
				return
			}
		case "-unit":
			new_partition.Unit = spplited_command[1][0]
		case "-path":
			if ContainsQuotes(spplited_command[1]){
				spplited_command[1] = DeleteQuotes(spplited_command[1])
			}
			if _, err := os.Stat(spplited_command[1]); !os.IsNotExist(err) {
				new_partition.Path = spplited_command[1]
			}else{
				fmt.Println("Especificated disk doesnt exist")
				return
			}
		case "-type":
			new_partition.Type =  strings.ToLower(spplited_command[1])[0]
		case "-fit":
			var fit_slice[2] byte
			copy(fit_slice[:], strings.ToLower(spplited_command[1])) 
			new_partition.Fit = fit_slice
		case "-delete":
			new_partition.Delete = true
		case "-name":
			new_partition.Name = spplited_command[1]
		case "-add":
			new_partition.Add = true
		default:
			if(strings.HasPrefix(trimmed,"#")){
				fmt.Println(element)
			}else if trimmed != "fdisk" {
				fmt.Println(trimmed, "command unknow")
			}
		}
	}
	if(new_partition.Path!="" && new_partition.Size!=0 && new_partition.Name!="" || new_partition.Delete){
		PartitionProcess(new_partition)
		//PrintMBR(ReadMBR(new_partition.Path))
	}else{
		fmt.Println("Not enough arguments")
	}
}
func PartitionProcess(cm Mfdisk_command){
	if(cm.Add && cm.Delete){
		fmt.Println("Invalid operation, canont combine add and delete")
	}else if(cm.Add){

	}else if(cm.Delete){
		mbrs:= ReadMBR(cm.Path)
		if verifyPartitionExistence(mbrs, cm.Name) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Are you sure of deleting this partition? [s/n]")
			input, _ := reader.ReadString('\n')
			input = Get_text(input)
			if strings.ToLower(input)=="s"{
				mounted, id := IsMounted(cm.Path, cm.Name)
				if(mounted){
					Unmount(id)
				}
				setDefaultValues(&mbrs, cm.Name)
				ModifyMBR(cm.Path, mbrs) //Sobreescribimos en el archivo binario la nueva tabla mbr
				fmt.Println("Partition deleted sucessfully")
				//PrintMBR(ReadMBR(cm.Path))
			}else{
				return
			}
		}else{
			fmt.Println("Partition doesnt exists on disk")
		}
	}else if(!cm.Delete && !cm.Add){
		mbr_table := ReadMBR(cm.Path) //OBTENEMOS LA TABLA MBR DEL DISCO ESPECIFICADO
		if(verifyDefaultValues(&cm, mbr_table)){
			if(createPart(&mbr_table, cm)){//Modificamos los datos del mbr
				ModifyMBR(cm.Path, mbr_table) //Sobreescribimos en el archivo binario la nueva tabla mbr
				//PrintMBR(ReadMBR(cm.Path))
			} else{
				fmt.Println("Error:Cannot create partition, internal error")
			}
		}else{
			fmt.Println("Error: Cannot create partition, values error")
		}
	}else{
		fmt.Println("Incorrect combination")
	}
}

//VERIFICA SI UNA PARTICION SE ENCUENTRA MONTADA, RECIBE EL PATH A LA PARTICION Y EL NOMBRE
func IsMounted(path string, name string)(bool, [20]string){
	var un [20] string
	var i int
	var exists bool =false 
	for _,element := range Partitions_m {
		if(CompareBytes(path, element.Path)){
			if(CompareBytes(name, element.Name) && i<20){
				exists =true
				un[i] = element.Identifier 
				i+=1
			}
		}
	}
	return exists,un
}

func setDefaultValues(m *mbr, name string){
	for i,e := range m.Partitions{
		n := string(e.Name[:])
		if(CompareBytes(name, n)){
			m.Partitions[i].Status = '0'
			m.Partitions[i].Type = 0
			m.Partitions[i].Fit[0] = 0
			m.Partitions[i].Fit[1] = 0
			m.Partitions[i].Start = 0
			m.Partitions[i].Size =0
			for z:=0;z<len(m.Partitions[i].Name);z++{
				m.Partitions[i].Name[z]=0
			}
		}
	}
}

func verifyPartitionExistence(m mbr, name string)(bool){
	for _,e := range m.Partitions{
		n := string(e.Name[:])
		if CompareBytes(name,n){
			return true
		}
	}
	return false
}

//SE HACEN TODAS LAS VERIFICACIONES ANTES DE CREAR LA PARTICION
func verifyDefaultValues(cm *Mfdisk_command, mbr_table mbr)(Part_error bool){
	if(cm.Fit[0]==0){
		copy(cm.Fit[:],"wf")
		fmt.Println()
	}
	if cm.Unit == 0 {
		cm.Unit = 'k'
	}
	if cm.Type == 0{
		cm.Type = 'p'
	}
	e,p,l := calcPart(mbr_table.Partitions) //OBTENEMOS E= PARTITIONES EXTENDIDAS, F= PARTICIONES LIBRES
	if(!(l+p>4)){ //SE VERIFICA SI EXISTEN PARTICIONES LIBRES
		var type_b byte = cm.Type
		if(type_b=='e'){ 
			if(e==1){//VERIFICAMOS SI EXISTE YA UNA PARTICION EXTENDIDA
				fmt.Println("There one extended partition already")
				return false
			}else{
				return true
			}
		}else if(type_b=='l'){
			if(e==0){
				fmt.Println("You need to create an extended partition before creating a logical partition")
				return false
			}else{
				return true
			}
		}
	}else{
		fmt.Println("Cannot create more partitions, theres 4 already")
		return false
	}
	return true
}

func createPart(mbr_table* mbr, cm Mfdisk_command) (created bool){
	created=false
	part_size := Calc_filesize(string(cm.Unit),int(cm.Size), true)
	i:=0
	for !created && !(i>len(mbr_table.Partitions)){
		if(mbr_table.Partitions[i].Status == '0'){			
			//SE VERIFICAN LOS VALORES DE INICIO
			if(i==0){
				n := mbr{}
				str_mbr := unsafe.Sizeof(n)
				mbr_table.Partitions[i].Start = int64(str_mbr)
				//SE VERIFICA QUE EN CASO DE TENER UNA PARTICION DELANTE DE ESTA HAYA ESPACIO SUFICIENTE
				if(part_size>mbr_table.Partitions[i+1].Start && mbr_table.Partitions[i+1].Start!=0){
					fmt.Println("There is not enough space ")
					created=false
					return
				}
			}else{
				//SI LA PARTICION QUE SE QUIERE CREAR EXCEDE EL TAMANIO DEL DISCO LA FUNCION RETORNARA Y NO SE CREARA LA PARTICION
				//SE CALCULA INICIO DE LA PARTICION ANTERIOR + TAMANIO DE LA PARTICION ANTERIOR + TAMANIO DE LA PARTICION QUE SE QUIERE CREAR
				verifyValue := mbr_table.Partitions[i-1].Start + mbr_table.Partitions[i-1].Size + part_size
				strt :=  mbr_table.Partitions[i-1].Start + mbr_table.Partitions[i-1].Size 
				if(i!=39){
					if(strt+part_size>mbr_table.Partitions[i+1].Start && mbr_table.Partitions[i+1].Start!=0){
						//SE VERIFICA QUE EN CASO DE TENER UNA PARTICION DELANTE DE ESTA HAYA ESPACIO SUFICIENTE
						fmt.Println("There is not enough space ")
						created=false
						return
					}
				}
				if(!(verifyValue>mbr_table.Size)){
					mbr_table.Partitions[i].Start = strt
				}else{
					fmt.Println("There is not enough space ")
					created=false
					return
				}
			}
			//SE LLENAN LOS VALORES DE LA PARTICION CONTENIDA EN EL MBR
			mbr_table.Partitions[i].Status = 'i'
			mbr_table.Partitions[i].Type = cm.Type
			mbr_table.Partitions[i].Fit = cm.Fit
			mbr_table.Partitions[i].Size = part_size
			copy(mbr_table.Partitions[i].Name[:], cm.Name)
			created=true
			return
		}
		i+=1
	}
	return
}

func calcPart(parti [40] Partition)(int, int, int){
	primary := 0
	logical:=0
	extended := 0
	for i:=0;i<len(parti);i++ {
		if(parti[i].Type == 'p'){
			primary += 1
		}else if(parti[i].Type == 'e'){
			extended +=1
		}else if(parti[i].Type == 'l'){
			logical +=1
		}
	}
	return extended, primary, logical
}