package gotlsg108g

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/thallosaurus/gotlsg108g/pkg/gotlsg108g/pages"
)

func CheckLogin(client Client) bool {
	login_status_res := Request(client, "logonInfo", pages.QVlanSet, nil)

	var res []int64
	json_err := json.Unmarshal(login_status_res, &res)

	if nil != json_err {
		log.Fatal(json_err)
	}

	log.Println("check login ", res)

	// 1 means no session
	return res[0] == 1
}

func Login(host string, username string, password string) bool {
	v := url.Values{}
	v.Add("username", username)
	v.Add("password", password)
	v.Add("cpassword", "")
	v.Add("logon", "Login")

	linfo := DataRequest(Client{host: host, login: false}, "logonInfo", pages.LOGON, v)

	var buf []int64
	err := json.Unmarshal(linfo, &buf)

	if nil != err {
		log.Fatal(err)
	}
	return buf[0] == 0
	//}
}

func Logout(host string) {
	Request(Client{host: host, login: false}, "", string(pages.LOGOUT), nil)

	//log.Println(res)
}
