package util

import (
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/parser/base"
)

// ChunkedSQLScript splits a SQL script into chunks.
func ChunkedSQLScript(script []base.SingleSQL, maxChunksCount int) ([][]base.SingleSQL, error) {
	var result [][]base.SingleSQL

	if maxChunksCount <= 0 {
		return nil, errors.New("maxChunksCount must be greater than 0")
	}

	if len(script) == 0 {
		return result, nil
	}

	roundDown := len(script) / maxChunksCount
	rest := len(script) % maxChunksCount

	// We have len(script) sqls, and we want to split it into no more than maxChunksCount chunks.
	// len(script) = roundDown * maxChunksCount + rest
	//             = rest * (roundDown + 1)  + (maxChunksCount - rest) * roundDown
	// So the first $rest chunks will have $(roundDown + 1) sqls, and the rest $(maxChunksCount) chunks will have $roundDown sqls.

	for i := 0; i < rest; i++ {
		start := i * (roundDown + 1)
		end := (i + 1) * (roundDown + 1)
		result = append(result, script[start:end])
	}

	for i := rest; i < maxChunksCount; i++ {
		start := rest*(roundDown+1) + (i-rest)*roundDown
		end := rest*(roundDown+1) + (i-rest+1)*roundDown
		result = append(result, script[start:end])
	}

	return result, nil
}
