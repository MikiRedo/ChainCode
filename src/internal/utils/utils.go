package utils

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric/common/flogging"
)

var logger = flogging.MustGetLogger("utils")

// TimeTrack traza el tiempo de ejecución
func timeTrack(start time.Time, name string) {
	logger := flogging.MustGetLogger("Timetrack")
	elapsed := time.Since(start)
	logger.Infof(fmt.Sprintf("%s took %s", name, elapsed))
}
