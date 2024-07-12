# gotlsg108e

Implements the Web API used by the WebUI in Go. Started out as a planned Terraform Module

## Supported Actions
- Managing VLAN8021Q
- Retrieving and Setting VLAN PVID
- Rudimentary Login/Logout

## Known Issues
- If you're logged into the Web UI you will be logged out there if you call `client.Logout()` because the switches register the Sessions using the IP Address