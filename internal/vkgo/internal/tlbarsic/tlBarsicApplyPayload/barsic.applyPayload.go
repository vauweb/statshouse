// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBarsicApplyPayload

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/internal"
	"github.com/vkcom/statshouse/internal/vkgo/internal/tl/tlTrue"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type BarsicApplyPayload struct {
	FieldsMask uint32
	// CommitASAP (TrueType) // Conditional: item.FieldsMask.0
	// Committed (TrueType) // Conditional: item.FieldsMask.1
	Offset  int64
	Payload string
}

func (BarsicApplyPayload) TLName() string { return "barsic.applyPayload" }
func (BarsicApplyPayload) TLTag() uint32  { return 0x8c981e13 }

func (item *BarsicApplyPayload) SetCommitASAP(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item BarsicApplyPayload) IsSetCommitASAP() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *BarsicApplyPayload) SetCommitted(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item BarsicApplyPayload) IsSetCommitted() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *BarsicApplyPayload) Reset() {
	item.FieldsMask = 0
	item.Offset = 0
	item.Payload = ""
}

func (item *BarsicApplyPayload) FillRandom(rg *basictl.RandGenerator) {
	var maskFieldsMask uint32
	maskFieldsMask = basictl.RandomUint(rg)
	item.FieldsMask = 0
	if maskFieldsMask&(1<<0) != 0 {
		item.FieldsMask |= (1 << 0)
	}
	if maskFieldsMask&(1<<1) != 0 {
		item.FieldsMask |= (1 << 1)
	}
	item.Offset = basictl.RandomLong(rg)
	item.Payload = basictl.RandomString(rg)
}

func (item *BarsicApplyPayload) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Payload)
}

// This method is general version of Write, use it instead!
func (item *BarsicApplyPayload) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *BarsicApplyPayload) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.Offset)
	w = basictl.StringWrite(w, item.Payload)
	return w
}

func (item *BarsicApplyPayload) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8c981e13); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BarsicApplyPayload) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *BarsicApplyPayload) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x8c981e13)
	return item.Write(w)
}

func (item *BarsicApplyPayload) ReadResult(w []byte, ret *tlTrue.True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *BarsicApplyPayload) WriteResult(w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *BarsicApplyPayload) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlTrue.True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *BarsicApplyPayload) WriteResultJSON(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BarsicApplyPayload) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *BarsicApplyPayload) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BarsicApplyPayload) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BarsicApplyPayload) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlTrue.True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item BarsicApplyPayload) String() string {
	return string(item.WriteJSON(nil))
}

