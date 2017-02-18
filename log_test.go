package log

import (
	"testing"

	"github.com/kasworld/log/loglevels"
)

func TestLog_Output(t *testing.T) {
	t.Logf("%v", logger)
	// logger.Fatal("Fatal %v", loglevels.LL_Fatal)
	logger.Error("Error %v", loglevels.LL_Error)
	logger.Warn("Warn %v", loglevels.LL_Warn)
	logger.Info("Info %v", loglevels.LL_Info)
	logger.Debug("Debug %v", loglevels.LL_Debug)
}
