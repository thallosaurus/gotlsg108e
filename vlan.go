package gotlsg108e

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/thallosaurus/gotlsg108e/pages"
)

type VlanData map[string]SingleVlanConfig

type SingleVlanConfig struct {
	Name            string
	Id              int64
	TaggedMembers   []bool
	UntaggedMembers []bool
}

type RawVlanData struct {
	State   int64    `json:"state"`
	PortNum int64    `json:"portNum"`
	VlanIds []int64  `json:"vids"`
	Count   int64    `json:"count"`
	MaxVids int64    `json:"maxVids"`
	Names   []string `json:"names`

	// Tagged VLAN Members
	TagMembers   []int64 `json:"tagMbrs"`
	UntagMembers []int64 `json:"untagMbrs"`
	LagIds       []int64 `json:"lagIds"`
	LagMembers   []int64 `json:"lagMbrs"`
}

type QVlan struct {
	VlanId   int64
	VlanName string
	Selected []SetVlanConfType
}

type SetVlanConfType string

const (
	SelTypeUntagged  SetVlanConfType = "0"
	SelTypeTagged    SetVlanConfType = "1"
	SelTypeNotMember SetVlanConfType = "2"
)

func (client Client) SetVlanConfig(qres QVlan) {
	v := url.Values{}
	v.Add("vid", fmt.Sprintf("%d", qres.VlanId))
	v.Add("vname", qres.VlanName)

	for key, va := range qres.Selected {
		k := fmt.Sprintf("selType_%d", key+1)
		v.Add(k, string(va))
	}

	log.Println(v)

	v.Add("qvlan_add", "Add/Modify")

	Request(client, "qvlan_ds", pages.QVlanSet, &v)
}

func (client Client) DeleteVlanConfig(vlan QVlan) {
	v := url.Values{}
	v.Add("selVlans", fmt.Sprintf("%d", vlan.VlanId))
	v.Add("qvlan_del", "Delete")

	Request(client, "qvlan_ds", pages.QVlanSet, &v)
}

func (client Client) GetRawVlanConfig() RawVlanData {

	vlan_data := Request(client, "qvlan_ds", "Vlan8021QRpm.htm", nil)

	var res RawVlanData
	json_err := json.Unmarshal(vlan_data, &res)

	if nil != json_err {
		log.Fatal(json_err)
	}

	return res
}

func (client Client) GetVlanConfig() VlanData {
	data := client.GetRawVlanConfig()
	log.Println(data)
	mapped := parseVlanConfig(data)

	return mapped
}

func binmaskToArray(mask int64) []bool {
	b := []bool{
		mask&(1<<0) > 0,
		mask&(1<<1) > 0,
		mask&(1<<2) > 0,
		mask&(1<<3) > 0,
		mask&(1<<4) > 0,
		mask&(1<<5) > 0,
		mask&(1<<6) > 0,
		mask&(1<<7) > 0,
	}

	return b

}

func arrayToBinmask(a []bool) int {
	mask := 0
	for index, v := range a {
		if v {
			mask += (1 << index)
		}
	}
	return mask
}

func parseVlanConfig(data RawVlanData) VlanData {
	m := make(VlanData)

	for key, val := range data.Names {
		m[val] = SingleVlanConfig{
			Name:            val,
			Id:              data.VlanIds[key],
			TaggedMembers:   binmaskToArray(data.TagMembers[key]),
			UntaggedMembers: binmaskToArray(data.UntagMembers[key]),
		}
	}

	return m
}
