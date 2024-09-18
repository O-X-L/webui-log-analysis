package u

import (
	"log"
)

func LogError(prefix string, err interface{}) {
	log.Fatalf("%v, Error | %v", prefix, err)
}
