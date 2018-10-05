package zmachine

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
)

// StoryFile is
type StoryFile struct {
	Raw    []byte
	Header *Header
}

// Cursor is a helper for reading a sequence of bytes from a story at a given location
type Cursor struct {
	Story *StoryFile
	Index uint32
}

type ByteAddress uint16
type WordAddress uint16
type PackedAddress uint16

// Header is
type Header struct {
	Version            byte        // 0
	Flags1             byte        // 1
	Release            uint16      // 2
	HighMemoryBase     ByteAddress // 4
	StartPC            ByteAddress // 6
	Dictionary         ByteAddress // 8
	ObjectTable        ByteAddress // A
	Globals            ByteAddress // C
	StaticMemory       ByteAddress // E
	Flags2             uint16      // 10
	Serial             [6]byte     // 12
	Abbreviations      ByteAddress // 18
	LengthOfFile       uint16      // 1A
	Checksum           uint16      // 1C
	InterpreterNumber  byte        // 1E
	InterpreterVersion byte        // 1F
}

func LoadStory(fname string) (*StoryFile, error) {
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(dat)
	hdr := &Header{}
	if err = binary.Read(r, binary.BigEndian, hdr); err != nil {
		return nil, err
	}
	story := &StoryFile{
		Raw:    dat,
		Header: hdr,
	}
	return story, nil
}
