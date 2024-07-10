package gotlsg108e

type Client struct {
	host  string
	login bool
}

func (self *Client) Close() {
	Logout(self.host)
}

func New(host string, username string, password string) *Client {
	if Login(host, username, password) {
		c := Client{
			login: true,
			host:  host,
		}

		return &c
	} else {
		return nil
	}
}
