package main

import (
	"log"

	driver "github.com/thallosaurus/gotlsg108e"
)

func main() {
	client := driver.New("10.0.1.4", "admin", "admin1")
	defer client.Close()

	vlan_data := client.GetVlanConfig()
	log.Printf("%+v\n", vlan_data)
	vlan_pvid_data := client.GetRawPvidConfig()
	log.Printf("%+v\n", vlan_pvid_data)

	vlan := driver.QVlan{
		VlanId:   50,
		VlanName: "test1234",
		Selected: []driver.SetVlanConfType{
			driver.SelTypeTagged,
			driver.SelTypeTagged,
			driver.SelTypeTagged,
			driver.SelTypeTagged,
			driver.SelTypeTagged,
			driver.SelTypeTagged,
			driver.SelTypeTagged,
			driver.SelTypeTagged,
		},
	}

	client.SetVlanConfig(vlan)
	client.SetPvidConfig(50, []bool{false, false, false, false, false, false, false, true})
	//client.AddPvidConfig(1, []bool{false, false, false, false, false, false, false, true})
	client.DeleteVlanConfig(vlan)
}
