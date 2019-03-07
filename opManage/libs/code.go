package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetAccessToken() string{
	appid := beego.AppConfig.String("appid")
	secretId := beego.AppConfig.String("sercetToken")
	baseUrl := "https://oapi.dingtalk.com/sns/gettoken?appid="
	AccessTokenUrl := baseUrl+appid+"&appsecret="+secretId
	responseToken,err:= http.Get(AccessTokenUrl)
	if err != nil{
		fmt.Println(err)
	}
	defer responseToken.Body.Close()//当获取到响应体时，客户端必须手动关闭链接
	jsonStr,err := ioutil.ReadAll(responseToken.Body)
	if err != nil{
		fmt.Println(err)
	}
	jsonTostr := strings.Split(string(jsonStr),",")[1]
	access_token := strings.Split(jsonTostr,":")[1]
	access_token = strings.Replace(access_token,"\"","",-1)
	//fmt.Println(access_token)
	return access_token
}

func GetPerpetualCode(code string,token string) (string,string) {
	data := make(map[string]string)
	data["tmp_auth_code"] = code
	jsonData,err := json.Marshal(data)

	baseurl := "https://oapi.dingtalk.com/sns/get_persistent_code?"
	urlencode := url.Values{}
	urlencode.Set("access_token",token)
	ResquestUrl := baseurl + urlencode.Encode()
	//fmt.Println(ResquestUrl)
	responseToken,err := http.NewRequest("POST",ResquestUrl,bytes.NewBuffer(jsonData))
	responseToken.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp, err := client.Do(responseToken)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	jsonStr,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	jsonTostrCode := strings.Split(string(jsonStr),",")[0]
	jsonTostrOpenid := strings.Split(string(jsonStr),",")[1]
	persistent_code := strings.Split(jsonTostrCode,":")[1]
	openid := strings.Split(jsonTostrOpenid,":")[1]
	persistent_code = strings.Replace(persistent_code,"\"","",-1)
	openid = strings.Replace(openid,"\"","",-1)
	return persistent_code,openid
}

func GetSnsToken(code string,openid string,token string) string {
	data := make(map[string]string)
	data["openid"] = openid
	data["persistent_code"] = code
	jsonData,err := json.Marshal(data)
	baseurl := "https://oapi.dingtalk.com/sns/get_sns_token?"
	urlencode := url.Values{}
	urlencode.Set("access_token",token)
	ResquestUrl := baseurl + urlencode.Encode()
	responseToken,err := http.NewRequest("POST",ResquestUrl,bytes.NewBuffer(jsonData))
	responseToken.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp, err := client.Do(responseToken)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	jsonStr,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	errmsgdict := strings.Split(string(jsonStr),",")[0]
	errmsg := strings.Split(string(errmsgdict),":")[1]
	errmsg = strings.Replace(errmsg,"\"","",-1)
	if errmsg == "ok"{
		jsonTostr := strings.Split(string(jsonStr),",")[2]
		sns_token := strings.Split(jsonTostr,":")[1]
		sns_token = strings.Replace(sns_token,"\"","",-1)
		return sns_token
	}else{
		sns_token := "errmsg"
		return sns_token
	}

}

func GetUserInfo(token string) (string,string) {
	baseurl := "https://oapi.dingtalk.com/sns/getuserinfo?sns_token="
	accessUrl := baseurl + token
	responseToken,err:= http.Get(accessUrl)
	if err != nil{
		fmt.Println(err)
	}
	defer responseToken.Body.Close()//当获取到响应体时，客户端必须手动关闭链接
	jsonStr,err := ioutil.ReadAll(responseToken.Body)
	if err != nil{
		fmt.Println(err)
	}
	errmsgdict := strings.Split(string(jsonStr),",")[0]
	errmsg := strings.Split(string(errmsgdict),":")[0]
	errmsg = strings.Replace(errmsg,"\"","",-1)
	if errmsg == "{errmsg"{
		name := "error"
		openid := "error"
		return name,openid
	}else{
		jsonTostrName := strings.Split(string(jsonStr),",")[1]
		name := strings.Split(jsonTostrName,":")[1]
		name = strings.Replace(name,"\"","",-1)
		jsonTostrOpenid := strings.Split(string(jsonStr),",")[3]
		openid := strings.Split(jsonTostrOpenid,":")[1]
		openid = strings.Replace(openid,"\"","",-1)
		openid = strings.Replace(openid,"}","",-1)
		return name,openid
	}
}