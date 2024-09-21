package climux

import (
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"

	// /random.go backend/utils
	xutils "github.com/bornjre/turnix/backend/utils"
)

var (
	configHome    = path.Join(xdg.ConfigHome, "turnix")
	DerivedSecret = ""
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
		secret, err := xutils.GenerateRandomString(32)
		if err != nil {
			panic(err)
		}

		secret = secret + "-" + string(sha256.New().Sum([]byte(secret)))

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

	DerivedSecret = string(sha.Sum([]byte(wd)))

}

type PeerAddr struct {
	Protocol string `json:"protocol,omitempty" toml:"protocol"`
	Addr     string `json:"addr,omitempty" toml:"addr"`
	Port     int    `json:"port,omitempty" toml:"port"`
	Token    string `json:"token,omitempty" toml:"token"`
}

type ConfigModel struct {
	MeshPort           int                  `json:"mesh_port,omitempty" toml:"mesh_port"`
	MeshKey            string               `json:"mesh_key,omitempty" toml:"mesh_port"`
	AllowTunnelAnyPort bool                 `json:"allow_tunnel_any_port,omitempty" toml:"allow_tunnel_any_port"`
	StaticAddrs        map[string]*PeerAddr `json:"static_addrs,omitempty" toml:"static_addrs"`
	StaticRelays       map[string]*PeerAddr `json:"static_relays,omitempty" toml:"static_relays"`
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

	return c
}

func (c *Configued) init() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if c.BasePath == "" {
		c.BasePath = path.Join(wd, ".turnix-data")
	}

	configPath := c.BasePath + "/config.toml"

	_, err = os.Stat(configPath)
	if err != nil {
		return err
	}

	c.config = &ConfigModel{}

	_, err = toml.DecodeFile(configPath, &c.config)
	if err != nil {
		return err
	}

	if c.config.MeshKey == "" {
		c.config.MeshKey = DerivedSecret
	}

	if c.config.MeshPort == 0 {
		if !isPortUsed(7704) {
			c.config.MeshPort = 7704
		}
	}

	return nil

}

func (c *Configued) GetConfig() *ConfigModel {
	return c.config
}

func isPortUsed(port int) bool {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return true
	}

	defer ln.Close()

	return false
}
