package functions

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"unsafe"
	"fmt"
)


func ModifyMBR(file_path string, rec mbr){
	file, err := os.Create(file_path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot write the file")
	}

	rec_size := unsafe.Sizeof(rec)
	fmt.Println("REC SIZE ", rec_size)
	ss := &rec
	
	var mbr_buf bytes.Buffer
	binary.Write(&mbr_buf, binary.BigEndian, ss)
	WriteBytes(file, mbr_buf.Bytes())



	var otro int8 = 0
	s := &otro
	file.Seek(rec.Size-1, 0)
	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
}

func WriteFile(file_path string, file_size int64) {
	file, err := os.Create(file_path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//ESTE CERO SERA PUESTO EN EL PRINCIPIO Y EL FINAL DE ARCHIVO
	var otro int8 = 0
	s := &otro
	//Escribimos un 0 en el inicio del archivo.
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	WriteBytes(file, binario.Bytes())
	//Nos posicionamos en el byte file_size -1  (primera posicion es 0)	
	file.Seek(file_size-1, 0)
	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
	file.Seek(0, 0) // nos posicionamos en el inicio del archivo.
	//MBR vacio
	var mbrs mbr
	//TAMANIO DEL MBR
	mbrs.Size = file_size
	//NUMERO RANDOM
	mbrs.Disk_signature = generateRandom()
	//OBTIENE UN BYTE[25] QUE REPRESENTA EL TIEMPO ACTUAL
	mbrs.Time = GetCurrentTime()
	//LLENAMOS TODOS LOS STATUS CON UN 0
	for i:=0;i<len(mbrs.Partitions);i++ {
		mbrs.Partitions[i].Status = '0'
	}
	s1 := &mbrs
	//Escribimos struct.
	var binario3 bytes.Buffer
	binary.Write(&binario3, binary.BigEndian, s1)
	WriteBytes(file, binario3.Bytes())
	//REGRESAMOS AL PRINCIPIO DEL ARCHIVO
	file.Seek(0, 0)
}

func ReadMBR(file_path string)(m mbr) {
	//Abrimos/creamos un archivo.
	file, err := os.Open(file_path)
	defer file.Close() 
	if err != nil { //validar que no sea nulo.
		log.Fatal(err)
	}
	//Obtenemos el tamanio del mbr
	var size int = int(unsafe.Sizeof(m))
	//Lee la cantidad de <size> bytes del archivo
	data := ReadBytes(file, size)
	//Convierte la data en un buffer,necesario para
	//decodificar binario
	buffer := bytes.NewBuffer(data)
	//Decodificamos y guardamos en la variable m
	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	return
}

func WriteBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes) 
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

//FUNCION PARA IMPRIMIR EL CONTENIDO DEL MBR
func PrintMBR(m mbr){
	fmt.Println("Disk size", m.Size)
	myString := string(m.Time[:])
	fmt.Println("Disk created at", myString)
	fmt.Println("Disk signature", m.Disk_signature)
	for i := 0; i < len(m.Partitions); i++ {
		par := m.Partitions[i]
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Partition", i)
		fmt.Println("Partition status", string(par.Status))
		fmt.Println("Partition type", string(par.Type))
		fmt.Println("Partition fit", string(par.Fit[0]))
		fmt.Println("Partition start",par.Start)
		fmt.Println("Partition size",par.Size)
		parName := string(par.Name[:])
		fmt.Println("Partition name",parName)
		fmt.Println("------------------------------------------------------------------")
	}
}
