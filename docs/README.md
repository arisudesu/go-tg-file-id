go-tg-file-id
=============
Decode Telegram Bot API file IDs, in Go.

Credits
-------
 * danog: for [original implementation in PHP](https://github.com/danog/tg-file-decoder).

Install
-------
```
go get github.com/arisudesu/go-tg-file-id
```

Usage
-----
```go
fileID, err := tg_file_id.DecodeFileID("CAACAgIAAxkBAAIEol9yQhBqFnT4HXldAh31a-hYXuDIAAIECwACAoujAAFFn1sl9AABHbkbBA")
// {Version:4
//  SubVersion:27
//  TypeID:8
//  DcID:2
//  FileReference:[1 0 0 4 162 95 114 66 16 106 22 116 248 29 121 93 2 29 245 107 232 88 94 224 200]
//  URL:[]
//  ID:46033261910035204
//  AccessHash:13338818719940058949}

uniqueID, err := tg_file_id.DecodeUniqueFileID("AgADBAsAAgKLowAB")
// {TypeID:2
//  ID:46033261910035204}
```
