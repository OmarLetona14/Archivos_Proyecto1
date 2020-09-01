package functions

import(
	"fmt"
	"strings"
	"strconv"
	"os"
)

func Exec_fdisk(com []string) {
	var new_partition Mfdisk_command
	for _, element := range com {
		spplited_command := strings.Split(element, Equalizer)
		switch strings.ToLower(spplited_command[0]) {
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
			if spplited_command[0] != "fdisk" {
				fmt.Println(spplited_command[0], "command unknow")
			}
		}
	}
	if(new_partition.Path!="" && new_partition.Size!=0 && new_partition.Name!=""){
		PartitionProcess(new_partition)
	}else{
		fmt.Println("Not enough arguments")
	}
}


func PartitionProcess(cm Mfdisk_command){
	if(cm.Add && cm.Delete){
		fmt.Println("Invalid operation, canont combine add and delete")
	}else if(cm.Add && !cm.Delete){

	}else if(cm.Delete && !cm.Add){

	}else if(!cm.Delete && !cm.Add){
		mbr_table := ReadMBR(cm.Path) //OBTENEMOS LA TABLA MBR DEL DISCO ESPECIFICADO
		if(verifyDefaultValues(&cm, mbr_table)){
			if(createPart(&mbr_table, cm)){//Modificamos los datos del mbr
				ModifyMBR(cm.Path, mbr_table) //Sobreescribimos en el archivo binario la nueva tabla mbr
				PrintMBR(ReadMBR(cm.Path))
			} 
		}
	}else{
		fmt.Println("Incorrect combination")
	}
}


//SE HACEN TODAS LAS VERIFICACIONES ANTES DE CREAR LA PARTICION
func verifyDefaultValues(cm *Mfdisk_command, mbr_table mbr)(Part_error bool){
	if(cm.Fit[0]==0){
		copy(cm.Fit[:],"wf")
	}
	if cm.Unit == 0 {
		cm.Unit = 'k'
	}
	if cm.Type == 0{
		cm.Type = 'p'
	}
	e,_,f := calcPart(mbr_table.Partitions) //OBTENEMOS E= PARTITIONES EXTENDIDAS, F= PARTICIONES LIBRES
	if(f>=1){ //SE VERIFICA SI EXISTEN PARTICIONES LIBRES
		if(cm.Type=='e'){ 
			if(e==1){//VERIFICAMOS SI EXISTE YA UNA PARTICION EXTENDIDA
				fmt.Println("There one extended partition already")
			}
			return false
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
	for !created && !(i>=len(mbr_table.Partitions)){
		if(mbr_table.Partitions[i].Status == '0'){
			//SE VERIFICAN LOS VALORES DE INICIO
			if(i==0){
				mbr_table.Partitions[i].Start = 0
			}else{
				//SI LA PARTICION QUE SE QUIERE CREAR EXCEDE EL TAMANIO DEL DISCO LA FUNCION RETORNARA Y NO SE CREARA LA PARTICION
				//SE CALCULA INICIO DE LA PARTICION ANTERIOR + TAMANIO DE LA PARTICION ANTERIOR + TAMANIO DE LA PARTICION QUE SE QUIERE CREAR
				verifyValue := mbr_table.Partitions[i-1].Start + mbr_table.Partitions[i-1].Size + part_size
				strt :=  mbr_table.Partitions[i-1].Start + mbr_table.Partitions[i-1].Size 
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

func calcPart(parti [4] Partition)(int, int, int){
	primary := 0
	free:=0
	extended := 0
	for i:=0;i<len(parti);i++ {
		if(parti[i].Type == 'p'){
			primary += 1
		}else if(parti[i].Type == 'e'){
			extended +=1
		}else{
			free +=1
		}
	}
	return extended, primary, free
}