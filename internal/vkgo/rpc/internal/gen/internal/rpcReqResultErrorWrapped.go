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

type RpcReqResultErrorWrapped struct {
	ErrorCode int32
	Error     string
}

func (RpcReqResultErrorWrapped) TLName() string { return "rpcReqResultErrorWrapped" }
func (RpcReqResultErrorWrapped) TLTag() uint32  { return 0x7ae432f6 }

func (item *RpcReqResultErrorWrapped) Reset() {
	item.ErrorCode = 0
	item.Error = ""
}

func (item *RpcReqResultErrorWrapped) FillRandom(rg *basictl.RandGenerator) {
	item.ErrorCode = basictl.RandomInt(rg)
	item.Error = basictl.RandomString(rg)
}

func (item *RpcReqResultErrorWrapped) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.ErrorCode); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Error)
}

// This method is general version of Write, use it instead!
func (item *RpcReqResultErrorWrapped) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *RpcReqResultErrorWrapped) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.ErrorCode)
	w = basictl.StringWrite(w, item.Error)
	return w
}

func (item *RpcReqResultErrorWrapped) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x7ae432f6); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *RpcReqResultErrorWrapped) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *RpcReqResultErrorWrapped) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x7ae432f6)
	return item.Write(w)
}

func (item RpcReqResultErrorWrapped) String() string {
	return string(item.WriteJSON(nil))
}

func (item *RpcReqResultErrorWrapped) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propErrorCodePresented bool
	var propErrorPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "error_code":
				if propErrorCodePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("rpcReqResultErrorWrapped", "error_code")
				}
				if err := Json2ReadInt32(in, &item.ErrorCode); err != nil {
					return err
				}
				propErrorCodePresented = true
			case "error":
				if propErrorPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("rpcReqResultErrorWrapped", "error")
				}
				if err := Json2ReadString(in, &item.Error); err != nil {
					return err
				}
				propErrorPresented = true
			default:
				return ErrorInvalidJSONExcessElement("rpcReqResultErrorWrapped", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propErrorCodePresented {
		item.ErrorCode = 0
	}
	if !propErrorPresented {
		item.Error = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *RpcReqResultErrorWrapped) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *RpcReqResultErrorWrapped) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *RpcReqResultErrorWrapped) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexErrorCode := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"error_code":`...)
	w = basictl.JSONWriteInt32(w, item.ErrorCode)
	if (item.ErrorCode != 0) == false {
		w = w[:backupIndexErrorCode]
	}
	backupIndexError := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"error":`...)
	w = basictl.JSONWriteString(w, item.Error)
	if (len(item.Error) != 0) == false {
		w = w[:backupIndexError]
	}
	return append(w, '}')
}

func (item *RpcReqResultErrorWrapped) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *RpcReqResultErrorWrapped) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("rpcReqResultErrorWrapped", err.Error())
	}
	return nil
}