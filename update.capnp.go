package main

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
	"unsafe"
)

type ValidateFile C.Struct

func NewValidateFile(s *C.Segment) ValidateFile      { return ValidateFile(s.NewStruct(8, 3)) }
func NewRootValidateFile(s *C.Segment) ValidateFile  { return ValidateFile(s.NewRootStruct(8, 3)) }
func AutoNewValidateFile(s *C.Segment) ValidateFile  { return ValidateFile(s.NewStructAR(8, 3)) }
func ReadRootValidateFile(s *C.Segment) ValidateFile { return ValidateFile(s.Root(0).ToStruct()) }
func (s ValidateFile) Name() string                  { return C.Struct(s).GetObject(0).ToText() }
func (s ValidateFile) SetName(v string)              { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s ValidateFile) Hash() []byte                  { return C.Struct(s).GetObject(1).ToData() }
func (s ValidateFile) SetHash(v []byte)              { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s ValidateFile) Buildnum() uint8               { return C.Struct(s).Get8(0) }
func (s ValidateFile) SetBuildnum(v uint8)           { C.Struct(s).Set8(0, v) }
func (s ValidateFile) Sign() []byte                  { return C.Struct(s).GetObject(2).ToData() }
func (s ValidateFile) SetSign(v []byte)              { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ValidateFile) MarshalJSON() (bs []byte, err error) { return }

type ValidateFile_List C.PointerList

func NewValidateFileList(s *C.Segment, sz int) ValidateFile_List {
	return ValidateFile_List(s.NewCompositeList(8, 3, sz))
}
func (s ValidateFile_List) Len() int { return C.PointerList(s).Len() }
func (s ValidateFile_List) At(i int) ValidateFile {
	return ValidateFile(C.PointerList(s).At(i).ToStruct())
}
func (s ValidateFile_List) ToArray() []ValidateFile {
	return *(*[]ValidateFile)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s ValidateFile_List) Set(i int, item ValidateFile) { C.PointerList(s).Set(i, C.Object(item)) }

type FileControl C.Struct

func NewFileControl(s *C.Segment) FileControl      { return FileControl(s.NewStruct(8, 3)) }
func NewRootFileControl(s *C.Segment) FileControl  { return FileControl(s.NewRootStruct(8, 3)) }
func AutoNewFileControl(s *C.Segment) FileControl  { return FileControl(s.NewStructAR(8, 3)) }
func ReadRootFileControl(s *C.Segment) FileControl { return FileControl(s.Root(0).ToStruct()) }
func (s FileControl) Name() string                 { return C.Struct(s).GetObject(0).ToText() }
func (s FileControl) SetName(v string)             { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s FileControl) Hashnext() []byte             { return C.Struct(s).GetObject(1).ToData() }
func (s FileControl) SetHashnext(v []byte)         { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s FileControl) Chunknum() uint8              { return C.Struct(s).Get8(0) }
func (s FileControl) SetChunknum(v uint8)          { C.Struct(s).Set8(0, v) }
func (s FileControl) Chunksize() uint32            { return C.Struct(s).Get32(4) ^ 1024 }
func (s FileControl) SetChunksize(v uint32)        { C.Struct(s).Set32(4, v^1024) }
func (s FileControl) Sign() []byte                 { return C.Struct(s).GetObject(2).ToData() }
func (s FileControl) SetSign(v []byte)             { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s FileControl) MarshalJSON() (bs []byte, err error) { return }

type FileControl_List C.PointerList

func NewFileControlList(s *C.Segment, sz int) FileControl_List {
	return FileControl_List(s.NewCompositeList(8, 3, sz))
}
func (s FileControl_List) Len() int             { return C.PointerList(s).Len() }
func (s FileControl_List) At(i int) FileControl { return FileControl(C.PointerList(s).At(i).ToStruct()) }
func (s FileControl_List) ToArray() []FileControl {
	return *(*[]FileControl)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s FileControl_List) Set(i int, item FileControl) { C.PointerList(s).Set(i, C.Object(item)) }

type FileChunk C.Struct

