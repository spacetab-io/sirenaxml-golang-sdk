package client

// Option describes a functional option for configuring the Client.
type Option func(*Channel)

func SetClientID(clientID uint16) Option {
	return func(c *Channel) {
		c.cfg.ClientID = clientID
	}
}

func SetClientPrivateKey(clientPrivateKey string) Option {
	return func(c *Channel) {
		c.cfg.ClientPrivateKey = clientPrivateKey
	}
}

func SetClientPrivateKeyPassword(clientPrivateKeyPassword string) Option {
	return func(c *Channel) {
		c.cfg.ClientPrivateKeyPassword = clientPrivateKeyPassword
	}
}

func SetClientPublicKey(clientPublicKey string) Option {
	return func(c *Channel) {
		c.cfg.ClientPublicKey = clientPublicKey
	}
}

func SetIp(ip string) Option {
	return func(c *Channel) {
		c.cfg.Ip = ip
	}
}

func SetEnvironment(environment string) Option {
	return func(c *Channel) {
		c.cfg.Environment = environment
	}
}

func SetZippedMessaging(zippedMessaging bool) Option {
	return func(c *Channel) {
		c.cfg.ZippedMessaging = zippedMessaging
	}
}

func SetServerPublicKey(serverPublicKey string) Option {
	return func(c *Channel) {
		c.cfg.ServerPublicKey = serverPublicKey
	}
}

func SetMaxConnections(maxConnections uint32) Option {
	return func(c *Channel) {
		c.cfg.MaxConnections = maxConnections
	}
}

func SetSendChannel(buffer int) Option {
	return func(c *Channel) {
		c.send = make(chan *Packet, buffer)
	}
}

func SetSocket(addr string) Option {
	return func(c *Channel) {
		c.socket = &socket{addr: addr}
	}
}
