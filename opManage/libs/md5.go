package libs


import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Md5 (buf []byte) string{
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func SetRole(roleinfo string) []string {
	str := strings.Split(roleinfo,",")
	array := make([]string,0)
	for i:=0;i<len(str);i++{
		if i == len(str)-1{
			str1 := strings.Split(str[i],"(")
			str2 := strings.Replace(str1[1],")\"","",-1)
			str3 := strings.Replace(str2,"}","",-1)
			array = append(array,str3)
		}else{
			str1 := strings.Split(str[i],"(")
			str2 := strings.Replace(str1[1],")\"","",-1)
			array = append(array,str2)
		}
	}
	return array
}
