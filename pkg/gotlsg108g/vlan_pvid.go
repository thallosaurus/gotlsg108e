package gotlsg108g

import (
	"encoding/json"
	"log"
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

func (client Client) GetRawPvidConfig() RawPvidData {
	vlan_data := Request(client, "pvid_ds", "Vlan8021QPvidRpm.htm", nil)

	var res RawPvidData
	json_err := json.Unmarshal(vlan_data, &res)

	if nil != json_err {
		log.Fatal(json_err)
	}

	return res
}