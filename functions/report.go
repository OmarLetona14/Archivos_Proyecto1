package functions

import(
	"strconv"
	"os"
	"bufio"
	"log"
	"fmt"
)

var Content string

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
}

func createTreeReport(r *avd){
	if(r!=nil){
		Content += "digraph G{ \n"
		GetContent(r)
		Content+="}"
	}else{
		fmt.Print("Empty tree")
	}
	
}

func GetContent(r *avd){
	Content += "c" + strconv.Itoa(r.Id) +" [shape =record label=\"{" + string(r.Directory_name[:]) + " | {"
	for i,e := range r.Sub_directory_pointers {
		if(i!=0){
			if(e!=nil){
				Content += "| <f" + strconv.Itoa(i) + "> " + strconv.Itoa(e.Id)
			}else{
				Content += "| <f" + strconv.Itoa(i) + "> -1"
			}
		}else{
			if(e!=nil){
				Content += " <f" + strconv.Itoa(i) + "> " + strconv.Itoa(e.Id)
			}else{
				Content += " <f" + strconv.Itoa(i) + "> -1"
			}
		}
	}
	if(r.Avd_next==nil){
		Content += "| <f6> -1"
	}else{

	}
	if(r.Directory_detail==nil){
		Content += "| <f7> -1"
	}else{

	}
	Content += "}}\"]\n"
	for i,e := range r.Sub_directory_pointers{
		if(e!=nil){
			Content += "c" + strconv.Itoa(r.Id) + ":f" + strconv.Itoa(i) + " -> c" + strconv.Itoa(e.Id) +"\n"
		}
	}
	for _,e :=range r.Sub_directory_pointers{
		if(e!=nil){
			GetContent(e)
		}
	}
}