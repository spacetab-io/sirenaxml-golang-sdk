package repository

// Option describes a functional option for configuring the Client.
type Option func(*Repository)

func SetProxyURL(proxyURL string) Option {
	return func(c *Repository) {
		c.Config.ProxyURL = proxyURL
	}
}
