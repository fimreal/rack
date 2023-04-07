package ngrok

import (
	"context"

	"github.com/spf13/viper"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func New() (ngrok.Tunnel, error) {
	token := viper.GetString("ngrok.token")
	domain := viper.GetString("ngrok.domain")

	if domain != "" {
		return ngrok.Listen(
			context.Background(),
			config.HTTPEndpoint(
				config.WithDomain(domain),
			),
			ngrok.WithAuthtoken(token),
		)
	}

	// Default tun
	return ngrok.Listen(
		context.Background(),
		config.HTTPEndpoint(),
		ngrok.WithAuthtoken(token),
	)
}
