package route

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/jaehoonkim/sentinel/pkg/manager/config"
	"github.com/jaehoonkim/sentinel/pkg/manager/control"
	"github.com/jaehoonkim/sentinel/pkg/manager/database/vanilla/excute"
	"github.com/jaehoonkim/sentinel/pkg/manager/macro/echoutil"
	"github.com/jaehoonkim/sentinel/pkg/manager/macro/logs"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func XAuthToken(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if cfg.Host.XAuthToken {
				// x_auth_token
				const (
					key   = "x_auth_token"
					value = "SENTINEL"
				)
				header_value := c.Request().Header.Get(key)

				if len(header_value) == 0 || strings.Compare(header_value, value) != 0 {
					err = errors.Errorf("not found request header%s",
						logs.KVL(
							"key", key,
						))
				}
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized).SetInternal(errors.Wrapf(err,
						http.StatusText(http.StatusUnauthorized)))
				}
				return
			}

			if err := next(c); err != nil {
				return err
			}

			return
		}
	}
}

func AgentSessionToken(db *sql.DB, dialect excute.SqlExcutor) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if _, err := control.GetAgentSessionClaims(c, db, dialect); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized).SetInternal(errors.Wrapf(err,
					http.StatusText(http.StatusUnauthorized)))
			}

			if err := next(c); err != nil {
				return err
			}
			return
		}
	}
}

func ServiceAuthorizationBearerToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// service Authorization: Bearer
			if _, err := control.GetServiceAuthorizationClaims(c); err != nil {
				// set http header "WWW-Authenticate" "Bearer"
				echoutil.SeHttpHeader(c.Response().Header(), "WWW-Authenticate", "Bearer")

				return echo.NewHTTPError(http.StatusUnauthorized).SetInternal(errors.Wrapf(err,
					http.StatusText(http.StatusUnauthorized)))
			}

			if err := next(c); err != nil {
				return err
			}

			return
		}
	}
}
