package main

import (
	"encoding/json"
	"fmt"
)

func prepareCheckinBody(aoIncluded bool) []byte {
	payloadJson := map[string]any{}
	ephemeral := "" //ephemeral--
	if len(ephemeral) > 0 {
		payloadJson["ephemeral"] = ephemeral
	}
	payloadJson["ztna_checksum"] = 1

	payloadJsonBytes, err := json.Marshal(payloadJson) // payload
	if err != nil {
		fmt.Println(err)
	}

	return payloadJsonBytes
}

func main() {
	payload := prepareCheckinBody(false)
	fmt.Println(string(payload))

}
