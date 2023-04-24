package main

type binary struct {
	Protoc string `default:"${BINARY_PROTOC=protoc}"`
	Lint   string `default:"${BINARY_LINT=protolint}"`
	Gtag   string `default:"${BINARY_GTAG=gtag}"`
}
