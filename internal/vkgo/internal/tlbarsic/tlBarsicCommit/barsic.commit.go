// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBarsicCommit

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/internal"
	"github.com/vkcom/statshouse/internal/vkgo/internal/tl/tlTrue"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type BarsicCommit struct {
	FieldsMask         uint32
	Offset             int64
	SnapshotMeta       string
	SafeSnapshotOffset int64
}

func (BarsicCommit) TLName() string { return "barsic.commit" }
func (BarsicCommit) TLTag() uint32  { return 0x12357324 }

func (item *BarsicCommit) Reset() {
	item.FieldsMask = 0
	item.Offset = 0
	item.SnapshotMeta = ""
	item.SafeSnapshotOffset = 0
}

func (item *BarsicCommit) FillRandom(rg *basictl.RandGenerator) {
	item.Offset = basictl.RandomLong(rg)
	item.SnapshotMeta = basictl.RandomString(rg)
	item.SafeSnapshotOffset = basictl.RandomLong(rg)
}

func (item *BarsicCommit) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.SnapshotMeta); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.SafeSnapshotOffset)
}

// This method is general version of Write, use it instead!
func (item *BarsicCommit) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *BarsicCommit) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.Offset)
	w = basictl.StringWrite(w, item.SnapshotMeta)
	w = basictl.LongWrite(w, item.SafeSnapshotOffset)
	return w
}

func (item *BarsicCommit) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x12357324); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BarsicCommit) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *BarsicCommit) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x12357324)
	return item.Write(w)
}

func (item *BarsicCommit) ReadResult(w []byte, ret *tlTrue.True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *BarsicCommit) WriteResult(w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *BarsicCommit) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlTrue.True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *BarsicCommit) WriteResultJSON(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BarsicCommit) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *BarsicCommit) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BarsicCommit) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BarsicCommit) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlTrue.True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item BarsicCommit) String() string {
	return string(item.WriteJSON(nil))
}

