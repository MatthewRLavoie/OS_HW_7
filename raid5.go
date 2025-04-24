package main

// RAID5 distributes parity across all disks
type RAID5 struct {
	Disks     []*Disk
	BlockSize int
}

// Returns number of data disks (1 disk used for parity)
func (r *RAID5) numDataDisks() int {
	return len(r.Disks) - 1
}

// Maps logical block to stripe, disk, and parity disk
func (r *RAID5) getStripeInfo(blockNum int) (stripe int, diskIndex int, parityDisk int) {
	numData := r.numDataDisks()
	stripe = blockNum / numData
	diskIndex = blockNum % numData
	parityDisk = stripe % len(r.Disks)
	if diskIndex >= parityDisk {
		diskIndex++
	}
	return
}

// Writes block with distributed parity
func (r *RAID5) Write(blockNum int, data []byte) error {
	stripe, diskIndex, parityDisk := r.getStripeInfo(blockNum)
	dataBlocks := make([][]byte, len(r.Disks))

	for i := range r.Disks {
		if i == diskIndex {
			dataBlocks[i] = data
		} else if i != parityDisk {
			block, err := r.Disks[i].ReadBlock(stripe, r.BlockSize)
			if err != nil {
				block = make([]byte, r.BlockSize)
			}
			dataBlocks[i] = block
		} else {
			dataBlocks[i] = make([]byte, r.BlockSize) // placeholder
		}
	}

	parity := xorBlocks(dataBlocks)

	if err := r.Disks[diskIndex].WriteBlock(stripe, r.BlockSize, data); err != nil {
		return err
	}
	return r.Disks[parityDisk].WriteBlock(stripe, r.BlockSize, parity)
}

// Reads block from appropriate disk
func (r *RAID5) Read(blockNum int) ([]byte, error) {
	stripe, diskIndex, _ := r.getStripeInfo(blockNum)
	return r.Disks[diskIndex].ReadBlock(stripe, r.BlockSize)
}