func NewFileChunk(s *C.Segment) FileChunk      { return FileChunk(s.NewStruct(8, 2)) }
func NewRootFileChunk(s *C.Segment) FileChunk  { return FileChunk(s.NewRootStruct(8, 2)) }
func AutoNewFileChunk(s *C.Segment) FileChunk  { return FileChunk(s.NewStructAR(8, 2)) }
func ReadRootFileChunk(s *C.Segment) FileChunk { return FileChunk(s.Root(0).ToStruct()) }
func (s FileChunk) Chunknum() uint8            { return C.Struct(s).Get8(0) }
func (s FileChunk) SetChunknum(v uint8)        { C.Struct(s).Set8(0, v) }
func (s FileChunk) Data() []byte               { return C.Struct(s).GetObject(0).ToData() }
func (s FileChunk) SetData(v []byte)           { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s FileChunk) Hashnext() []byte           { return C.Struct(s).GetObject(1).ToData() }
func (s FileChunk) SetHashnext(v []byte)       { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s FileChunk) MarshalJSON() (bs []byte, err error) { return }

type FileChunk_List C.PointerList

func NewFileChunkList(s *C.Segment, sz int) FileChunk_List {
	return FileChunk_List(s.NewCompositeList(8, 2, sz))
}
func (s FileChunk_List) Len() int           { return C.PointerList(s).Len() }
func (s FileChunk_List) At(i int) FileChunk { return FileChunk(C.PointerList(s).At(i).ToStruct()) }
func (s FileChunk_List) ToArray() []FileChunk {
	return *(*[]FileChunk)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s FileChunk_List) Set(i int, item FileChunk) { C.PointerList(s).Set(i, C.Object(item)) }

type GetManifest C.Struct

func NewGetManifest(s *C.Segment) GetManifest      { return GetManifest(s.NewStruct(8, 0)) }
func NewRootGetManifest(s *C.Segment) GetManifest  { return GetManifest(s.NewRootStruct(8, 0)) }
func AutoNewGetManifest(s *C.Segment) GetManifest  { return GetManifest(s.NewStructAR(8, 0)) }
func ReadRootGetManifest(s *C.Segment) GetManifest { return GetManifest(s.Root(0).ToStruct()) }
func (s GetManifest) Buildnum() uint8              { return C.Struct(s).Get8(0) }
func (s GetManifest) SetBuildnum(v uint8)          { C.Struct(s).Set8(0, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s GetManifest) MarshalJSON() (bs []byte, err error) { return }

type GetManifest_List C.PointerList

func NewGetManifestList(s *C.Segment, sz int) GetManifest_List {
	return GetManifest_List(s.NewUInt8List(sz))
}
func (s GetManifest_List) Len() int             { return C.PointerList(s).Len() }
func (s GetManifest_List) At(i int) GetManifest { return GetManifest(C.PointerList(s).At(i).ToStruct()) }
func (s GetManifest_List) ToArray() []GetManifest {
	return *(*[]GetManifest)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s GetManifest_List) Set(i int, item GetManifest) { C.PointerList(s).Set(i, C.Object(item)) }

type SendManifest C.Struct

func NewSendManifest(s *C.Segment) SendManifest      { return SendManifest(s.NewStruct(8, 2)) }
func NewRootSendManifest(s *C.Segment) SendManifest  { return SendManifest(s.NewRootStruct(8, 2)) }
func AutoNewSendManifest(s *C.Segment) SendManifest  { return SendManifest(s.NewStructAR(8, 2)) }
func ReadRootSendManifest(s *C.Segment) SendManifest { return SendManifest(s.Root(0).ToStruct()) }
func (s SendManifest) Buildnum() uint8               { return C.Struct(s).Get8(0) }
func (s SendManifest) SetBuildnum(v uint8)           { C.Struct(s).Set8(0, v) }
func (s SendManifest) Data() []byte                  { return C.Struct(s).GetObject(0).ToData() }
func (s SendManifest) SetData(v []byte)              { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s SendManifest) Sign() []byte                  { return C.Struct(s).GetObject(1).ToData() }
func (s SendManifest) SetSign(v []byte)              { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s SendManifest) MarshalJSON() (bs []byte, err error) { return }

type SendManifest_List C.PointerList

func NewSendManifestList(s *C.Segment, sz int) SendManifest_List {
	return SendManifest_List(s.NewCompositeList(8, 2, sz))
}
func (s SendManifest_List) Len() int { return C.PointerList(s).Len() }
func (s SendManifest_List) At(i int) SendManifest {
	return SendManifest(C.PointerList(s).At(i).ToStruct())
}
func (s SendManifest_List) ToArray() []SendManifest {
	return *(*[]SendManifest)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s SendManifest_List) Set(i int, item SendManifest) { C.PointerList(s).Set(i, C.Object(item)) }
