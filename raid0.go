package main

// RAID0 splits data across disks (striping), no redundancy
type RAID0 struct {
	Disks     []*Disk
	BlockSize int
}

// Writes block to striped disk
func (r *RAID0) Write(blockNum int, data []byte) error {
	disk := r.Disks[blockNum%len(r.Disks)]
	return disk.WriteBlock(blockNum/len(r.Disks), r.BlockSize, data)
}

// Reads block from striped disk
func (r *RAID0) Read(blockNum int) ([]byte, error) {
	disk := r.Disks[blockNum%len(r.Disks)]
	return disk.ReadBlock(blockNum/len(r.Disks), r.BlockSize)
}
