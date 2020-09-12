package functions

import(
	"time"
	"math/rand"
	"strings"
	"runtime"
	"fmt"
	"os/exec"
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
	var commands []string
	if(!strings.ContainsAny(txt, "\"")){
		commands = strings.Split(txt, " ")
	}else{
		precommands := strings.Split(txt, " ")
		commands = QuotesPath(precommands)
	}
	return commands
}

func QuotesPath(str []string)[]string{
	pt := ""
	for i,_ := range str{
		if(strings.HasPrefix(str[i],"-path")){
			nxt := 0
			for !strings.HasSuffix(str[i+nxt], "\""){
				if nxt!=0{
					pt +=" "+ str[i+nxt]
				}else{
					pt += str[i+nxt]
				}
				str[i+nxt] = ""
				nxt +=1
			}
			pt += " " + str[i+nxt]
			str[i+nxt] = ""
			str[i] = pt
		}
	}
	return str
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

func GetWSpace(r string)string{
	var str string = ""
	if(strings.HasPrefix(r," ")){
		str = strings.TrimRight(r, " ")
	}else if(strings.HasSuffix(r, " ")){
		str = strings.TrimLeft(r, " ")
	}else{
		str = r
	}
	return str
}

func execDot(input string, output string){
	cmd := exec.Command("dot","-Tpng", input, "-o", output)
	err:= cmd.Run()
	if(err!=nil){
		fmt.Println("**** ERROR: CANNOT CREATE REPORT ******")
	}else{
		fmt.Println("Report created correctly")
	}
}
