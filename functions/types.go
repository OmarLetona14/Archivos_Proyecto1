package functions


type mbr struct{
	Size int64
	Time[25] byte
	Disk_signature int8
	Partitions[4] Partition
}


type ebr struct{
	Status byte
	Fit byte
	Start int64 
	Size int64
	Next int64
	Name [16]byte
}

type Partition struct{
	Status byte
	Type byte 
	Fit [2]byte
	Start int64
	Size int64
	Name[16] byte
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

type Mkfs_command struct{
	Identifier string
	Type string
	Add bool
	Unit byte	
}

type Mount_command struct{
	Path string
	Name string
}

type Unmount_command struct{
	List_id [20]string
}

type Mkfile_command struct{
	Id string
	Path string
	P bool
	Size int64
	Cont string
}

type Mkdir_command struct{
	Id string
	Path string
	P bool
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
	Init int64
	Size int64
	Identifier string
	Formatted bool
	Dsk Mounted_disk
}

type Super_Boot struct{
	Virtual_disk_name [20]byte
	Virtual_tree_count int64
	Directory_detail_count int64
	Inodes_count int64
	Block_count int64
	Free_virtual_tree_count int64
	Free_directory_detail_count int64
	Free_inodes_count int64
	Free_block_count int64
	Creation_date [25] byte
	Last_mount_date [25] byte
	Mount_count int64
	Inp_bitmap_directory_tree int64
	Inp_directory_tree int64
	Inp_bitmap_directory_detail int64
	Inp_directory_detail int64
	Inp_bitmap_inode_table int64
	Inp_inode_table int64
	Inp_bitmap_block int64
	Inp_block int64
	Inp_bitacora int64
	Directory_tree_size int64
	Directory_detail_size int64
	Inode_size int64
	Block_size int64
	Ffb_directory_tree int64
	Ffb_directory_detail int64
	Ffb_inode_table int64
	Ffb_block int64
	Magic_num int64
}

type avd_binary struct{
	Id int64
	Creation_date [25]byte
	Directory_name [25]byte
	Sub_directory_pointers [6] int64
	Directory_detail int64
	Avd_next int64
	Proper [25]byte
}

type dd struct{
	Table [5] inode_register
	Next_pointer *dd
}

type inode_register struct{
	File_name [25] byte
	Inode_pointer *inode
	Creation_date [25] byte
	Modification_date [25]byte
}


type inode struct{
	Inode_number int64
	File_size int64
	Block_count int64
	Block_array [4] block
	Indirect_pointer *inode
}

type block struct{
	Data [25] byte
}


type bitacora struct{
	operation_type [25]byte
	typee byte
	file_name [25]byte
	content [50]byte
	transaction_date [25]byte
}
