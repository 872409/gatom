package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	l := New("aa")
	l.SaveToFile("aa", "vv.log")
	l.SetLevel(DebugLevel)
	l.Logger.Infoln("Test..")
	SetDefault(l)

	Debugln("Dbug")
	l.Infoln("AAA")
	l.Warnln("AAA")
}
