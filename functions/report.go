package functions

import(
	"strconv"
	"os"
	"bufio"
	"log"
	"unsafe"
	"fmt"
)

var Content string
var curr_partition Mounted_partition

func createDotFile(path string){
	file, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(Content)
        if err != nil {
            log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	}
	writer.Flush()
	Content=""
}

func createTreeReport(r avd_binary,p Mounted_partition, path string){
	curr_partition = p
	if(r.Creation_date[0]!=0){
		Content += "digraph G{ \n"
		GetContent(r)
		Content+="}"
	}else{
		fmt.Print("Empty tree")
	}
	pth := GetPathWODot(path)
	createDotFile(pth + ".dot")
	execDot(pth + ".dot", pth+ ".png")
}

func createSbReport(r Super_Boot,p Mounted_partition, path string){
	curr_partition = p
	if(r.Creation_date[0]!=0){
		Content += "digraph G{ \n"
		SbReport(r)
		Content+="}"
	}else{
		fmt.Print("Empty table")
	}
	pth := GetPathWODot(path)
	createDotFile(pth + ".dot")
	execDot(pth + ".dot", pth+ ".png")
}

func createDiskReport(r mbr,p Mounted_partition, path string){
	curr_partition = p
	if(r.Time[0]!=0){
		Content += "digraph G{ \n"
		diskReport(r)
		Content+="}"
	}else{
		fmt.Print("Empty table")
	}
	pth := GetPathWODot(path)
	createDotFile(pth + ".dot")
	execDot(pth + ".dot", pth+ ".png")
}

func createMbrReport(r mbr, path string){
	if(r.Time[0]!=0){
		Content += "digraph G{ \n"
		MbrReport(r)
		Content+="}"
	}else{
		fmt.Print("Empty table")
	}
	pth := GetPathWODot(path)
	createDotFile(pth + ".dot")
	execDot(pth + ".dot", pth+ ".png")
}

func GetContent(r avd_binary){
	str:=GetString(r.Directory_name[:])
	Content += "c" + strconv.Itoa(int(r.Id)) +" [shape =record label=\"{" + str + " | {"
	for i,e := range r.Sub_directory_pointers {
		if(i!=0){
			if(e!=0){
				current_avd := ReadAVD(curr_partition.Path, e)
				Content += "| <f" + strconv.Itoa(i) + "> " + strconv.Itoa(int(current_avd.Id))
			}else{
				Content += "| <f" + strconv.Itoa(i) + "> -1"
			}
		}else{
			if(e!=0){
				current_avd := ReadAVD(curr_partition.Path, e)
				Content += " <f" + strconv.Itoa(i) + "> " + strconv.Itoa(int(current_avd.Id))
			}else{
				Content += " <f" + strconv.Itoa(i) + "> -1"
			}
		}
	}
	if(r.Avd_next==0){
		Content += "| <f6> -1"
	}else{

	}
	if(r.Directory_detail==0){
		Content += "| <f7> -1"
	}else{

	}
	Content += "}}\"]\n"
	for i,e := range r.Sub_directory_pointers{
		if(e!=0){
			current_avd := ReadAVD(curr_partition.Path, e)
			Content += "c" + strconv.Itoa(int(r.Id)) + ":f" + strconv.Itoa(i) + " -> c" + strconv.Itoa(int(current_avd.Id)) +"\n"
		}
	}
	for _,e :=range r.Sub_directory_pointers{
		if(e!=0){
			current_avd := ReadAVD(curr_partition.Path, e)
			GetContent(current_avd)
		}
	}
}


func diskReport(r mbr){
	total_disk := r.Size
	Content += "label=<" + "\n"
	Content += "<table border='1' cellborder='1'>" + "\n"
	mbr_per := CalcPercentage(r.Size, int64(unsafe.Sizeof(r)))
	total_disk -= int64(unsafe.Sizeof(r))
	Content += "<tr><td>MBR "+ strconv.Itoa(int(mbr_per)) + "%</td>" + "\n"
	for _,e :=range r.Partitions{
		if e.Status!='0'{
			if e.Type=='p' {
				Content += "<td>Primaria " +strconv.Itoa(CalcPercentage(r.Size, e.Size))  + "%</td>"+ "\n"
				total_disk -= e.Size
			}else if e.Type == 'e'{
				_,_, logical := calcPart(r.Partitions)
				Content += "<td>" +"\n"
				Content += "<table border='1' cellborder='1'>" + "\n"
				Content += "<tr><td colspan=\"" + strconv.Itoa(logical*3) + "\">Extendida " +strconv.Itoa(CalcPercentage(r.Size, e.Size)) + "%</td></tr>"+ "\n"
				Content += "<tr>"+ "\n"
				for _,element := range r.Partitions{
					if(element.Status!='0'){
						if(element.Type=='l' || element.Type=='L'){
							eb :=ebr{}
							ebr_size := unsafe.Sizeof(eb)
							Content += "<td>EBR "+ strconv.Itoa(CalcPercentage(e.Size, int64(ebr_size))) + "%</td>" + "\n"
							Content += "<td>Logica " + strconv.Itoa(CalcPercentage(e.Size, element.Size))  + "%</td>"+ "\n"
							
						}
					}
				}
				Content += "</tr>"+ "\n"
				Content += "</table>"+ "\n"
				Content += "</td>" +"\n"
				total_disk -= e.Size
			}
		}
	}
	if(total_disk!=0){
		Content += "<td>Libre " + strconv.Itoa(CalcPercentage(r.Size, total_disk))+ "%</td>"+ "\n"
	}
	Content += "</tr>" + "\n"
	Content += "</table>"+ "\n"
	Content += ">"
}

