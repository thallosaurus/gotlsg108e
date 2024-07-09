package main

import (
	"log"

	driver "github.com/thallosaurus/gotlsg108g"
)

func main() {
	//tlsg108g.SetHost("10.0.1.4")
	/*if tlsg108g.Login("admin", "admin1") {
		vlan_data := tlsg108g.GetVlanConfig()
		log.Printf("%+v\n", vlan_data)
		vlan_pvid_data := tlsg108g.GetRawPvidConfig()
		log.Printf("%+v\n", vlan_pvid_data)

		vlan := tlsg108g.QVlan{
			VlanId:   50,
			VlanName: "test1234",
			Selected: []tlsg108g.SetVlanConfType{
				tlsg108g.SelTypeTagged,
				tlsg108g.SelTypeTagged,
				tlsg108g.SelTypeTagged,
				tlsg108g.SelTypeTagged,
				tlsg108g.SelTypeTagged,
				tlsg108g.SelTypeTagged,
				tlsg108g.SelTypeTagged,
				tlsg108g.SelTypeTagged,
			},
		}

		tlsg108g.SetVlanConfig(vlan)
		tlsg108g.DeleteVlanConfig(vlan)
		tlsg108g.Logout()
	} */

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
	client.DeleteVlanConfig(vlan)
}
