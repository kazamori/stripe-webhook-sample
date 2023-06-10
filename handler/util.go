package handler

import (
	"bytes"
	"encoding/json"
	"fmt"

	"golang.org/x/exp/slog"
)

func prettyPrint(v any) {
	var buf bytes.Buffer
	data := mustMarshal(v)
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		slog.Error("dump", err, "data", string(data))
		return
	}
	fmt.Println(string(buf.Bytes())) // pretty print
}

func mustMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
