package reversebits

// do not understand this solution at all
// got it from internet

func reverseBits(n uint32) uint32 {
	n = (n>>1)&0x55555555 | (n<<1)&0xaaaaaaaa
	n = (n>>2)&0x33333333 | (n<<2)&0xcccccccc
	n = (n>>4)&0x0f0f0f0f | (n<<4)&0xf0f0f0f0
	n = (n>>8)&0x00ff00ff | (n<<8)&0xff00ff00
	n = (n>>16)&0x0000ffff | (n<<16)&0xffff0000
	return n
}
