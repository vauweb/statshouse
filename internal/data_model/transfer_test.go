// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package data_model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"

	"github.com/vkcom/statshouse/internal/data_model/gen2/tlstatshouse"
	"github.com/vkcom/statshouse/internal/format"
)

const (
	maxSTagLength = 128             // Maximum length for string tags
	now           = 3600 * 24 * 365 // 1 year in seconds
)

func getActualSize(t require.TestingT, key Key, defaultTimestamp uint32) int {
	item := key.TLMultiItemFromKey(defaultTimestamp)
	buf := make([]byte, 0, 1024) // Initial buffer
	buf, err := item.WriteGeneral(buf)
	require.NoError(t, err)
	return len(buf)
}

// Custom generators for Key components
func genKey() *rapid.Generator[Key] {
	return rapid.Custom(func(t *rapid.T) Key {
		key := Key{
			Timestamp: uint32(rapid.Int64Range(1, now).Draw(t, "timestamp")),
			Metric:    int32(rapid.Int64Range(0, 1000).Draw(t, "metric")),
		}

		// Generate random tags
		for i := range key.Tags {
			key.Tags[i] = int32(rapid.Int64Range(0, 1000).Draw(t, "tag"))
		}

		// Generate STags with different patterns
		pattern := rapid.IntRange(0, 3).Draw(t, "stags_pattern")
		switch pattern {
		case 0: // All empty
			// STags already initialized as empty strings
		case 1: // One non-empty
			idx := rapid.IntRange(0, format.MaxTags-1).Draw(t, "stag_index")
			key.SetSTag(idx, rapid.StringN(1, 20, maxSTagLength).Draw(t, "stag"))
		case 2: // Half filled
			for i := 0; i < format.MaxTags/2; i++ {
				key.SetSTag(i, rapid.StringN(1, 20, maxSTagLength).Draw(t, "stag"))
			}
		case 3: // All filled
			for i := range key.Tags {
				key.SetSTag(i, rapid.StringN(1, 20, maxSTagLength).Draw(t, "stag"))
			}
		}

		return key
	})
}

func TestKeySizeEstimationProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		key := genKey().Draw(t, "key")
		defaultTimestamp := uint32(time.Now().Unix())

		estimated := key.TLSizeEstimate(defaultTimestamp)
		actual := getActualSize(t, key, defaultTimestamp)

		// The estimated size should never be less than actual
		if estimated < actual {
			t.Fatalf("Estimated size %d is less than actual size %d for key: %+v",
				estimated, actual, key)
		}

		// The estimation shouldn't be too far off (within 20%)
		if float64(estimated) > float64(actual)*1.2 {
			t.Fatalf("Estimated size %d is too large compared to actual size %d for key: %+v",
				estimated, actual, key)
		}
	})
}

