package functions


type mbr struct{
	Size int64
	Time[25] byte
	Disk_signature int8
	Partitions[4] Partition
}

type Partition struct{
	Status byte
	Type byte 
	Fit [2]byte
	Start int64
	Size int64
	Name[20] byte
}

type Mkdisk_command struct {
	Size int
	Path string
	Name string
	Unit string
}

type Mfdisk_command struct{
	Size int64
	Unit byte
	Path string
	Type byte
	Fit[2] byte
	Delete bool
	Name string
	Add bool
}

type Mount_command struct{
	Path string
	Name string
}

type Unmount_command struct{
	List_id [20]string
}
type Mounted_disk struct{
	Name string
	Path string
	Size int
	Identifier string
	Count_mounted int
}

type Mounted_partition struct{
	Name string
	Path string
	Identifier string
	dsk Mounted_disk
}

type Super_Boot struct{
	Virtual_disk_name [10]byte
	Virtual_tree_count float64
	Directory_detail_count float64
	Inodes_count float64
	Block_count float64
	Free_virtual_tree_count float64
	Free_directory_detail_count float64
	Free_inodes_count float64
	Free_block_count float64
	Creation_date [25] byte
	Last_mount_date [25] byte
	Mount_count float64
	Inp_bitmap_directory_tree float64
	Inp_directory_tree float64
	Inp_bitmap_directory_detail float64
	Inp_directory_detail float64
	Inp_bitmap_inode_table float64
	Inp_inode_table float64
	Inp_bitmap_block float64
	Inp_block float64
	Inp_bitacora float64
	Directory_tree_size float64
	Directory_detail_size float64
	Inode_size float64
	Block_size float64
	Ffb_directory_tree float64
	Ffb_directory_detail float64
	Ffb_inode_table float64
	Ffb_block float64
	Magic_num float64
}

type avd struct{
	creation_date [25]byte
	directory_name [25]byte
	sub_directory_pointers [6] *avd
	directory_detail dd
	avd_next *avd 
	proper [25]byte
}

type dd struct{
	table [5] inode_register
	next_pointer *dd
}

type inode_register struct{
	file_name [25] byte
	inode_pointer *inode
	creation_date [25] byte
	modification_date [25]byte
}


type inode struct{
	inode_number int64
	file_size int64
	block_count int64
	block_array [4] block
	indirect_pointer *inode
}

type block struct{
	data [25] byte
}


type bitacora struct{
	operation_type [25]byte
	typee byte
	file_name [25]byte
	content [50]byte
	transaction_date [25]byte
}
