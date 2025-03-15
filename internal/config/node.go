package config

type Node struct {
	ip       string
	port     int
	username string
	password string
}

func (n *Node) GetIp() string {
	return n.ip
}

func (n *Node) GetPort() int {
	return n.port
}

func (n *Node) GetUsername() string {
	return n.username
}

func (n *Node) Authenticate(password string) bool {
	return n.password == password
}

func NewNode(ip string, port int, username string, password string) *Node {
	if ip == "" || port <= 0 || username == "" || password == "" {
		panic("Invalid parameters for Node")
	}
	return &Node{ip: ip, port: port, username: username, password: password}
}