func TestKeySizeEstimationEdgeCases(t *testing.T) {
	testCases := []struct {
		name string
		key  Key
	}{
		{
			name: "Empty key",
			key:  Key{},
		},
		{
			name: "Only timestamp",
			key: Key{
				Timestamp: 12345,
			},
		},
		{
			name: "Only metric",
			key: Key{
				Metric: 67890,
			},
		},
		{
			name: "Single tag",
			key: Key{
				Tags: [format.MaxTags]int32{42},
			},
		},
		{
			name: "Single STag",
			key: Key{
				sTags: &sTagsHolder{[format.MaxTags]string{"test"}},
			},
		},
		{
			name: "Single STag max length",
			key: Key{
				sTags: &sTagsHolder{[format.MaxTags]string{strings.Repeat("a", maxSTagLength)}},
			},
		},
		{
			name: "All fields filled with max length STags",
			key: Key{
				Timestamp: 12345,
				Metric:    67890,
				Tags:      [format.MaxTags]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				sTags: func() *sTagsHolder {
					var stags sTagsHolder
					for i := 0; i < format.MaxTags; i++ {
						stags.values[i] = strings.Repeat("a", maxSTagLength)
					}
					return &stags
				}(),
			},
		},
		{
			name: "Timestamp equals defaultTimestamp",
			key: Key{
				Timestamp: 12345, // Will be passed as defaultTimestamp in test
				Metric:    67890,
			},
		},
		{
			name: "Mixed length STags",
			key: Key{
				sTags: func() *sTagsHolder {
					var stags sTagsHolder
					for i := 0; i < format.MaxTags; i++ {
						stags.values[i] = strings.Repeat("a", i)
					}
					return &stags
				}(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defaultTimestamp := tc.key.Timestamp
			estimated := tc.key.TLSizeEstimate(defaultTimestamp)
			actual := getActualSize(t, tc.key, defaultTimestamp)

			require.Equal(t, estimated, actual,
				"Estimated size should be equal to actual size\nKey: %+v\nEstimated: %d\nActual: %d",
				tc.key, estimated, actual)
		})
	}
}

// Helper function to convert Key to MultiItemBytes and back
func roundTripKey(key Key, bucketTimestamp uint32, newestTime uint32) Key {
	// Convert Key to MultiItem
	item := key.TLMultiItemFromKey(bucketTimestamp)

	// Convert MultiItem to MultiItemBytes
	multiItemBytes := &tlstatshouse.MultiItemBytes{
		Metric:     item.Metric,
		Keys:       item.Keys,
		FieldsMask: item.FieldsMask,
	}
	if item.IsSetT() {
		multiItemBytes.SetT(item.T)
	}
	if item.IsSetSkeys() {
		skeysBytes := make([][]byte, 0, len(item.Skeys))
		for _, skey := range item.Skeys {
			skeysBytes = append(skeysBytes, []byte(skey))
		}
		multiItemBytes.SetSkeys(skeysBytes)
	}

	// Convert MultiItemBytes back to Key
	reconstructedKey, _ := KeyFromStatshouseMultiItem(multiItemBytes, bucketTimestamp, newestTime)
	return reconstructedKey
}

func timestampValid(t require.TestingT, originalTs, newestTime, reconstructedTs, bucketTimestamp uint32) {
	var oldestTime uint32
	if bucketTimestamp > BelieveTimestampWindow {
		oldestTime = bucketTimestamp - BelieveTimestampWindow
	}
	switch {
	case originalTs == 0:
		require.Equal(t, originalTs, reconstructedTs, "Zero timestamp should be replaced with bucketTimestamp")
	case originalTs > newestTime:
		require.Equal(t, newestTime, reconstructedTs, "Timestamp should be clamped to newestTime")
	case originalTs < oldestTime:
		require.Equal(t, oldestTime, reconstructedTs,
			"Timestamp should be clamped to bucketTimestamp-BelieveTimestampWindow")
	default:
		require.Equal(t, originalTs, reconstructedTs, "Timestamps should match")
	}
}

func TestKeyFromStatshouseMultiItem(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate timestamps within reasonable bounds
		bucketTimestamp := uint32(rapid.Int64Range(1, now).Draw(t, "bucket_timestamp"))
		newestTime := bucketTimestamp + uint32(rapid.Int64Range(0, 3600*24).Draw(t, "time_offset")) // Up to 24 hours newer

		// Generate a random key using existing generator
		originalKey := genKey().Draw(t, "key")

		// Perform roundtrip
		reconstructedKey := roundTripKey(originalKey, bucketTimestamp, newestTime)

		// Verify key components
		require.Equal(t, originalKey.Metric, reconstructedKey.Metric, "Metrics should match")
		require.Equal(t, originalKey.Tags, reconstructedKey.Tags, "Tags should match")
		require.Equal(t, originalKey.sTags, reconstructedKey.sTags, "STags should match")
		timestampValid(t, originalKey.Timestamp, newestTime, reconstructedKey.Timestamp, bucketTimestamp)
	})
}