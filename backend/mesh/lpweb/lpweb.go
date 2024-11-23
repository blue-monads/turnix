package lpweb

import (
	"net/http"

	"github.com/blue-monads/lpweb/code/core/mesh"
	"github.com/blue-monads/lpweb/code/proxy"
	"github.com/blue-monads/lpweb/code/tunnel"
)

type LPWebMesh struct {
	tunnel *tunnel.HttpTunnel
	proxy  *proxy.WebProxy
}

func New(key string, tunnelPort, proxyPort int) *LPWebMesh {

	m, err := mesh.New(key, 0)
	if err != nil {
		panic(err)
	}

	wproxy := proxy.NewUsingMesh(proxyPort, m)
	tunnel := tunnel.NewUsingMesh(tunnelPort, m, false)

	return &LPWebMesh{
		tunnel: tunnel,
		proxy:  wproxy,
	}

}

func (wp *LPWebMesh) HandleHttpIn(r *http.Request, w http.ResponseWriter) {
	wp.proxy.HandleHttp3(r, w)
}

func (wp *LPWebMesh) HandleWsIn(r *http.Request, w http.ResponseWriter) {
	wp.proxy.HandleWS(r, w)
}
