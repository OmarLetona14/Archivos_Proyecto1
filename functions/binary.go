package functions

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"unsafe"
	"fmt"
)

func WriteSuperB(file_path string, super Super_Boot, init int64, final_bit int64){
	file, err := os.OpenFile(file_path, os.O_RDWR, 0664)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot write the file")
	}
	//Escribir el superboot en el principio de la particion
	file.Seek(init, 0)
	ss := &super
	var mbr_buf bytes.Buffer
	binary.Write(&mbr_buf, binary.BigEndian, ss)
	WriteBytes(file, mbr_buf.Bytes())
	//Escribimos un 0 al final del archivo.
	file.Seek(final_bit,0)
	var otro int8 = 0
	s := &otro
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
	file.Seek(0,0)
}


func WriteAVD(file_path string, n_avd avd_binary, init int64, final_bit int64){
	file, err := os.OpenFile(file_path, os.O_RDWR, 0664)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot write the file")
	}
		//Escribir el superboot en el principio de la particion
	file.Seek(init, 0)
	ss := &n_avd
	var mbr_buf bytes.Buffer
	binary.Write(&mbr_buf, binary.BigEndian, ss)
	WriteBytes(file, mbr_buf.Bytes())
	//Escribimos un 0 al final del archivo.
	file.Seek(final_bit,0)
	var otro int8 = 0
	s := &otro
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
	file.Seek(0,0)
}


func ModifyMBR(file_path string, rec mbr){
	file, err := os.OpenFile(file_path, os.O_RDWR, 0664)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot write the file")
	}
	ss := &rec
	file.Seek(0,0)
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
	file.Seek(0,0)
}


func ModifySB(file_path string, rec Super_Boot, part_init int64, final int64){
	file, err := os.OpenFile(file_path, os.O_RDWR, 0664)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot write the file")
	}
	ss := &rec
	file.Seek(part_init, 0)
	var mbr_buf bytes.Buffer
	binary.Write(&mbr_buf, binary.BigEndian, ss)
	WriteBytes(file, mbr_buf.Bytes())

	var otro int8 = 0
	s := &otro
	file.Seek(final, 0)
	
	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
	file.Seek(0,0)
}

func ModifyAVD(file_path string, a avd_binary, part_init int64, final int64){
	file, err := os.OpenFile(file_path, os.O_RDWR, 0664)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot write the file")
	}
	ss := &a
	file.Seek(part_init, 0)
	var mbr_buf bytes.Buffer
	binary.Write(&mbr_buf, binary.BigEndian, ss)
	WriteBytes(file, mbr_buf.Bytes())

	var otro int8 = 0
	s := &otro
	file.Seek(final, 0)
	
	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
	file.Seek(0,0)
}

func ModifyBitmap(file_path string, bitmap_init int64, final int64){
	file, err := os.OpenFile(file_path, os.O_RDWR, 0664)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot write the file")
	}
	var otro1 int8 = 1
	ss := &otro1
	file.Seek(bitmap_init, 0)
	var mbr_buf bytes.Buffer
	binary.Write(&mbr_buf, binary.BigEndian, ss)
	WriteBytes(file, mbr_buf.Bytes())

	var otro int8 = 0
	s := &otro
	file.Seek(final, 0)
	
	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
	file.Seek(0,0)
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
	file.Seek(0, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	WriteBytes(file, binario.Bytes())
	//Nos posicionamos en el byte file_size -1  (primera posicion es 0)	
	file.Seek(file_size-1, 0)
	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	WriteBytes(file, binario2.Bytes())
	file.Seek(0, 0)// nos posicionamos en el inicio del archivo.
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
	binario3 := bytes.NewBuffer([]byte{})
	binary.Write(binario3, binary.BigEndian, s1)
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
	file.Seek(0,0)
	mr := mbr{}
	//Obtenemos el tamanio del mbr
	var size int = int(unsafe.Sizeof(mr))
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
	file.Seek(0,0)
	return
}

func ReadSB(file_path string, part_init int64)(m Super_Boot) {
	//Abrimos/creamos un archivo.
	file, err := os.Open(file_path)
	defer file.Close() 
	if err != nil { //validar que no sea nulo.
		log.Fatal(err)
	}
	file.Seek(part_init, 0)
	//Obtenemos el tamanio del SUPER_BOOT
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
	file.Seek(0,0)
	return
}

func ReadAVD(file_path string, part_init int64)(m avd_binary) {
	//Abrimos/creamos un archivo.
	file, err := os.Open(file_path)
	defer file.Close() 
	if err != nil { //validar que no sea nulo.
		log.Fatal("Path " ,file_path, err)
	}
	file.Seek(part_init, 0)
	//Obtenemos el tamanio del AVD
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
	file.Seek(0,0)
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
