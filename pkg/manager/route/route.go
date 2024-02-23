// @title MORPHEUS
// @version 0.0.1
// @description this is a morpheus manager.
// @contact.url https://github.com/jaehoonkim/morpheus
// @contact.email jaehoon@gmail.io
// @securityDefinitions.apikey ClientAuth
// @in header
// @name x-morpheus-agent-token
// @description token for client api
// @securityDefinitions.apikey XAuthToken
// @in header
// @name x_auth_token
// @description limit for access morpheus api
package route

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/jaehoonkim/morpheus/pkg/manager/config"
	"github.com/jaehoonkim/morpheus/pkg/manager/control"
	"github.com/jaehoonkim/morpheus/pkg/manager/database/vanilla/excute"
	mysqlFlavor "github.com/jaehoonkim/morpheus/pkg/manager/database/vanilla/excute/dialects/mysql"
	pprof "github.com/jaehoonkim/morpheus/pkg/manager/route/pprof"
	"github.com/jaehoonkim/morpheus/pkg/version"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/jaehoonkim/morpheus/pkg/manager/route/docs"
)

func init() {
	//swago docs version
	docs.SwaggerInfo.Version = version.Version
}

type Route struct {
	e *echo.Echo

	Port                   int32
	UseTls                 bool
	TlsCertificateFilename string
	TlsPrivateKeyFilename  string
}

