package main

import "fmt"

func main() {
	v := "=dfd="
	fmt.Println(v[:4])
}

/*
func ParseIni(initData []string)map[string]map[string]string {
    mapIniData:=make(map[string]map[string]string,10)
	t:=false
	for index,value:=range initData {
		if value[0]==59 || value[0]==44 ||len(value)==0 {
			continue
		}
		if value[0]==91 && value[len(value)-1]==93 {
            t=true
		}
		if t {

		}
		t=false
	}

}
func ParsIniOne(initData []string, index int) map[string]string {
	t := func(value string) int {
		for i := 0; i < len(value); i++ {
			if value[i] == 61 {
				return len(value) - i
				break
			}
		}
	}
	mapData := make(map[string]string, 10)
	for i := index + 1; i < len(initData); i++ {
		if initData[i][0] == 91 && initData[i][len(initData[i])-1] == 93 {
			break
		}
		mapData[initData[i][:t(value1)]] = initData[i][t(value1):]
	}

}
*/