func (item *BarsicApplyPayload) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var trueTypeCommitASAPPresented bool
	var trueTypeCommitASAPValue bool
	var trueTypeCommittedPresented bool
	var trueTypeCommittedValue bool
	var propOffsetPresented bool
	var propPayloadPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "commitASAP":
				if trueTypeCommitASAPPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "commitASAP")
				}
				if err := internal.Json2ReadBool(in, &trueTypeCommitASAPValue); err != nil {
					return err
				}
				trueTypeCommitASAPPresented = true
			case "committed":
				if trueTypeCommittedPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "committed")
				}
				if err := internal.Json2ReadBool(in, &trueTypeCommittedValue); err != nil {
					return err
				}
				trueTypeCommittedPresented = true
			case "offset":
				if propOffsetPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "offset")
				}
				if err := internal.Json2ReadInt64(in, &item.Offset); err != nil {
					return err
				}
				propOffsetPresented = true
			case "payload":
				if propPayloadPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "payload")
				}
				if err := internal.Json2ReadString(in, &item.Payload); err != nil {
					return err
				}
				propPayloadPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("barsic.applyPayload", key)
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
	if !propPayloadPresented {
		item.Payload = ""
	}
	if trueTypeCommitASAPPresented {
		if trueTypeCommitASAPValue {
			item.FieldsMask |= 1 << 0
		}
	}
	if trueTypeCommittedPresented {
		if trueTypeCommittedValue {
			item.FieldsMask |= 1 << 1
		}
	}
	// tries to set bit to zero if it is 1
	if trueTypeCommitASAPPresented && !trueTypeCommitASAPValue && (item.FieldsMask&(1<<0) != 0) {
		return internal.ErrorInvalidJSON("barsic.applyPayload", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	// tries to set bit to zero if it is 1
	if trueTypeCommittedPresented && !trueTypeCommittedValue && (item.FieldsMask&(1<<1) != 0) {
		return internal.ErrorInvalidJSON("barsic.applyPayload", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BarsicApplyPayload) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *BarsicApplyPayload) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BarsicApplyPayload) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"commitASAP":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"committed":true`...)
	}
	backupIndexOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"offset":`...)
	w = basictl.JSONWriteInt64(w, item.Offset)
	if (item.Offset != 0) == false {
		w = w[:backupIndexOffset]
	}
	backupIndexPayload := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"payload":`...)
	w = basictl.JSONWriteString(w, item.Payload)
	if (len(item.Payload) != 0) == false {
		w = w[:backupIndexPayload]
	}
	return append(w, '}')
}

func (item *BarsicApplyPayload) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *BarsicApplyPayload) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	return nil
}

type BarsicApplyPayloadBytes struct {
	FieldsMask uint32
	// CommitASAP (TrueType) // Conditional: item.FieldsMask.0
	// Committed (TrueType) // Conditional: item.FieldsMask.1
	Offset  int64
	Payload []byte
}

func (BarsicApplyPayloadBytes) TLName() string { return "barsic.applyPayload" }
func (BarsicApplyPayloadBytes) TLTag() uint32  { return 0x8c981e13 }

func (item *BarsicApplyPayloadBytes) SetCommitASAP(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item BarsicApplyPayloadBytes) IsSetCommitASAP() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *BarsicApplyPayloadBytes) SetCommitted(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item BarsicApplyPayloadBytes) IsSetCommitted() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *BarsicApplyPayloadBytes) Reset() {
	item.FieldsMask = 0
	item.Offset = 0
	item.Payload = item.Payload[:0]
}

func (item *BarsicApplyPayloadBytes) FillRandom(rg *basictl.RandGenerator) {
	var maskFieldsMask uint32
	maskFieldsMask = basictl.RandomUint(rg)
	item.FieldsMask = 0
	if maskFieldsMask&(1<<0) != 0 {
		item.FieldsMask |= (1 << 0)
	}
	if maskFieldsMask&(1<<1) != 0 {
		item.FieldsMask |= (1 << 1)
	}
	item.Offset = basictl.RandomLong(rg)
	item.Payload = basictl.RandomStringBytes(rg)
}

func (item *BarsicApplyPayloadBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.Payload)
}

// This method is general version of Write, use it instead!
func (item *BarsicApplyPayloadBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *BarsicApplyPayloadBytes) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.Offset)
	w = basictl.StringWriteBytes(w, item.Payload)
	return w
}

func (item *BarsicApplyPayloadBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8c981e13); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BarsicApplyPayloadBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *BarsicApplyPayloadBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x8c981e13)
	return item.Write(w)
}

func (item *BarsicApplyPayloadBytes) ReadResult(w []byte, ret *tlTrue.True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *BarsicApplyPayloadBytes) WriteResult(w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *BarsicApplyPayloadBytes) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlTrue.True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *BarsicApplyPayloadBytes) WriteResultJSON(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BarsicApplyPayloadBytes) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *BarsicApplyPayloadBytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BarsicApplyPayloadBytes) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BarsicApplyPayloadBytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlTrue.True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item BarsicApplyPayloadBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *BarsicApplyPayloadBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var trueTypeCommitASAPPresented bool
	var trueTypeCommitASAPValue bool
	var trueTypeCommittedPresented bool
	var trueTypeCommittedValue bool
	var propOffsetPresented bool
	var propPayloadPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "commitASAP":
				if trueTypeCommitASAPPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "commitASAP")
				}
				if err := internal.Json2ReadBool(in, &trueTypeCommitASAPValue); err != nil {
					return err
				}
				trueTypeCommitASAPPresented = true
			case "committed":
				if trueTypeCommittedPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "committed")
				}
				if err := internal.Json2ReadBool(in, &trueTypeCommittedValue); err != nil {
					return err
				}
				trueTypeCommittedPresented = true
			case "offset":
				if propOffsetPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "offset")
				}
				if err := internal.Json2ReadInt64(in, &item.Offset); err != nil {
					return err
				}
				propOffsetPresented = true
			case "payload":
				if propPayloadPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("barsic.applyPayload", "payload")
				}
				if err := internal.Json2ReadStringBytes(in, &item.Payload); err != nil {
					return err
				}
				propPayloadPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("barsic.applyPayload", key)
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
	if !propPayloadPresented {
		item.Payload = item.Payload[:0]
	}
	if trueTypeCommitASAPPresented {
		if trueTypeCommitASAPValue {
			item.FieldsMask |= 1 << 0
		}
	}
	if trueTypeCommittedPresented {
		if trueTypeCommittedValue {
			item.FieldsMask |= 1 << 1
		}
	}
	// tries to set bit to zero if it is 1
	if trueTypeCommitASAPPresented && !trueTypeCommitASAPValue && (item.FieldsMask&(1<<0) != 0) {
		return internal.ErrorInvalidJSON("barsic.applyPayload", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	// tries to set bit to zero if it is 1
	if trueTypeCommittedPresented && !trueTypeCommittedValue && (item.FieldsMask&(1<<1) != 0) {
		return internal.ErrorInvalidJSON("barsic.applyPayload", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BarsicApplyPayloadBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *BarsicApplyPayloadBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BarsicApplyPayloadBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"commitASAP":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"committed":true`...)
	}
	backupIndexOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"offset":`...)
	w = basictl.JSONWriteInt64(w, item.Offset)
	if (item.Offset != 0) == false {
		w = w[:backupIndexOffset]
	}
	backupIndexPayload := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"payload":`...)
	w = basictl.JSONWriteStringBytes(w, item.Payload)
	if (len(item.Payload) != 0) == false {
		w = w[:backupIndexPayload]
	}
	return append(w, '}')
}

func (item *BarsicApplyPayloadBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *BarsicApplyPayloadBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	return nil
}