func New(cfg *config.Config, db *sql.DB) *Route {

	e := echo.New()
	ctl := control.NewVanilla(db, excute.GetSqlExcutor(mysqlFlavor.Dialect()))

	//echo cors config
	e.Use(echoCORSConfig(cfg))

	if true {
		//request id generator
		e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Generator: func() func() string {
				var (
					id uint64
				)
				return func() string {
					id := atomic.AddUint64(&id, 1)
					return fmt.Sprintf("%d", id)
				}
			}(),
		}))
	}
	//logger
	if true {
		e.Use(echoServiceLogger(config.LoggerInfoOutput))
	}

	// enable request 'Content-Encoding' header handler
	if true {
		e.Use(middleware.Decompress())
	}

	// enable response 'Content-Encoding' header handler
	if true {
		e.Use(middleware.Gzip())
	}

	//echo error handler
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		echoErrorResponder(err, ctx)
		echoErrorLogger(err, ctx)
	}
	//echo recover
	e.Use(echoRecover())

	//swago
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// pprof
	pprof.Wrap("/debug/pprof", e)

	{
		// /client/auth*
		e.POST("/agent/auth", ctl.AuthAgent)
		e.POST("/morpheus/agent/auth", ctl.AuthAgent)

		group := e.Group("")
		// @Security ClientSessionToken
		group.Use(AgentSessionToken(db, excute.GetSqlExcutor(mysqlFlavor.Dialect())))
		// /client/service*
		group.GET("/agent/service", ctl.PollingService)
		group.PUT("/agent/service", ctl.UpdateService)
		group.GET("/morpheus/agent/service", ctl.PollingService)
		group.PUT("/morpheus/agent/service", ctl.UpdateService)
	}

	{
		group := e.Group("")
		group.Use(XAuthToken(cfg)) // @Security XAuthToken

		// /manager/auth*
		group.POST("/manager/tenant", ctl.Tenant)
		group.POST("/morpheus/manager/tenant", ctl.Tenant)

		// /manager/template_recipe*
		group.GET("/manager/template_recipe", ctl.FindTemplateRecipe)
		group.GET("/morpheus/manager/template_recipe", ctl.FindTemplateRecipe)

		// /manager/template*
		group.GET("/manager/template", ctl.FindTemplate)
		group.GET("/manager/template/:uuid", ctl.GetTemplate)
		group.GET("/morpheus/manager/template", ctl.FindTemplate)
		group.GET("/morpheus/manager/template/:uuid", ctl.GetTemplate)

		// /manager/template_command*
		group.GET("/manager/template_command", ctl.FindTemplateCommand)
		group.GET("/manager/template_command/:uuid", ctl.GetTemplateCommand)
		group.GET("/morpheus/manager/template_command", ctl.FindTemplateCommand)
		group.GET("/morpheus/manager/template_command/:uuid", ctl.GetTemplateCommand)

		// /manager/global_variables*
		group.GET("/manager/global_variables", ctl.FindGlobalVariables)
		group.GET("/manager/global_variables/:uuid", ctl.GetGlobalVariables)
		group.PUT("/manager/global_variables/:uuid", ctl.UpdateGlobalVariablesValue)
		group.GET("/morpheus/manager/global_variables", ctl.FindGlobalVariables)
		group.GET("/morpheus/manager/global_variables/:uuid", ctl.GetGlobalVariables)
		group.PUT("/morpheus/manager/global_variables/:uuid", ctl.UpdateGlobalVariablesValue)

	}

	{
		group := e.Group("")
		group.Use(ServiceAuthorizationBearerToken()) // @Security ServiceAuthorizationBearerToken

		// /manager/cluster*
		group.GET("/manager/cluster", ctl.FindCluster)
		group.GET("/manager/cluster/:uuid", ctl.GetCluster)
		group.POST("/manager/cluster", ctl.CreateCluster)
		group.PUT("/manager/cluster/:uuid", ctl.UpdateCluster)
		group.PUT("/manager/cluster/:uuid/polling/regular", ctl.UpdateClusterPollingRegular)
		group.PUT("/manager/cluster/:uuid/polling/smart", ctl.UpdateClusterPollingSmart)
		group.DELETE("/manager/cluster/:uuid", ctl.DeleteCluster)
		group.GET("/morpheus/manager/cluster", ctl.FindCluster)
		group.GET("/morpheus/manager/cluster/:uuid", ctl.GetCluster)
		group.POST("/morpheus/manager/cluster", ctl.CreateCluster)
		group.PUT("/morpheus/manager/cluster/:uuid", ctl.UpdateCluster)
		group.PUT("/morpheus/manager/cluster/:uuid/polling/regular", ctl.UpdateClusterPollingRegular)
		group.PUT("/morpheus/manager/cluster/:uuid/polling/smart", ctl.UpdateClusterPollingSmart)
		group.DELETE("/morpheus/manager/cluster/:uuid", ctl.DeleteCluster)

		// /manager/service*
		group.GET("/manager/service", ctl.FindService)
		group.GET("/manager/service/:uuid", ctl.GetService)
		group.POST("/manager/service", ctl.CreateService)
		group.GET("/manager/service/:uuid/result", ctl.GetServiceResult)
		group.GET("/morpheus/manager/service", ctl.FindService)
		group.GET("/morpheus/manager/service/:uuid", ctl.GetService)
		group.POST("/morpheus/manager/service", ctl.CreateService)
		group.GET("/morpheus/manager/service/:uuid/result", ctl.GetServiceResult)

		// /manager/session*
		group.GET("/manager/session", ctl.FindSession)
		group.GET("/manager/session/:uuid", ctl.GetSession)
		group.DELETE("/manager/session/:uuid", ctl.DeleteSession)
		group.GET("/manager/session/cluster/:cluster_uuid/alive", ctl.AliveClusterSession)
		group.GET("/morpheus/manager/session", ctl.FindSession)
		group.GET("/morpheus/manager/session/:uuid", ctl.GetSession)
		group.DELETE("/morpheus/manager/session/:uuid", ctl.DeleteSession)
		group.GET("/morpheus/manager/session/cluster/:cluster_uuid/alive", ctl.AliveClusterSession)

		// /manager/cluster_token*
		group.GET("/manager/cluster_token", ctl.FindClusterToken)
		group.GET("/manager/cluster_token/:uuid", ctl.GetClusterToken)
		group.PUT("/manager/cluster_token/:uuid/label", ctl.UpdateClusterTokenLabel)
		group.DELETE("/manager/cluster_token/:uuid", ctl.DeleteClusterToken)
		group.POST("/manager/cluster_token", ctl.CreateClusterToken)
		group.PUT("/manager/cluster_token/:uuid/refresh", ctl.RefreshClusterTokenTime)
		group.PUT("/manager/cluster_token/:uuid/expire", ctl.ExpireClusterToken)
		group.GET("/morpheus/manager/cluster_token", ctl.FindClusterToken)
		group.GET("/morpheus/manager/cluster_token/:uuid", ctl.GetClusterToken)
		group.PUT("/morpheus/manager/cluster_token/:uuid/label", ctl.UpdateClusterTokenLabel)
		group.DELETE("/morpheus/manager/cluster_token/:uuid", ctl.DeleteClusterToken)
		group.POST("/morpheus/manager/cluster_token", ctl.CreateClusterToken)
		group.PUT("/morpheus/manager/cluster_token/:uuid/refresh", ctl.RefreshClusterTokenTime)
		group.PUT("/morpheus/manager/cluster_token/:uuid/expire", ctl.ExpireClusterToken)

		// /manager/channels*
		group.POST("/manager/channels", ctl.CreateChannel)
		group.GET("/manager/channels", ctl.FindChannel)
		group.GET("/manager/channels/:uuid", ctl.GetChannel)
		group.PUT("/manager/channels/:uuid", ctl.UpdateChannel)
		group.DELETE("/manager/channels/:uuid", ctl.DeleteChannel)
		group.POST("/morpheus/manager/channels", ctl.CreateChannel)
		group.GET("/morpheus/manager/channels", ctl.FindChannel)
		group.GET("/morpheus/manager/channels/:uuid", ctl.GetChannel)
		group.PUT("/morpheus/manager/channels/:uuid", ctl.UpdateChannel)
		group.DELETE("/morpheus/manager/channels/:uuid", ctl.DeleteChannel)

		// /manager/channels/:uuid/notifiers/*
		group.GET("/manager/channels/:uuid/notifiers/edge", ctl.GetChannelNotifierEdge)
		group.PUT("/manager/channels/:uuid/notifiers/console", ctl.UpdateChannelNotifierConsole)
		group.PUT("/manager/channels/:uuid/notifiers/rabbitmq", ctl.UpdateChannelNotifierRabbitMq)
		group.PUT("/manager/channels/:uuid/notifiers/webhook", ctl.UpdateChannelNotifierWebhook)
		group.PUT("/manager/channels/:uuid/notifiers/slackhook", ctl.UpdateChannelNotifierSlackhook)
		group.GET("/morpheus/manager/channels/:uuid/notifiers/edge", ctl.GetChannelNotifierEdge)
		group.PUT("/morpheus/manager/channels/:uuid/notifiers/console", ctl.UpdateChannelNotifierConsole)
		group.PUT("/morpheus/manager/channels/:uuid/notifiers/rabbitmq", ctl.UpdateChannelNotifierRabbitMq)
		group.PUT("/morpheus/manager/channels/:uuid/notifiers/webhook", ctl.UpdateChannelNotifierWebhook)
		group.PUT("/morpheus/manager/channels/:uuid/notifiers/slackhook", ctl.UpdateChannelNotifierSlackhook)

		// /manager/channels/status
		group.GET("/manager/channels/status", ctl.FindChannelStatus)
		group.GET("/morpheus/manager/channels/status", ctl.FindChannelStatus)

		// /manager/channels/:uuid/status*
		group.GET("/manager/channels/:uuid/status", ctl.ListChannelStatus)
		group.DELETE("/manager/channels/:uuid/status/purge", ctl.PurgeChannelStatus)
		group.PUT("/manager/channels/:uuid/status/option", ctl.UpdateChannelStatusOption)
		group.GET("/manager/channels/:uuid/status/option", ctl.GetChannelStatusOption)
		group.GET("/morpheus/manager/channels/:uuid/status", ctl.ListChannelStatus)
		group.DELETE("/morpheus/manager/channels/:uuid/status/purge", ctl.PurgeChannelStatus)
		group.PUT("/morpheus/manager/channels/:uuid/status/option", ctl.UpdateChannelStatusOption)
		group.GET("/morpheus/manager/channels/:uuid/status/option", ctl.GetChannelStatusOption)

		// /manager/channels/:uuid/format*
		group.GET("/manager/channels/:uuid/format", ctl.GetChannelFormat)
		group.PUT("/manager/channels/:uuid/format", ctl.UpdateChannelFormat)
		group.GET("/morpheus/manager/channels/:uuid/format", ctl.GetChannelFormat)
		group.PUT("/morpheus/manager/channels/:uuid/format", ctl.UpdateChannelFormat)

	}

	return &Route{
		e:                      e,
		Port:                   cfg.Host.Port,
		UseTls:                 cfg.Host.TlsEnable,
		TlsCertificateFilename: cfg.Host.TlsCertificateFilename,
		TlsPrivateKeyFilename:  cfg.Host.TlsPrivateKeyFilename,
	}
}

func (r *Route) Start() error {
	if r.UseTls {
		crt, err := os.ReadFile(r.TlsCertificateFilename)
		if err != nil {
			err = errors.Wrapf(err, "faild to read tls certificate file name=%v", r.TlsCertificateFilename)
			return err
		}
		key, err := os.ReadFile(r.TlsPrivateKeyFilename)
		if err != nil {
			err = errors.Wrapf(err, "faild to read tls privateKey file name=%v", r.TlsPrivateKeyFilename)
			return err
		}

		return StartTLS(r.e, r.Port, crt, key)
	}
	return Start(r.e, r.Port)
}

func Start(e *echo.Echo, port int32) error {
	go func() {
		address := fmt.Sprintf(":%d", port)
		if err := e.Start(address); err != nil {
			e.Logger.Info("shut down the manager")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func StartTLS(e *echo.Echo, port int32, crt, key []byte) error {
	go func() {
		address := fmt.Sprintf(":%d", port)
		if err := e.StartTLS(address, crt, key); err != nil {
			e.Logger.Info("shut down the manager")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
