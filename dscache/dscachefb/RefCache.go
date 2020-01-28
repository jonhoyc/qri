// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package dscachefb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RefCache struct {
	_tab flatbuffers.Table
}

func GetRootAsRefCache(buf []byte, offset flatbuffers.UOffsetT) *RefCache {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RefCache{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RefCache) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RefCache) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RefCache) InitID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RefCache) ProfileID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RefCache) TopIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefCache) MutateTopIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *RefCache) CursorIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefCache) MutateCursorIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *RefCache) PrettyName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RefCache) MetaTitle() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RefCache) ThemeList() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RefCache) BodySize() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefCache) MutateBodySize(n int64) bool {
	return rcv._tab.MutateInt64Slot(18, n)
}

func (rcv *RefCache) BodyRows() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefCache) MutateBodyRows(n int32) bool {
	return rcv._tab.MutateInt32Slot(20, n)
}

func (rcv *RefCache) CommitTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefCache) MutateCommitTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(22, n)
}

func (rcv *RefCache) NumErrors() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefCache) MutateNumErrors(n int32) bool {
	return rcv._tab.MutateInt32Slot(24, n)
}

func (rcv *RefCache) HeadRef() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RefCache) FsiPath() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func RefCacheStart(builder *flatbuffers.Builder) {
	builder.StartObject(13)
}
func RefCacheAddInitID(builder *flatbuffers.Builder, initID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(initID), 0)
}
func RefCacheAddProfileID(builder *flatbuffers.Builder, profileID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(profileID), 0)
}
func RefCacheAddTopIndex(builder *flatbuffers.Builder, topIndex int32) {
	builder.PrependInt32Slot(2, topIndex, 0)
}
func RefCacheAddCursorIndex(builder *flatbuffers.Builder, cursorIndex int32) {
	builder.PrependInt32Slot(3, cursorIndex, 0)
}
func RefCacheAddPrettyName(builder *flatbuffers.Builder, prettyName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(prettyName), 0)
}
func RefCacheAddMetaTitle(builder *flatbuffers.Builder, metaTitle flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(metaTitle), 0)
}
func RefCacheAddThemeList(builder *flatbuffers.Builder, themeList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(themeList), 0)
}
func RefCacheAddBodySize(builder *flatbuffers.Builder, bodySize int64) {
	builder.PrependInt64Slot(7, bodySize, 0)
}
func RefCacheAddBodyRows(builder *flatbuffers.Builder, bodyRows int32) {
	builder.PrependInt32Slot(8, bodyRows, 0)
}
func RefCacheAddCommitTime(builder *flatbuffers.Builder, commitTime int64) {
	builder.PrependInt64Slot(9, commitTime, 0)
}
func RefCacheAddNumErrors(builder *flatbuffers.Builder, numErrors int32) {
	builder.PrependInt32Slot(10, numErrors, 0)
}
func RefCacheAddHeadRef(builder *flatbuffers.Builder, headRef flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(headRef), 0)
}
func RefCacheAddFsiPath(builder *flatbuffers.Builder, fsiPath flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(fsiPath), 0)
}
func RefCacheEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
