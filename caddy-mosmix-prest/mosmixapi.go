package mosmixapi

import (
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"github.com/prest/adapters/postgres"
	prestconfig "github.com/prest/config"
	"github.com/prest/controllers"
)

type Handler struct {
	Next    httpserver.Handler
	Configs []*config
}

type config struct {
	cannedQueryPaths map[string]string
	pgHost           string
	pgPort           int
	pgUser           string
	pgPass           string
	pgDatabase       string
}

func init() {
	caddy.RegisterPlugin("mosmixapi", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func parseConfigs(c *caddy.Controller) ([]*config, error) {
	configs := []*config{}

	for c.Next() {
		conf := &config{
			cannedQueryPaths: make(map[string]string),
			pgPort:           5432,
		}

		args := c.RemainingArgs()

		if len(args) != 0 {
			return nil, c.ArgErr()
		}

		for c.NextBlock() {
			switch c.Val() {
			case "pg_host":
				if !c.NextArg() {
					return nil, c.ArgErr()
				}
				conf.pgHost = c.Val()
			case "pg_port":
				if !c.NextArg() {
					return nil, c.ArgErr()
				}
				pgPort, err := strconv.Atoi(c.Val())
				if err != nil {
					return nil, c.Errf("Unable to parse pg_port value %s to int. %s", c.Val(), err)
				}
				conf.pgPort = pgPort
			case "pg_user":
				if !c.NextArg() {
					return nil, c.ArgErr()
				}
				conf.pgUser = c.Val()
			case "pg_pass":
				if !c.NextArg() {
					return nil, c.ArgErr()
				}
				conf.pgPass = c.Val()
			case "pg_db":
				if !c.NextArg() {
					return nil, c.ArgErr()
				}
				conf.pgDatabase = c.Val()
			case "canned_queries":
				c.Next()
				for c.NextBlock() {
					url := c.Val()
					if !c.NextArg() {
						return nil, c.ArgErr()
					}

					conf.cannedQueryPaths[url] = c.Val()
				}
			}
		}

		if len(conf.cannedQueryPaths) > 0 {
			configs = append(configs, conf)
		}
	}

	return configs, nil
}

func setup(c *caddy.Controller) error {
	configs, err := parseConfigs(c)
	if err != nil {
		return err
	}
	prestconfig.PrestConf = &prestconfig.Prest{
		PGHost:     configs[0].pgHost,
		PGUser:     configs[0].pgUser,
		PGDatabase: configs[0].pgDatabase,
		PGPort:     configs[0].pgPort,
		PGPass:     configs[0].pgPass,
		SSLMode:    "disable",
	}
	postgres.Load()

	httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return Handler{
			Next:    next,
			Configs: configs,
		}
	})
	return nil
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	for url, directory := range h.Configs[0].cannedQueryPaths {
		if httpserver.Path(r.URL.Path).Matches(url) {
			result, err := controllers.ExecuteScriptQuery(r, directory, path.Base(r.URL.Path))
			if err != nil {
				if strings.HasPrefix(err.Error(), "could not get script") {
					return http.StatusNotFound, err
				}
				return http.StatusBadRequest, err
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(result)

			return 0, nil
		}
	}

	return h.Next.ServeHTTP(w, r)
}
