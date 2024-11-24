package lpweb

import (
	"net/http"

	"github.com/blue-monads/lpweb/code/core/mesh"
	"github.com/blue-monads/lpweb/code/proxy"
	"github.com/blue-monads/lpweb/code/tunnel"
	"github.com/k0kubun/pp"
)

type LPWebMesh struct {
	tunnel *tunnel.HttpTunnel
	proxy  *proxy.WebProxy
}

func New(key string, meshPort, tunnelPort int) *LPWebMesh {

	pp.Println("@LPWebMesh/0")
	m, err := mesh.NewWithOptions(key, meshPort, true)
	if err != nil {
		panic(err)
	}

	pp.Println("@LPWebMesh/1", m.PublicIp)

	// we are not going to do `wproxy.Run()` so port doesnot matter
	wproxy := proxy.NewUsingMesh(0, m)
	tunnel := tunnel.NewUsingMesh(tunnelPort, m, false)

	return &LPWebMesh{
		tunnel: tunnel,
		proxy:  wproxy,
	}

}

func (wp *LPWebMesh) HandleHttpIn(r *http.Request, w http.ResponseWriter) {

	pp.Println("@HandleHttpIn", r.URL.String())
	// http://mnop.lpweb

	wp.proxy.HandleHttp3(r, w)
}

func (wp *LPWebMesh) HandleWsIn(r *http.Request, w http.ResponseWriter) {
	wp.proxy.HandleWS(r, w)
}
