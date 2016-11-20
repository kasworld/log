package log

import "testing"

func TestLog_Output(t *testing.T) {
	t.Logf("%v", logger)
	// logger.Fatal("Fatal %v", LL_Fatal)
	logger.Error("Error %v", LL_Error)
	logger.Warn("Warn %v", LL_Warn)
	logger.Info("Info %v", LL_Info)
	logger.Debug("Debug %v", LL_Debug)
}