func (item *BarsicCommit) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propOffsetPresented bool
	var propSnapshotMetaPresented bool
	var propSafeSnapshotOffsetPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "offset":
				if propOffsetPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "offset")
				}
				if err := internal.Json2ReadInt64(in, &item.Offset); err != nil {
					return err
				}
				propOffsetPresented = true
			case "snapshot_meta":
				if propSnapshotMetaPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "snapshot_meta")
				}
				if err := internal.Json2ReadString(in, &item.SnapshotMeta); err != nil {
					return err
				}
				propSnapshotMetaPresented = true
			case "safe_snapshot_offset":
				if propSafeSnapshotOffsetPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "safe_snapshot_offset")
				}
				if err := internal.Json2ReadInt64(in, &item.SafeSnapshotOffset); err != nil {
					return err
				}
				propSafeSnapshotOffsetPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("barsic.commit", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propOffsetPresented {
		item.Offset = 0
	}
	if !propSnapshotMetaPresented {
		item.SnapshotMeta = ""
	}
	if !propSafeSnapshotOffsetPresented {
		item.SafeSnapshotOffset = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BarsicCommit) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *BarsicCommit) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BarsicCommit) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"offset":`...)
	w = basictl.JSONWriteInt64(w, item.Offset)
	if (item.Offset != 0) == false {
		w = w[:backupIndexOffset]
	}
	backupIndexSnapshotMeta := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"snapshot_meta":`...)
	w = basictl.JSONWriteString(w, item.SnapshotMeta)
	if (len(item.SnapshotMeta) != 0) == false {
		w = w[:backupIndexSnapshotMeta]
	}
	backupIndexSafeSnapshotOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"safe_snapshot_offset":`...)
	w = basictl.JSONWriteInt64(w, item.SafeSnapshotOffset)
	if (item.SafeSnapshotOffset != 0) == false {
		w = w[:backupIndexSafeSnapshotOffset]
	}
	return append(w, '}')
}

func (item *BarsicCommit) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *BarsicCommit) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("barsic.commit", err.Error())
	}
	return nil
}

type BarsicCommitBytes struct {
	FieldsMask         uint32
	Offset             int64
	SnapshotMeta       []byte
	SafeSnapshotOffset int64
}

func (BarsicCommitBytes) TLName() string { return "barsic.commit" }
func (BarsicCommitBytes) TLTag() uint32  { return 0x12357324 }

func (item *BarsicCommitBytes) Reset() {
	item.FieldsMask = 0
	item.Offset = 0
	item.SnapshotMeta = item.SnapshotMeta[:0]
	item.SafeSnapshotOffset = 0
}

func (item *BarsicCommitBytes) FillRandom(rg *basictl.RandGenerator) {
	item.Offset = basictl.RandomLong(rg)
	item.SnapshotMeta = basictl.RandomStringBytes(rg)
	item.SafeSnapshotOffset = basictl.RandomLong(rg)
}

func (item *BarsicCommitBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	if w, err = basictl.StringReadBytes(w, &item.SnapshotMeta); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.SafeSnapshotOffset)
}

// This method is general version of Write, use it instead!
func (item *BarsicCommitBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *BarsicCommitBytes) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.Offset)
	w = basictl.StringWriteBytes(w, item.SnapshotMeta)
	w = basictl.LongWrite(w, item.SafeSnapshotOffset)
	return w
}

func (item *BarsicCommitBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x12357324); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BarsicCommitBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *BarsicCommitBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x12357324)
	return item.Write(w)
}

func (item *BarsicCommitBytes) ReadResult(w []byte, ret *tlTrue.True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *BarsicCommitBytes) WriteResult(w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *BarsicCommitBytes) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlTrue.True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *BarsicCommitBytes) WriteResultJSON(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BarsicCommitBytes) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *BarsicCommitBytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BarsicCommitBytes) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BarsicCommitBytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlTrue.True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item BarsicCommitBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *BarsicCommitBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propOffsetPresented bool
	var propSnapshotMetaPresented bool
	var propSafeSnapshotOffsetPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "offset":
				if propOffsetPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "offset")
				}
				if err := internal.Json2ReadInt64(in, &item.Offset); err != nil {
					return err
				}
				propOffsetPresented = true
			case "snapshot_meta":
				if propSnapshotMetaPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "snapshot_meta")
				}
				if err := internal.Json2ReadStringBytes(in, &item.SnapshotMeta); err != nil {
					return err
				}
				propSnapshotMetaPresented = true
			case "safe_snapshot_offset":
				if propSafeSnapshotOffsetPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.commit", "safe_snapshot_offset")
				}
				if err := internal.Json2ReadInt64(in, &item.SafeSnapshotOffset); err != nil {
					return err
				}
				propSafeSnapshotOffsetPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("barsic.commit", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propOffsetPresented {
		item.Offset = 0
	}
	if !propSnapshotMetaPresented {
		item.SnapshotMeta = item.SnapshotMeta[:0]
	}
	if !propSafeSnapshotOffsetPresented {
		item.SafeSnapshotOffset = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BarsicCommitBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *BarsicCommitBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BarsicCommitBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"offset":`...)
	w = basictl.JSONWriteInt64(w, item.Offset)
	if (item.Offset != 0) == false {
		w = w[:backupIndexOffset]
	}
	backupIndexSnapshotMeta := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"snapshot_meta":`...)
	w = basictl.JSONWriteStringBytes(w, item.SnapshotMeta)
	if (len(item.SnapshotMeta) != 0) == false {
		w = w[:backupIndexSnapshotMeta]
	}
	backupIndexSafeSnapshotOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"safe_snapshot_offset":`...)
	w = basictl.JSONWriteInt64(w, item.SafeSnapshotOffset)
	if (item.SafeSnapshotOffset != 0) == false {
		w = w[:backupIndexSafeSnapshotOffset]
	}
	return append(w, '}')
}

func (item *BarsicCommitBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *BarsicCommitBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("barsic.commit", err.Error())
	}
	return nil
}