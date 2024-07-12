package gotlsg108e

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/thallosaurus/gotlsg108e/pages"
)

type RawPvidData struct {
	State      int64   `json:"state"`
	PortNumber int64   `json:"portNum"`
	VlanIds    []int64 `json:"vids"`
	Count      int64   `json:"count"`
	Members    []int64 `json:"mbrs"`
	LagIds     []int64 `json:"lagIds"`
	LagMembers []int64 `json:"lagMbrs"`
}

type SetPvidData struct {
}

func (client Client) GetRawPvidConfig() RawPvidData {
	vlan_data := Request(client, "pvid_ds", "Vlan8021QPvidRpm.htm", nil)

	var res RawPvidData
	json_err := json.Unmarshal(vlan_data, &res)

	if nil != json_err {
		log.Fatal(json_err)
	}

	return res
}

func (client Client) SetPvidConfig(pvid int64, ports []bool) {
	payload := url.Values{}
	payload.Add("pbm", fmt.Sprintf("%d", arrayToBinmask(ports)))
	payload.Add("pvid", fmt.Sprintf("%d", pvid))

	fmt.Print(payload)

	Request(client, "pvid_ds", pages.PVIDSet, &payload)
}
