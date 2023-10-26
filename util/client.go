package util

import (
	tlsClient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/rs/zerolog/log"
)

func NewClient() (*tlsClient.HttpClient, error) {
	jar := tlsClient.NewCookieJar()

	options := []tlsClient.HttpClientOption{
		tlsClient.WithTimeoutSeconds(30),
		tlsClient.WithClientProfile(profiles.Chrome_105),
		// tlsClient.WithRandomTLSExtensionOrder(),
		tlsClient.WithCookieJar(jar),
	}

	client, err := tlsClient.NewHttpClient(tlsClient.NewNoopLogger(), options...)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create a client")
		return nil, err
	}

	return &client, nil
}
