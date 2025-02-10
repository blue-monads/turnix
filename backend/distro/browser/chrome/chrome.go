package chrome

import (
	"context"

	"github.com/chromedp/chromedp"
)

type Chrome struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func New() *Chrome {
	return &Chrome{
		ctx:        nil,
		cancelFunc: nil,
	}
}

func (c *Chrome) Start() error {

	allocatorContext, _ := chromedp.NewRemoteAllocator(context.Background(), "ws://127.0.0.1:9222")

	ctx, cancel := chromedp.NewContext(
		allocatorContext,
	)

	defer cancel()

	c.ctx = ctx
	c.cancelFunc = cancel

	return nil
}

func (c *Chrome) Navigate(url string) error {
	return chromedp.Run(c.ctx,
		chromedp.Navigate(url),
	)
}
