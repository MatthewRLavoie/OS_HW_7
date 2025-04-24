package main

type RAID1 struct {
	Disks     []*Disk
	BlockSize int
}

func (r *RAID1) Write(blockNum int, data []byte) error {
	for _, disk := range r.Disks {
		if err := disk.WriteBlock(blockNum, r.BlockSize, data); err != nil {
			return err
		}
	}
	return nil
}

func (r *RAID1) Read(blockNum int) ([]byte, error) {
	return r.Disks[0].ReadBlock(blockNum, r.BlockSize)
}
