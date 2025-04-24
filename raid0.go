package main

type RAID0 struct {
	Disks     []*Disk
	BlockSize int
}

func (r *RAID0) Write(blockNum int, data []byte) error {
	disk := r.Disks[blockNum%len(r.Disks)]
	return disk.WriteBlock(blockNum/len(r.Disks), r.BlockSize, data)
}

func (r *RAID0) Read(blockNum int) ([]byte, error) {
	disk := r.Disks[blockNum%len(r.Disks)]
	return disk.ReadBlock(blockNum/len(r.Disks), r.BlockSize)
}
