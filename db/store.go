package db

import "fmt"

type Store struct {
	Segment
	position int
	length   int
}

type Segment struct {
	name string
	num  uint
	size int
}

func (s *Segment) Increment() {
	s.num++
	s.name = SetName(s.name, s.num)
}

func NewSegment(name string, num uint) Segment {
	return Segment{
		name: SetName(name, num),
	}
}

func (s *Segment) UpdateSize(length int) {
	s.size += length
}

func SetName(name string, num uint) string {
	return fmt.Sprintf("%s%d", name, num)
}

func NewStore(dataLength int, segment *Segment) Store {
	store := Store{
		Segment:  *segment,
		position: segment.size,
		length:   dataLength,
	}

	segment.UpdateSize(dataLength)

	return store
}
