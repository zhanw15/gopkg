package time

/*
   const char* build_time(void)
   {
   	return __TIME__;    //"10:59:19"
   }
   const char* build_date(void)
   {
   	return __DATE__;    //"Sep 18 2010"
   }
*/
import "C"

import (
"fmt"
"strings"
)

var monthMap = map[string]string{
	"Jan":"01",
	"Feb":"02",
	"Mar":"03",
	"Apr":"04",
	"May":"05",
	"Jun":"06",
	"Jul":"07",
	"Aug":"08",
	"Sep":"09",
	"Oct":"10",
	"Nov":"11",
	"Dec":"12",
}

func GetBuildTime() string {
	buildTime := string(C.GoString(C.build_time()))
	buildDate := string(C.GoString(C.build_date()))

	dateArr := filterSpace(strings.Split(buildDate, " "))
	if len(dateArr)!=3 {
		return ""
	}

	month, day, year := dateArr[0], dateArr[1], dateArr[2]
	month = monthMap[month]

	return fmt.Sprintf("%s-%s-%s %s", month, day, year, buildTime)
}

func filterSpace(sls []string) (res []string){
	for _, v := range sls {
		if v!="" && v!=" " {
			res = append(res, v)
		}
	}
	return
}