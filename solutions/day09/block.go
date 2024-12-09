package day09

import "fmt"

type Block struct {
	id    int
	start int
	end   int
}

func (b Block) getLength() int {
	return b.end - b.start + 1
}
func (b *Block) setBound(newStart, newEnd int) {
	b.start = newStart
	b.end = newEnd
}
func (b *Block) setBoundWithBlock(other Block) {
	b.start = other.start
	b.end = other.end
}

func (b Block) IsEmpty() bool {
	return b.start == 0 &&
		b.end == 0
}

func (b *Block) splitBlock(other *Block) (otherPart Block, err error) {
	if b.getLength() < other.getLength() {
		return Block{}, fmt.Errorf("block to small")
	}
	if b.getLength() == other.getLength() {

		newStart := b.start
		newEnd := b.end
		b.setBoundWithBlock(*other)
		other.setBound(newStart, newEnd)
		return Block{}, nil
	}

	oldStart, oldEnd := other.start, other.end
	other.setBound(b.start, b.start+other.getLength()-1)
	otherPart = Block{-1, other.end + 1, b.end}
	b.setBound(oldStart, oldEnd)

	return otherPart, nil
}
