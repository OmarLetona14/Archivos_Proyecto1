package functions

import(
	"fmt"
	"unsafe"
	
)

func Format(sb *Super_Boot, disk_name string, partition_size float64, partition_init float64){
	avd_struct := avd{}
	dd_struct := dd{}
	inode_struct := inode{}
	block_struct := block{}
	bit := bitacora{}
	//Calcular los tamanios de cada estructura
	avd_size := float64(unsafe.Sizeof(avd_struct))
	dd_size := float64(unsafe.Sizeof(dd_struct))
	inode_size := float64(unsafe.Sizeof(inode_struct))
	block_size := float64(unsafe.Sizeof(block_struct))
	bita_size := float64(unsafe.Sizeof(bit))
	sb_size := float64(unsafe.Sizeof(sb))
	//Numero de estructuras
	struct_count := (partition_size - (2.0*sb_size)) / (27.0+ avd_size + dd_size +(5.0*inode_size +(20.0*block_size)+ bita_size)) 
	//Calcular la cantidad de cada estructura
	avd_count := struct_count
	dd_count := struct_count
	inode_count := 5*struct_count
	block_count := 4*inode_count
	bita_count := struct_count
	
	//Calcular los inicios de cada estructura dentro de la particion
	sb_init := partition_init
	avd_bitmap_init := sb_init + sb_size
	avd_init := avd_bitmap_init+avd_count
	dd_bitmap_init := avd_init + (avd_count*avd_size)
	dd_init := dd_bitmap_init + dd_count
	inode_bitmap_init := dd_init + (dd_count*dd_size)
	inode_init := inode_bitmap_init + inode_count
	block_bitmap_init := inode_init + (inode_count*inode_size)
	block_init := block_bitmap_init + block_count
	bita_init := block_init + (block_size*block_count)
	bita_fin := bita_init + (bita_size*bita_count)
	sb_backup := bita_fin + sb_size	
	fmt.Println("FIN DE LA PARTICION", sb_backup)
	//Asignar valores del formateo al Super Boot
	sb.Virtual_tree_count = avd_count
	sb.Directory_detail_count = dd_count
	sb.Inodes_count = inode_count
	sb.Block_count = block_count
	sb.Free_virtual_tree_count,sb.Ffb_directory_tree = calcFreeAvd(avd_bitmap_init,avd_init)
	sb.Free_directory_detail_count, sb.Ffb_directory_detail = calcFreeDd(dd_bitmap_init, dd_init)
	sb.Free_inodes_count, sb.Ffb_inode_table = calcFreeInode(inode_bitmap_init, inode_init)
	sb.Free_block_count, sb.Ffb_block = calcFreeBlock(block_bitmap_init, block_init)
	sb.Creation_date = GetCurrentTime()
	sb.Last_mount_date = GetCurrentTime()
	sb.Mount_count = 1
	sb.Inp_bitmap_directory_tree = avd_bitmap_init
	sb.Inp_directory_tree = avd_init
	sb.Inp_directory_tree = dd_bitmap_init
	sb.Inp_directory_detail = dd_init
	sb.Inp_bitmap_inode_table = inode_bitmap_init
	sb.Inp_inode_table = inode_init
	sb.Inp_bitmap_block = block_bitmap_init
	sb.Inp_block = block_init
	sb.Inp_bitacora = bita_init
	sb.Directory_tree_size = avd_size
	sb.Directory_detail_size = dd_size
	sb.Inode_size = inode_size
	sb.Block_size = block_size
	printSB(sb, int64(partition_init))
}

func printSB(sb *Super_Boot, part_init int64){
	//Inicio detalles generales
	fmt.Println("DISK NAME", sb.Virtual_disk_name)
	//Cantidades
	fmt.Println("CANTIDAD DE ESTRUCTURAS DEL ARBOL VIRTUAL:",sb.Virtual_tree_count)
	fmt.Println("CANTIDAD DE ESTRUCTURAS DEL DETALLE DE DIRECTORIO:",sb.Directory_detail_count)
	fmt.Println("CANTIDAD DE INODOS:",sb.Inodes_count)
	fmt.Println("CANTIDAD DE BLOQUES DE DATOS:",sb.Block_count)
	fmt.Println("CANTIDAD DE ESTRUCTURAS DEL ARBOL VIRTUAL LIBRES:",sb.Virtual_tree_count)
	fmt.Println("CANTIDAD DE ESTRUCTURAS DEL DETALLE DE DIRECTORIO LIBRES:",sb.Virtual_tree_count)
	fmt.Println("CANTIDAD DE INODOS LIBRES:",sb.Virtual_tree_count)
	fmt.Println("CANTIDAD DE BLOQUES DE DATOS LIBRES:",sb.Virtual_tree_count)
	fmt.Println("FECHA DE CREACION:", string(sb.Creation_date[:]))
	fmt.Println("FECHA DE ULTIMA MODIFICACION:", string(sb.Last_mount_date[:]))
	fmt.Println("CONTADOR DE MONTAJES DEL SISTEMA DE ARCHIVOS:", sb.Mount_count)
	//Inicios
	fmt.Println("INICIO DEL SUPER BLOQUE:",part_init)
	fmt.Println("INICIO DEL BITMAP DEL AVD:",sb.Inp_bitmap_directory_tree )
	fmt.Println("INICIO DEL AVD:",sb.Inp_directory_tree)
	fmt.Println("INICIO DEL BITMAP DEL DD:",sb.Inp_directory_tree)
	fmt.Println("INICIO DEL DD:",sb.Inp_directory_detail)
	fmt.Println("INICIO DEL BITMAP DEL INODO:",sb.Inp_bitmap_inode_table)
	fmt.Println("INICIO DEL INODO:",sb.Inp_inode_table)
	fmt.Println("INICIO DEL BITMAP DEL BLOQUE:",sb.Inp_bitmap_block)
	fmt.Println("INICIO DEL BLOQUE:",sb.Inp_block)
	fmt.Println("INICIO DE LA BITACORA:",sb.Inp_bitacora)
	//Tamanios
	fmt.Println("TAMANIO DEL AVD", sb.Directory_tree_size)
	fmt.Println("TAMANIO DEL DD", sb.Directory_detail_size)
	fmt.Println("TAMANIO DEL INODO", sb.Inode_size)
	fmt.Println("TAMANIO DEL BLOQUE DE DATOS", sb.Block_size)

}


func update_SB(u_sb Super_Boot){
	
}

func calcFreeAvd(init, final float64) (free, first_free float64){
	free =0
	first_free = 0
	return
}

func calcFreeDd(init, final float64) (free, first_free float64){
	free =0
	first_free = 0
	return
}

func calcFreeInode(init, final float64) (free, first_free float64){
	free =0
	first_free = 0
	return
}

func calcFreeBlock(init, final float64) (free, first_free float64){
	free =0
	first_free=0
	return
}