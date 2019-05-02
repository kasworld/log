// Code generated by "gen_logfn2 -packagename globallogbase -structname LogBase -outputfilename lib/log/globallogbase/loglevelfn_gen.go -levelpackagefolder lib/log/globalloglevels -leveltypename LL_Type"

package globallogbase

import (
	"fmt"

	"github.com/kasworld/log/globalloglevels"
)

func (l *LogBase) Fatal(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_Fatal, format, v...)
	err := l.Output(globalloglevels.LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithFatalLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_Fatal, format, v...)
	err := l.Output(globalloglevels.LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Error(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_Error, format, v...)
	err := l.Output(globalloglevels.LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithErrorLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_Error, format, v...)
	err := l.Output(globalloglevels.LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Warn(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_Warn, format, v...)
	err := l.Output(globalloglevels.LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithWarnLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_Warn, format, v...)
	err := l.Output(globalloglevels.LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceService(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceService, format, v...)
	err := l.Output(globalloglevels.LL_TraceService, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceServiceLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceService, format, v...)
	err := l.Output(globalloglevels.LL_TraceService, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Monitor(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_Monitor, format, v...)
	err := l.Output(globalloglevels.LL_Monitor, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithMonitorLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_Monitor, format, v...)
	err := l.Output(globalloglevels.LL_Monitor, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Debug(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_Debug, format, v...)
	err := l.Output(globalloglevels.LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithDebugLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_Debug, format, v...)
	err := l.Output(globalloglevels.LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) AdminAudit(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_AdminAudit, format, v...)
	err := l.Output(globalloglevels.LL_AdminAudit, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithAdminAuditLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_AdminAudit, format, v...)
	err := l.Output(globalloglevels.LL_AdminAudit, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Analysis(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_Analysis, format, v...)
	err := l.Output(globalloglevels.LL_Analysis, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithAnalysisLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_Analysis, format, v...)
	err := l.Output(globalloglevels.LL_Analysis, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceUser(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceUser, format, v...)
	err := l.Output(globalloglevels.LL_TraceUser, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceUserLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceUser, format, v...)
	err := l.Output(globalloglevels.LL_TraceUser, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceClient(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceClient, format, v...)
	err := l.Output(globalloglevels.LL_TraceClient, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceClientLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceClient, format, v...)
	err := l.Output(globalloglevels.LL_TraceClient, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceAO(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceAO, format, v...)
	err := l.Output(globalloglevels.LL_TraceAO, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceAOLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceAO, format, v...)
	err := l.Output(globalloglevels.LL_TraceAO, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceAI(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceAI, format, v...)
	err := l.Output(globalloglevels.LL_TraceAI, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceAILog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceAI, format, v...)
	err := l.Output(globalloglevels.LL_TraceAI, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceTask(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceTask, format, v...)
	err := l.Output(globalloglevels.LL_TraceTask, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceTaskLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceTask, format, v...)
	err := l.Output(globalloglevels.LL_TraceTask, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceSuspect(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceSuspect, format, v...)
	err := l.Output(globalloglevels.LL_TraceSuspect, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceSuspectLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceSuspect, format, v...)
	err := l.Output(globalloglevels.LL_TraceSuspect, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) TraceRPC(format string, v ...interface{}) {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceRPC, format, v...)
	err := l.Output(globalloglevels.LL_TraceRPC, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithTraceRPCLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, globalloglevels.LL_TraceRPC, format, v...)
	err := l.Output(globalloglevels.LL_TraceRPC, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}
