package assetserve

import (
	"sync"

	"github.com/bornjre/turnix/backend/distro"
)

func Run(app *distro.DistroApp, wg *sync.WaitGroup) {
	defer wg.Done()

}
