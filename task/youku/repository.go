package youku

import (
	"encoding/json"

	"github.com/dcb9/steamer/task/common"
)

var r = common.ResourceInfoRepository{GenerateStreams}

func GenerateStreams(streams []byte) map[string]common.Stream {
	s := make(map[string]YoukuStream)
	json.Unmarshal([]byte(streams), &s)

	streamsMap := make(map[string]common.Stream)
	for key, val := range s {
		streamsMap[key] = val
	}

	return streamsMap
}
