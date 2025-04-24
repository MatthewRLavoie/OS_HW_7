package main

// RAID1 mirrors data to all disks for redundancy
type RAID1 struct {
	Disks     []*Disk
	BlockSize int
}

// Writes same data to all disks
func (r *RAID1) Write(blockNum int, data []byte) error {
	for _, disk := range r.Disks {
		if err := disk.WriteBlock(blockNum, r.BlockSize, data); err != nil {
			return err
		}
	}
	return nil
}

// Reads from first disk (assuming all are same)
func (r *RAID1) Read(blockNum int) ([]byte, error) {
	return r.Disks[0].ReadBlock(blockNum, r.BlockSize)
}
