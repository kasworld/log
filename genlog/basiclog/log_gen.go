// Code generated by "genlog -leveldatafile basiclevel.data -packagename=basiclog"

package basiclog

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/kasworld/log/logdestination_stdio"
	"github.com/kasworld/log/logdestinationgroup"
	"github.com/kasworld/log/logdestinationi"
	"github.com/kasworld/log/logflagi"
)

type LL_Type uint64

func (ll LL_Type) LevelsString() string {
	var buff bytes.Buffer
	buff.WriteString("LL_Type[")
	for ll := LL_Type(1); ll < LL_END; ll <<= 1 {
		if ll.IsLevel(ll) {
			fmt.Fprintf(&buff, "%s, ", ll)
		}
	}
	buff.WriteString("]")
	return buff.String()
}
func (ll LL_Type) String() string {
	if str, ok := leveldata[ll]; ok {
		return str
	}
	return "LL_Type(" + strconv.FormatInt(int64(ll), 10) + ")"
}
func (ll LL_Type) IsLevel(l2 LL_Type) bool {
	return ll&l2 != 0
}
func (ll LL_Type) AllLevel() LL_Type {
	return LL_All
}
func (ll LL_Type) StartLevel() LL_Type {
	return LL_Fatal
}
func (ll LL_Type) LevelCount() int {
	return LL_Count
}
func (ll LL_Type) IsLastLevel() bool {
	return ll == LL_END
}
func (ll LL_Type) NextLevel(n uint) LL_Type {
	return ll << n
}
func (ll LL_Type) PreLevel(n uint) LL_Type {
	return ll >> n
}
func (ll LL_Type) BitAnd(l2 LL_Type) LL_Type {
	return ll & l2
}
func (ll LL_Type) BitOr(l2 LL_Type) LL_Type {
	return ll | l2
}
func (ll LL_Type) BitXor(l2 LL_Type) LL_Type {
	return ll ^ l2
}
func (ll LL_Type) BitClear(l2 LL_Type) LL_Type {
	return ll &^ l2
}
func (ll LL_Type) BitTest(l2 LL_Type) bool {
	return ll&l2 != 0
}
func (ll LL_Type) TestAt(n int) bool {
	return ll&LL_Type(1<<n) != 0
}

const (
	LL_Fatal LL_Type = 1 << iota //
	LL_Error                     //
	LL_Warn                      //
	LL_Debug                     //

	LL_END
	LL_All   = LL_END - 1
	LL_Count = 4
)

var leveldata = map[LL_Type]string{
	1: "Fatal",
	2: "Error",
	4: "Warn",
	8: "Debug",

	16: "END",
}

//////////////////////////////////////////////////////////////////

var (
	OutputStdout = logdestination_stdio.NewStdOut()
	OutputStderr = logdestination_stdio.NewStdErr()
)

//////////////////////////////////////////////////////////////////

type logDestInfo struct {
	refCntByLogLv int // count referenced by each loglv
	dest          logdestinationi.LogDestinationI
}

type LogBase struct {
	mutex sync.RWMutex

	flag     logflagi.LogFlagI // properties
	prefix   string            // prefix to write at beginning of each line
	loglevel LL_Type

	ltype2destgrp     []*logdestinationgroup.LogDestinationGroup
	allDestInfoByName map[string]*logDestInfo
}

func New(prefix string, lf logflagi.LogFlagI, lv LL_Type) *LogBase {

	dstgrp := make([]*logdestinationgroup.LogDestinationGroup, lv.LevelCount())
	for i := 0; i < lv.LevelCount(); i++ {
		dstgrp[i] = logdestinationgroup.New()
	}

	return &LogBase{
		ltype2destgrp:     dstgrp,
		allDestInfoByName: make(map[string]*logDestInfo),
		flag:              lf,
		prefix:            prefix,
		loglevel:          lv,
	}
}

func makeLogFilename(logdir string, ll string) string {
	basename := filepath.Base(logdir)
	filename := fmt.Sprintf("%s.%s.%s", basename, ll, "log")
	return filepath.Join(logdir, filename)
}
func (lg *LogBase) AddDestination(
	ll LL_Type, o logdestinationi.LogDestinationI) {

	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	for i := 0; i < ll.LevelCount(); i++ {
		if ll.TestAt(i) {
			lg.addDestination1DestGrp(i, o)
		}
	}
}

func (lg *LogBase) addDestination1DestGrp(
	i int, o logdestinationi.LogDestinationI) {

	added := lg.ltype2destgrp[i].AddDestination(o)
	if !added {
		if _, ok := lg.allDestInfoByName[o.Name()]; !ok {
			panic(fmt.Sprintf(
				"%v failed to AddDestination to destgroup index:%v, abnormal state",
				lg, i))
		}
		fmt.Printf("%v not added to destgroup index:%v\n", o, i)
		return
	}

	if dstinfo, ok := lg.allDestInfoByName[o.Name()]; ok {
		dstinfo.refCntByLogLv++
	} else {
		lg.allDestInfoByName[o.Name()] = &logDestInfo{
			refCntByLogLv: 1,
			dest:          o,
		}
	}
}

