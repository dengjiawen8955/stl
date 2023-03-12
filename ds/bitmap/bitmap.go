package bitmap

// Bitmap is a mapping from some domain (for example, a range of integers) to bits. It is also called a bit array or bitmap index
type Bitmap struct {
	 //TODO: Complete me!
}

//New creates a new bitmap
func New(size uint64) *Bitmap {
	 //TODO: Complete me!
}

// NewFromData creates a bitmap from the exported data
func NewFromData(data []byte) *Bitmap {
	 //TODO: Complete me!
}

// Set sets 1 at position pos
func (b *Bitmap) Set(pos uint64) bool {
	 //TODO: Complete me!
}

// Unset sets 0 at position pos
func (b *Bitmap) Unset(pos uint64) bool {
	 //TODO: Complete me!
}

// IsSet returns true if the position pos is 1
func (b *Bitmap) IsSet(pos uint64) bool {
	 //TODO: Complete me!
}

// Resize resizes the bitmap with the passed size
func (b *Bitmap) Resize(size uint64) {
	 //TODO: Complete me!
}

// Size returns the bitmap's size in bit
func (b *Bitmap) Size() uint64 {
	 //TODO: Complete me!
}

// Clear clear the bitmap's data
func (b *Bitmap) Clear() {
	 //TODO: Complete me!
}

// Data returns the bitmap's internal data slice
func (b *Bitmap) Data() []byte {
	 //TODO: Complete me!
}
