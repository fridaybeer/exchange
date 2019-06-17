// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"exchange/models"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"exchange/restapi/operations"
	"exchange/restapi/operations/asset"
	"exchange/restapi/operations/order"
	"exchange/restapi/operations/trader"
)

//go:generate swagger generate server --target ../../exchange --name ExchangeApp --spec ../swagger.yaml

var traders map[string]models.Portfolio

func configureFlags(api *operations.ExchangeAppAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ExchangeAppAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	traders = make(map[string]models.Portfolio)

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-TOKEN" header is set
	api.AccessTokenAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (accessToken) X-TOKEN from header param [X-TOKEN] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.AssetAssetListHandler = asset.AssetListHandlerFunc(func(params asset.AssetListParams) middleware.Responder {
		return middleware.NotImplemented("operation asset.AssetList has not yet been implemented")
	})
	api.TraderCreateNewtraderHandler = trader.CreateNewtraderHandlerFunc(func(params trader.CreateNewtraderParams) middleware.Responder {
		token := "token_test"

		traders[token] = models.Portfolio{}

		return trader.NewCreateNewtraderOK().WithPayload(token)
	})
	api.OrderProcessOrderHandler = order.ProcessOrderHandlerFunc(func(params order.ProcessOrderParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation order.ProcessOrder has not yet been implemented")
	})
	api.TraderTraderPortfolioHandler = trader.TraderPortfolioHandlerFunc(func(params trader.TraderPortfolioParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation trader.TraderPortfolio has not yet been implemented")
	})
	api.TraderTraderPortfoliosHandler = trader.TraderPortfoliosHandlerFunc(func(params trader.TraderPortfoliosParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation trader.TraderPortfolios has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
