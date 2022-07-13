//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InirializedService() (*SimpleService, error) {
	// kasih tahu google wire, function yg akan diguynakan untuk dependecy injectoin
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}
