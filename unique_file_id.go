package tg_file_id

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/arisudesu/go-tg-file-id/internal"
)

type UniqueFileID struct {
	TypeID uint32
	ID     uint64
}

func DecodeUniqueFileID(uniqueFileID string) (UniqueFileID, error) {
	bs, err := base64.RawURLEncoding.DecodeString(uniqueFileID)
	if err != nil {
		return UniqueFileID{}, err
	}

	rld := internal.RLEDecode(bs)

	uid := UniqueFileID{}
	uid.TypeID, rld = binary.LittleEndian.Uint32(rld[0:4]), rld[4:]

	if len(rld) == 8 {
		// Any other document
		uid.ID, rld = binary.LittleEndian.Uint64(rld), rld[8:]
	}

	// TODO: other types
	// TODO: len check
	return uid, nil
}
