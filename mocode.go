package mocode

import (
	"github.com/cocotyty/httpclient"
	"fmt"
	"encoding/json"
	"time"
	"strings"
)

type Res struct {
	Code int `json:"code" db:"code"`
	Data struct {
		ID string `json:"id" db:"id"`
	} `json:"data" db:"data"`
}

type CapachaRes struct {
	Code int `json:"code" db:"code"`
	Data struct {
		Recaptcha string `json:"recaptcha" db:"recaptcha"`
	} `json:"data" db:"data"`
}
type MoCodeGoogleClient struct {
	User string
	Pass string
	Act  string
	Refer string
	K string

}

func NewMoCode(user string,pass string,act string,refer string,k string) *MoCodeGoogleClient{
	return &MoCodeGoogleClient{
		User:user,
		Pass:pass,
		Act:act,
		Refer:refer,
		K:k,
	}
}

func (cli *MoCodeGoogleClient) Request() (*Res,error){
	res:=&Res{}
	str,err:=httpclient.Post("https://yzcode.xinby.cn/api.php").Param("user",cli.User).
		        Param("pass",cli.Pass).
			    Param("act","google").
				Param("type","post").
				Param("k",cli.K).
				Param("referer",cli.Refer).Send().String()

				if err!=nil {
					return res,err
				}
				fmt.Println(str)
				err=json.Unmarshal([]byte(str),res)
	if err!=nil {
		return res,err
	}
				return res,nil

}
func (cli *MoCodeGoogleClient) Status(id string) (string,error){
	time.Sleep(time.Second*30)
	fmt.Println("等待验证码识别")
	res:=&CapachaRes{}
	i:=0
	for {
		i++
		fmt.Println("重试第",i,"次")
		time.Sleep(time.Second*10)
		str,err:=httpclient.Post("https://yzcode.xinby.cn/api.php").
			Param("user",cli.User).
			Param("pass",cli.Pass).
			Param("act","google").
			Param("type","res").
			Param("id",id).Send().String()
			if strings.Contains(str,"1811"){
				fmt.Println("稍后重试")
				continue
			}else {
			err=json.Unmarshal([]byte(str),res)
			if err!=nil {
				return "",err
			}
			return res.Data.Recaptcha,nil
			}
	}

}
