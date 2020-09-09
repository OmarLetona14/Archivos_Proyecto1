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

func DeleteQuotes(str string) (strQu string) {
	quLeft := strings.TrimLeft(str, "\"")
	strQu = strings.TrimRight(quLeft, "\"")
	return
}

func ContainsQuotes(str string)bool{
	if(strings.HasPrefix(str, "\"")){
		return true
	}
	return false
}

func CompareBytes(str1 string, str2 string)bool{
	if(str2!=""){
		for i:=0;i< len(str1); i++{
			if(!(i>=len(str1)) && !(i>=len(str2))){
				if(!(str1[i]==str2[i])){
					return false
				}
			}
		}
		return true
	}
	return false
}

func GetString(e []byte)string{
	s := ""
	for _,element := range e{
		if(element!=0){
			s += string(element)
		}
	}
	return s
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

func GetPath(p string)(string, string){
	sp := strings.Split(p, "/")
	name := sp[len(sp)-1]
	path := strings.TrimRight(p, name)
	return path, name
}