func SbReport(s Super_Boot){
	Content += "label=<" + "\n"
	Content += "<table border='1' cellborder='1'>" + "\n"
	Content += "<tr><td>Nombre</td><td>Valor</td></tr>" + "\n"
	Content += "<tr><td>Nombre del disco</td><td>"+GetString(s.Virtual_disk_name[:])+ "</td></tr>" + "\n"
	Content += "<tr><td>Cantidad de estructuras en el arbol del directorio</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de estructuras en el detalle de directorio</td><td>"+strconv.Itoa(int(s.Directory_detail_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de inodos</td><td>"+strconv.Itoa(int(s.Inodes_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de bloques de datos</td><td>"+strconv.Itoa(int(s.Block_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de estructuras en el arbol del directorio libres</td><td>"+strconv.Itoa(int(s.Free_virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de estructuras en el detalle de directorio libres</td><td>"+strconv.Itoa(int(s.Free_directory_detail_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de inodos libres</td><td>"+strconv.Itoa(int(s.Free_inodes_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de bloques de datos libres</td><td>"+strconv.Itoa(int(s.Free_block_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Fecha y hora de creacion</td><td>"+GetString(s.Creation_date[:])+ "</td></tr>"+ "\n"
	Content += "<tr><td>Fecha y hora de ultima modificacion</td><td>"+GetString(s.Last_mount_date[:])+ "</td></tr>"+ "\n"
	Content += "<tr><td>Contador de montajes</td><td>"+strconv.Itoa(int(s.Mount_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio del bitmap de avd</td><td>"+strconv.Itoa(int(s.Inp_bitmap_directory_tree))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio del avd</td><td>"+strconv.Itoa(int(s.Inp_directory_tree))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio del bitmap de dd</td><td>"+strconv.Itoa(int(s.Inp_bitmap_directory_detail))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio del dd</td><td>"+strconv.Itoa(int(s.Inp_directory_detail))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio del bitmap de inodo</td><td>"+strconv.Itoa(int(s.Inp_bitmap_inode_table))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio del inodo</td><td>"+strconv.Itoa(int(s.Inp_inode_table))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio del bitmap de bloques</td><td>"+strconv.Itoa(int(s.Inp_bitmap_block))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio de los bloques</td><td>"+strconv.Itoa(int(s.Inp_block))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Apuntador al inicio de los bloques</td><td>"+strconv.Itoa(int(s.Inp_bitacora))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Tamanio de la estructura del avd</td><td>"+strconv.Itoa(int(s.Directory_tree_size))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Tamanio de la estructura del dd</td><td>"+strconv.Itoa(int(s.Directory_detail_size))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Tamanio de la estructura del inodo</td><td>"+strconv.Itoa(int(s.Inode_size))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Tamanio de la estructura de los bloques</td><td>"+strconv.Itoa(int(s.Block_size))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Primer bit libre del bitmap de avd</td><td>"+strconv.Itoa(int(s.Ffb_directory_tree))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Primer bit libre del bitmap de dd</td><td>"+strconv.Itoa(int(s.Ffb_directory_detail))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Primer bit libre del bitmap de inodo</td><td>"+strconv.Itoa(int(s.Ffb_inode_table))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Primer bit libre del bitmap de bloques</td><td>"+strconv.Itoa(int(s.Ffb_block))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Numero magico</td><td>"+strconv.Itoa(int(s.Magic_num))+ "</td></tr>"+ "\n"
	Content += "</table>"+ "\n"
	Content += ">"
}


func MbrReport(m mbr){
	Content += "label=<" + "\n"
	Content += "<table border='1' cellborder='1'>"+ "\n"
	Content += "<tr><td>Nombre</td><td>Valor</td></tr>"+ "\n"
	Content += "<tr><td>Disk size</td><td>"+strconv.Itoa(int(m.Size))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Creation date</td><td>"+GetString(m.Time[:])+ "</td></tr>"+ "\n"
	Content += "<tr><td>Disk signature</td><td>"+strconv.Itoa(int(m.Disk_signature))+ "</td></tr>"+ "\n"
	for _,e := range m.Partitions{
		if e.Status!='0' && e.Type != '0' {
			Content +="<tr><td colspan=\"3\">"+ GetString(e.Name[:]) + "</td></tr>"+ "\n"
			Content += "<tr><td>Status</td><td>"+string(e.Status)+ "</td></tr>"+ "\n"
			Content += "<tr><td>Type</td><td>"+string(e.Type)+ "</td></tr>"+ "\n"
			Content += "<tr><td>Fit</td><td>"+GetString(e.Fit[:])+ "</td></tr>"+ "\n"
			Content += "<tr><td>Start</td><td>"+strconv.Itoa(int(e.Start))+ "</td></tr>"+ "\n"
			Content += "<tr><td>Size</td><td>"+strconv.Itoa(int(e.Size))+ "</td></tr>"+ "\n"
		}
	}
	Content += "</table>"+ "\n"
	Content += ">"
}