package helper

type CustomColor struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c CustomColor) RGBA() (uint32, uint32, uint32, uint32) {
	// convert components to float32 in the range [0.0, 1.0]
	r32f := float32(c.R) / float32(255.0)
	g32f := float32(c.G) / float32(255.0)
	b32f := float32(c.B) / float32(255.0)
	a32f := float32(c.A) / float32(255.0)

	// perform alpha-premultiplication
	r32f = r32f * a32f
	g32f = g32f * a32f
	b32f = b32f * a32f

	// convert back to 16 bit
	r16 := uint16(r32f * 65535.0)
	g16 := uint16(g32f * 65535.0)
	b16 := uint16(b32f * 65535.0)
	a16 := uint16(a32f * 65535.0)

	// return as 32 bit (without upscaling)
	return uint32(r16), uint32(g16), uint32(b16), uint32(a16)
}
