// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package welcome

// Injectors from wire.go:

// Wire initializes and returns a new Handler instance.
//
// This function uses the wire package to build the Handler instance by calling the ProviderSet function.
// The wire package is used for dependency injection and ensures that all dependencies are properly initialized.
// The function panics if there is an error during the build process.
//
// Returns:
// - *Handler: The initialized Handler instance.
func Wire() *Handler {
	handler := ProvideHandler()
	return handler
}