func (lg *LogBase) DelDestination(
	ll LL_Type, o logdestinationi.LogDestinationI) {

	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	for i := 0; i < ll.LevelCount(); i++ {
		if ll.TestAt(i) {
			lg.delDestinationFrom1DestGrp(i, o)
		}
	}
}

func (lg *LogBase) delDestinationFrom1DestGrp(
	i int, o logdestinationi.LogDestinationI) {

	deleted := lg.ltype2destgrp[i].DelDestination(o)
	if !deleted {
		fmt.Printf("%v not deleted from destgroup index:%v\n", o, i)
		return
	}

	if dstinfo, ok := lg.allDestInfoByName[o.Name()]; ok {
		dstinfo.refCntByLogLv--
		if dstinfo.refCntByLogLv <= 0 {
			delete(lg.allDestInfoByName, o.Name())
		}
	} else {
		panic(fmt.Sprintf(
			"%v failed to DelDestination %v from destgroup index:%v, abnormal state",
			lg, o, i,
		))
	}
}

func (lg *LogBase) Reload() error {
	lg.mutex.RLock()
	defer lg.mutex.RUnlock()

	for _, v := range lg.allDestInfoByName {
		if err := v.dest.Reload(); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func (lg *LogBase) LogPrintf(calldepth int, ll LL_Type,
	format string, v ...interface{}) ([]byte, error) {
	s := lg.Format2Bytes(calldepth+1, ll, format, v...)
	err := lg.Output(ll, s)
	return s, err
}

func (lg *LogBase) Format2Bytes(calldepth int, ll LL_Type,
	format string, v ...interface{}) []byte {

	if !lg.loglevel.IsLevel(ll) {
		return nil
	}
	s := fmt.Sprintf(format, v...)
	var buf []byte
	llinfo := fmt.Sprintf("%s", ll)
	lg.flag.FormatHeader(&buf, calldepth+2, time.Now(), lg.prefix, llinfo)
	buf = append(buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	return buf
}

func (lg *LogBase) Output(ll LL_Type, b []byte) error {
	var err error
	for i := 0; i < ll.LevelCount(); i++ {
		if ll.TestAt(i) {
			if lerr := lg.ltype2destgrp[i].Write(b); lerr != nil {
				err = lerr
			}
		}
	}
	return err
}

func (lg *LogBase) SetLevel(level LL_Type) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.loglevel = level
}

func (lg *LogBase) GetLevel() LL_Type {
	return lg.loglevel
}

func (lg *LogBase) SetPrefix(p string) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.prefix = p
}

// Prefix returns the output prefix for the logger.
func (lg *LogBase) GetPrefix() string {
	return lg.prefix
}

// Flags returns the output flags for the logger.
func (lg *LogBase) GetFlags() logflagi.LogFlagI {
	return lg.flag
}

// SetFlags sets the output flags for the logger.
func (lg *LogBase) SetFlags(flag logflagi.LogFlagI) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.flag = flag
}

func (lg *LogBase) String() string {
	return fmt.Sprintf("LogBase[%v %v]",
		lg.flag.FlagString(), lg.loglevel.LevelsString(),
	)
}

func (l *LogBase) Fatal(format string, v ...interface{}) {
	if !l.GetLevel().IsLevel(LL_Fatal) {
		return
	}
	s := l.Format2Bytes(1, LL_Fatal, format, v...)
	err := l.Output(LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogBase) Error(format string, v ...interface{}) {
	if !l.GetLevel().IsLevel(LL_Error) {
		return
	}
	s := l.Format2Bytes(1, LL_Error, format, v...)
	err := l.Output(LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogBase) Warn(format string, v ...interface{}) {
	if !l.GetLevel().IsLevel(LL_Warn) {
		return
	}
	s := l.Format2Bytes(1, LL_Warn, format, v...)
	err := l.Output(LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogBase) Debug(format string, v ...interface{}) {
	if !l.GetLevel().IsLevel(LL_Debug) {
		return
	}
	s := l.Format2Bytes(1, LL_Debug, format, v...)
	err := l.Output(LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
}

//////////////////////////////////////////////////////////////////
var GlobalLogger *LogBase

func Fatal(format string, v ...interface{}) {
	if !GlobalLogger.GetLevel().IsLevel(LL_Fatal) {
		return
	}
	s := GlobalLogger.Format2Bytes(1, LL_Fatal, format, v...)
	err := GlobalLogger.Output(LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
}

func Error(format string, v ...interface{}) {
	if !GlobalLogger.GetLevel().IsLevel(LL_Error) {
		return
	}
	s := GlobalLogger.Format2Bytes(1, LL_Error, format, v...)
	err := GlobalLogger.Output(LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
}

func Warn(format string, v ...interface{}) {
	if !GlobalLogger.GetLevel().IsLevel(LL_Warn) {
		return
	}
	s := GlobalLogger.Format2Bytes(1, LL_Warn, format, v...)
	err := GlobalLogger.Output(LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
}

func Debug(format string, v ...interface{}) {
	if !GlobalLogger.GetLevel().IsLevel(LL_Debug) {
		return
	}
	s := GlobalLogger.Format2Bytes(1, LL_Debug, format, v...)
	err := GlobalLogger.Output(LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
}
