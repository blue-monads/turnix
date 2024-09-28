package climux

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/k0kubun/pp"

	xutils "github.com/bornjre/turnix/backend/utils"
)

var (
	configHome          = path.Join(xdg.ConfigHome, "turnix")
	derivedMasterSecret = ""
	derivedMeshKey      = ""
)

const (
	SecretFileName = "secret.txt"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// insure config dir, resursesively create
	err = os.MkdirAll(configHome, os.ModePerm)
	if err != nil {
		panic(err)
	}

	secretFile := path.Join(configHome, SecretFileName)

	// check file exists
	_, err = os.Stat(secretFile)
	if err != nil {
		// genrate random secret base
		secret, err := xutils.GenerateRandomString(128)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(secretFile, []byte(secret), 0600)
		if err != nil {
			panic(err)
		}
	}

	sout, err := os.ReadFile(secretFile)
	if err != nil {
		panic(err)
	}

	sha := sha256.New()
	sha.Write(sout)
	sha.Write([]byte(wd))

	derivedMasterSecret = base64.StdEncoding.EncodeToString(sha.Sum(nil))
	derivedMeshKey = base64.StdEncoding.EncodeToString(sha.Sum([]byte("node")))

}

type PeerAddr struct {
	Protocol string `json:"protocol,omitempty" toml:"protocol"`
	Addr     string `json:"addr,omitempty" toml:"addr"`
	Port     int    `json:"port,omitempty" toml:"port"`
	Token    string `json:"token,omitempty" toml:"token"`
}

type ConfigModel struct {
	HttpServerPort     int                  `json:"http_server_port,omitempty" toml:"http_server_port"`
	MeshPort           int                  `json:"mesh_port,omitempty" toml:"mesh_port"`
	MeshKey            string               `json:"mesh_key,omitempty" toml:"mesh_port"`
	AllowTunnelAnyPort bool                 `json:"allow_tunnel_any_port,omitempty" toml:"allow_tunnel_any_port"`
	StaticAddrs        map[string]*PeerAddr `json:"static_addrs,omitempty" toml:"static_addrs"`
	StaticRelays       map[string]*PeerAddr `json:"static_relays,omitempty" toml:"static_relays"`
	NodeCtrlKey        string               `json:"node_ctrl_key,omitempty" toml:"node_ctrl_key"`
	LocalSocket        string               `json:"local_socket,omitempty" toml:"local_socket"`
	MasterKey          string               `json:"master_key,omitempty" toml:"master_key"`
	DatabaseFile       string               `json:"database_file,omitempty" toml:"database_file"`
}

type Configued struct {
	BasePath string
	config   *ConfigModel
}

func NewConfigued(basepath string) *Configued {
	c := &Configued{
		BasePath: basepath,
		config:   nil,
	}

	err := c.init()
	if err != nil {
		pp.Println("config init error", err)
		panic(err)
	}

	return c
}

func (c *Configued) init() error {
	pp.Println("init/1")

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	pp.Println("init/2")

	if c.BasePath == "" {
		c.BasePath = path.Join(wd, ".turnix-data")
	}

	pp.Println("init/3")

	configPath := path.Join(c.BasePath, "/config.toml")

	_, err = os.Stat(configPath)
	if err == nil {
		c.config = &ConfigModel{}

		_, err = toml.DecodeFile(configPath, &c.config)
		if err != nil {
			return err
		}
	} else {
		c.config = &ConfigModel{}
	}

	pp.Println("init/4")

	if c.config.MeshKey == "" {
		c.config.MeshKey = derivedMeshKey
	}

	if c.config.MasterKey == "" {
		c.config.MasterKey = derivedMasterSecret
	}

	if c.config.LocalSocket == "" {
		c.config.LocalSocket = path.Join(c.BasePath, "/local.sock")
	}

	if c.config.DatabaseFile == "" {
		c.config.DatabaseFile = path.Join(c.BasePath, "/data.db")
	}

	pp.Println("init/5")

	if c.config.MeshPort == 0 {
		if !isPortUsed(7704) {
			c.config.MeshPort = 7704
		}
	}

	pp.Println("init/6")

	if c.config.HttpServerPort == 0 {
		if !isPortUsed(7703) {
			c.config.HttpServerPort = 7703
		}
	}

	pp.Println("init/7")

	if c.config.NodeCtrlKey == "" {

		nodeCtrl, err := xutils.GenerateRandomString(32)
		if err != nil {
			return err
		}

		c.config.NodeCtrlKey = nodeCtrl
	}

	return nil

}

func (c *Configued) GetConfig() *ConfigModel {
	return c.config
}

func (c *Configued) InitPath() error {
	return os.MkdirAll(c.BasePath, os.ModePerm)
}

func isPortUsed(port int) bool {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return true
	}

	defer ln.Close()

	return false
}
