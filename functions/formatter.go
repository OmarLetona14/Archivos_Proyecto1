package functions

import(
	"fmt"
	"unsafe"
	
)

func Format(sb *Super_Boot, partition_size float64, partition_init float64){
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
	fmt.Println("SUPER BLOCK SIZE",sb_size)
	fmt.Println("AVD SIZE", avd_size)
	fmt.Println("DD SIZE", dd_size)
	fmt.Println("INODE SIZE", inode_size)
	fmt.Println("BLOCK SIZE", block_size)
	fmt.Println("BITACORA SIZE", bita_size)
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

	fmt.Println("SUPER BLOCK INIT",sb_init)
	fmt.Println("AVD BITMAP INIT",avd_bitmap_init)
	fmt.Println("AVD INIT",avd_init)
	fmt.Println("DD BITMAP INIT",dd_bitmap_init)
	fmt.Println("DD INIT",dd_init)
	fmt.Println("INODE BITMAP INIT",inode_bitmap_init)
	fmt.Println("INODE INIT",inode_init)
	fmt.Println("BLOCK BITMAP INIT",block_bitmap_init)
	fmt.Println("BLOCK INIT",block_init)
	fmt.Println("BITACAORA INIT",bita_init)
	fmt.Println("BITACORA FIN",bita_fin)
	fmt.Println("SUPER BLOQUE COPIA",sb_backup)
	
	//Verificar si llegamos a el final de la particion, sino, ocurrio un error
	fmt.Println("FINAL VALUE", sb_backup)
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
	sb.Inp_bitmap_directory_detail = dd_bitmap_init
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