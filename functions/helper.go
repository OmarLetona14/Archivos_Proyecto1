package functions

import(
	"time"
	"math/rand"
	"strings"
	"runtime"
	"fmt"
)

func GetCurrentTime()(time_array[25] byte){
	dt := time.Now()
	tiempo := dt.Format("01-02-2006 15:04:00")
	copy(time_array[:],tiempo)
	return
}

func generateRandom() int8{
	return int8(rand.Int())
}

func Splitter(txt string) []string {
	commands := strings.Split(txt, " ")
	return commands
}

func CompareBytes(str1 string, str2 string)bool{
	for i:=0;i< len(str1); i++{
		if(!(i>=len(str1)) && !(i>=len(str2))){
			if(!(str1[i]==str2[i])){
				return false
			}
		}
	}
	return true
}

func Calc_filesize(unit string, size int, partition bool)int64{
	if(unit=="" && !partition){
		unit = "m"
	}else if(unit=="" && partition){
		unit = "k"
	}
	switch strings.ToLower(unit) {
	case "k":
		return 1024*int64(size)
	case "m":
		return 1024*1024*int64(size)
	case "b":
		return int64(size)
	default:
		fmt.Println("Invalid unit formmat")
	}
	return 0
}
func Get_text(txt string) string {
	if runtime.GOOS == "windows" {
		txt = strings.TrimRight(txt, "\r\n")
	} else {
		txt = strings.TrimRight(txt, "\n")
	}
	return txt
}