// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

var _ = basictl.NatWrite

type ExactlyOnceSlotResponse struct {
	PersistentQueryUuid ExactlyOnceUuid
	PersistentSlotUuid  ExactlyOnceUuid
}

func (ExactlyOnceSlotResponse) TLName() string { return "exactlyOnce.slotResponse" }
func (ExactlyOnceSlotResponse) TLTag() uint32  { return 0x95f25c81 }

func (item *ExactlyOnceSlotResponse) Reset() {
	item.PersistentQueryUuid.Reset()
	item.PersistentSlotUuid.Reset()
}

func (item *ExactlyOnceSlotResponse) FillRandom(rg *basictl.RandGenerator) {
	item.PersistentQueryUuid.FillRandom(rg)
	item.PersistentSlotUuid.FillRandom(rg)
}

func (item *ExactlyOnceSlotResponse) Read(w []byte) (_ []byte, err error) {
	if w, err = item.PersistentQueryUuid.Read(w); err != nil {
		return w, err
	}
	return item.PersistentSlotUuid.Read(w)
}

// This method is general version of Write, use it instead!
func (item *ExactlyOnceSlotResponse) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *ExactlyOnceSlotResponse) Write(w []byte) []byte {
	w = item.PersistentQueryUuid.Write(w)
	w = item.PersistentSlotUuid.Write(w)
	return w
}

func (item *ExactlyOnceSlotResponse) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x95f25c81); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *ExactlyOnceSlotResponse) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *ExactlyOnceSlotResponse) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x95f25c81)
	return item.Write(w)
}

func (item ExactlyOnceSlotResponse) String() string {
	return string(item.WriteJSON(nil))
}

func (item *ExactlyOnceSlotResponse) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propPersistentQueryUuidPresented bool
	var propPersistentSlotUuidPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "persistent_query_uuid":
				if propPersistentQueryUuidPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("exactlyOnce.slotResponse", "persistent_query_uuid")
				}
				if err := item.PersistentQueryUuid.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propPersistentQueryUuidPresented = true
			case "persistent_slot_uuid":
				if propPersistentSlotUuidPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("exactlyOnce.slotResponse", "persistent_slot_uuid")
				}
				if err := item.PersistentSlotUuid.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propPersistentSlotUuidPresented = true
			default:
				return ErrorInvalidJSONExcessElement("exactlyOnce.slotResponse", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propPersistentQueryUuidPresented {
		item.PersistentQueryUuid.Reset()
	}
	if !propPersistentSlotUuidPresented {
		item.PersistentSlotUuid.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *ExactlyOnceSlotResponse) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *ExactlyOnceSlotResponse) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *ExactlyOnceSlotResponse) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"persistent_query_uuid":`...)
	w = item.PersistentQueryUuid.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"persistent_slot_uuid":`...)
	w = item.PersistentSlotUuid.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *ExactlyOnceSlotResponse) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *ExactlyOnceSlotResponse) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("exactlyOnce.slotResponse", err.Error())
	}
	return nil
}