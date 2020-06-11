package book

import "fmt"

type book struct {
	wear uint
}

func CreateBook(wear uint) *book {
	return &book{wear: wear}
}

func (b *book) checkWear() bool {
	return b.wear < 100
}

func (b *book) Open() bool {
	success := b.checkWear()
	if success {
		b.wear -= 1
	}
	return success
}

func (b *book) Read() bool {
	success := b.checkWear()
	if success {
		b.wear -= 5
	}
	return success
}

func (b *book) Close() {
	b.wear -= 1
}

func (b *book) State() string {
	return fmt.Sprintf("book - wear: %v", b.wear)
}