package browser

type Browser interface {
	Start() error
	Navigate(url string) error
}

var InstanceBrowser Browser = nil
