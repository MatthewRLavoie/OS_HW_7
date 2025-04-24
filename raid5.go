package main

type RAID5 struct {
	Disks     []*Disk
	BlockSize int
}

func (r *RAID5) numDataDisks() int {
	return len(r.Disks) - 1
}

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

func (r *RAID5) Write(blockNum int, data []byte) error {
	stripe, diskIndex, parityDisk := r.getStripeInfo(blockNum)
	dataBlocks := make([][]byte, len(r.Disks))

	for i := range r.Disks {
		if i == diskIndex {
			dataBlocks[i] = data
		} else if i != parityDisk {
			block, err := r.Disks[i].ReadBlock(stripe, r.BlockSize)
			if err != nil {
				block = make([]byte, r.BlockSize) // fallback to zeros
			}
			dataBlocks[i] = block
		} else {
			dataBlocks[i] = make([]byte, r.BlockSize) // temp placeholder
		}
	}

	parity := xorBlocks(dataBlocks)

	if err := r.Disks[diskIndex].WriteBlock(stripe, r.BlockSize, data); err != nil {
		return err
	}
	return r.Disks[parityDisk].WriteBlock(stripe, r.BlockSize, parity)
}

func (r *RAID5) Read(blockNum int) ([]byte, error) {
	stripe, diskIndex, _ := r.getStripeInfo(blockNum)
	return r.Disks[diskIndex].ReadBlock(stripe, r.BlockSize)
}
