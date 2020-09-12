package functions

import(
	"strconv"
	"os"
	"bufio"
	"log"
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
	nm:="Directorio" + p.Identifier +"_Report"
	createDotFile(path + nm + ".dot")
	execDot(path + nm + ".dot", path + nm + ".png")
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
	nm:="SB" + p.Identifier +"_Report"
	createDotFile(path + nm + ".dot")
	execDot(path + nm + ".dot", path + nm + ".png")
}

func createMbrReport(r mbr, path string){
	if(r.Time[0]!=0){
		Content += "digraph G{ \n"
		MbrReport(r)
		Content+="}"
	}else{
		fmt.Print("Empty table")
	}
	nm:="MBR" + strconv.Itoa(int(r.Disk_signature)) +"_Report"
	createDotFile(path + nm + ".dot")
	execDot(path + nm + ".dot", path + nm + ".png")
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

func SbReport(s Super_Boot){
	Content += "tbl [" + "\n"
	Content += "label=<" + "\n"
	Content += "<table border='1' cellborder='1'>" + "\n"
	Content += "<tr><td>Nombre</td><td>Valor</td></tr>" + "\n"
	Content += "<tr><td>Nombre del disco</td><td>"+GetString(s.Virtual_disk_name[:])+ "</td></tr>" + "\n"
	Content += "<tr><td>Cantidad de estructuras en el arbol del directorio</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de estructuras en el detalle de directorio</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de inodos</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de bloques de datos</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de estructuras en el arbol del directorio libres</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de estructuras en el detalle de directorio libres</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de inodos libres</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Cantidad de bloques de datos libres</td><td>"+strconv.Itoa(int(s.Virtual_tree_count))+ "</td></tr>"+ "\n"
	Content += "<tr><td>Fecha y hora de creacion</td><td>"+GetString(s.Creation_date[:])+ "</td></tr>"+ "\n"
	Content += "<tr><td>Fecha y hora de ultima moficacion</td><td>"+GetString(s.Last_mount_date[:])+ "</td></tr>"+ "\n"
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
	Content += ">];"
}


func MbrReport(m mbr){
	Content += "tbl [" + "\n"
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
	Content += ">];"
}