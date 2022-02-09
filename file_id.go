package tg_file_id

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/arisudesu/go-tg-file-id/internal"
)

type FileID struct {
	Version       uint8
	SubVersion    uint8
	TypeID        uint32
	DcID          uint32
	FileReference []byte
	URL           []byte
	ID            uint64
	AccessHash    uint64
}

const (
	webLocationFlag   = uint32(1 << 24)
	fileReferenceFlag = uint32(1 << 25)
)

func DecodeFileID(fileID string) (FileID, error) {
	bs, err := base64.RawURLEncoding.DecodeString(fileID)
	if err != nil {
		return FileID{}, err
	}

	rld := internal.RLEDecode(bs)

	fid := FileID{}
	fid.Version = rld[len(rld)-1]
	if fid.Version == 4 {
		fid.SubVersion = rld[len(rld)-2]
	}

	fid.TypeID, rld = binary.LittleEndian.Uint32(rld), rld[4:]
	fid.DcID, rld = binary.LittleEndian.Uint32(rld), rld[4:]

	hasReference := fid.TypeID&fileReferenceFlag != 0
	hasWebLocation := fid.TypeID&webLocationFlag != 0
	fid.TypeID &= ^fileReferenceFlag
	fid.TypeID &= ^webLocationFlag

	if hasReference {
		fid.FileReference, rld, err = internal.TLDecode(rld)
		if err != nil {
			return FileID{}, err
		}
	}
	if hasWebLocation {
		fid.URL, rld, err = internal.TLDecode(rld)
		if err != nil {
			return FileID{}, err
		}
		fid.AccessHash, rld = binary.LittleEndian.Uint64(rld), rld[8:]
		return fid, nil
	}

	fid.ID, rld = binary.LittleEndian.Uint64(rld), rld[8:]
	fid.AccessHash, rld = binary.LittleEndian.Uint64(rld), rld[8:]

	// TODO: type data
	// TODO: len check
	return fid, nil
}
