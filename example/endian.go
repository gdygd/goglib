package main

import (
	"goglib"
)

// ---------------------------------------------------------------------------
// Log
// ---------------------------------------------------------------------------
var Applog *goglib.OLog2 = goglib.InitLogEnv("./logdir", "process1", 2) // "path, stamp string, initial debug lv"
func main() {
	var buf1 []byte = make([]byte, 2)
	var buf2 []byte = make([]byte, 4)
	var buf3 []byte = make([]byte, 8)

	goglib.SetNumber(buf1, 0, 1, 2, goglib.ED_BIG)
	goglib.SetNumber(buf2, 0, 2, 4, goglib.ED_BIG)
	goglib.SetNumber(buf3, 0, 1152921504606846977, 8, goglib.ED_LITTLE)

	Applog.Print(2, "buf1 (big endian)")
	Applog.Dump(2, "buf1 rawdata", buf1, len(buf1))

	Applog.Print(2, "buf2 (big endian)")
	Applog.Dump(2, "buf2 rawdata", buf2, len(buf2))

	Applog.Print(2, "buf3 (little endian)")
	Applog.Dump(2, "buf3 rawdata", buf3, len(buf3))

	val1 := goglib.GetNumber(buf1, 0, 2, goglib.ED_BIG)
	val2 := goglib.GetNumber(buf2, 0, 4, goglib.ED_BIG)
	val3 := goglib.GetNumber(buf3, 0, 8, goglib.ED_LITTLE)

	Applog.Print(2, "buf1 val : %d", val1)
	Applog.Print(2, "buf2 val : %d", val2)
	Applog.Print(2, "buf3 val : %d", val3)

	Applog.Fileclose()
}
