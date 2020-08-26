package env

import (
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/mon"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
	"net/url"
)

func init() {
	componentFactories[componentMySql] = new(mysqlFactory)
}

const componentMySql = "mysql"

type mysqlCredentials struct {
	DatabaseName string `cfg:"database_name" default:"gosoline"`
	UserName     string `cfg:"user_name" default:"gosoline"`
	UserPassword string `cfg:"user_password" default:"gosoline"`
	RootPassword string `cfg:"root_password" default:"gosoline"`
}

type mysqlSettings struct {
	ComponentBaseSettings
	ComponentContainerSettings
	Port        int              `cfg:"port" default:"0"`
	Version     string           `cfg:"version" default:"8.0"`
	Credentials mysqlCredentials `cfg:"credentials"`
}

type mysqlFactory struct {
}

func (f mysqlFactory) Detect(config cfg.Config, manager *ComponentsConfigManager) error {
	if !config.IsSet("db") {
		return nil
	}

	if manager.HasType(componentMySql) {
		return nil
	}

	components := config.GetStringMap("db")

	for name := range components {
		driver := config.Get(fmt.Sprintf("db.%s.driver", name))

		if driver != componentMySql {
			continue
		}

		settings := &mysqlSettings{}
		config.UnmarshalDefaults(settings)

		settings.Type = componentMySql
		settings.Name = name

		tmpfs, err := cast.ToStringMapStringE(settings.TmpfsIface)
		if err != nil {
			return err
		}

		settings.Tmpfs = tmpfs

		if err := manager.Add(settings); err != nil {
			return fmt.Errorf("can not add default mysql component: %w", err)
		}
	}

	return nil
}

func (f mysqlFactory) GetSettingsSchema() ComponentBaseSettingsAware {
	return &mysqlSettings{}
}

func (f mysqlFactory) ConfigureContainer(settings interface{}) *containerConfig {
	s := settings.(*mysqlSettings)

	env := []string{
		fmt.Sprintf("MYSQL_DATABASE=%s", s.Credentials.DatabaseName),
		fmt.Sprintf("MYSQL_USER=%s", s.Credentials.UserName),
		fmt.Sprintf("MYSQL_PASSWORD=%s", s.Credentials.UserPassword),
		fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", s.Credentials.RootPassword),
	}

	return &containerConfig{
		Repository: "mysql/mysql-server",
		Tmpfs:      s.Tmpfs,
		Tag:        s.Version,
		Env:        env,
		Cmd:        []string{"--sql_mode=NO_ENGINE_SUBSTITUTION", "--log-bin-trust-function-creators=TRUE"},
		PortBindings: portBindings{
			"3306/tcp": s.Port,
		},
		ExpireAfter: s.ExpireAfter,
	}
}

func (f mysqlFactory) HealthCheck(settings interface{}) ComponentHealthCheck {
	return func(container *container) error {
		s := settings.(*mysqlSettings)
		binding := container.bindings["3306/tcp"]
		client, err := f.connection(s, binding)

		if err != nil {
			return fmt.Errorf("can not create client: %w", err)
		}

		return client.Ping()
	}
}

func (f mysqlFactory) Component(_ cfg.Config, _ mon.Logger, container *container, settings interface{}) (Component, error) {
	s := settings.(*mysqlSettings)
	binding := container.bindings["3306/tcp"]
	client, err := f.connection(s, binding)

	if err != nil {
		return nil, fmt.Errorf("can not create client: %w", err)
	}

	component := &mysqlComponent{
		baseComponent: baseComponent{
			name: s.Name,
		},
		client:      client,
		credentials: s.Credentials,
		binding:     binding,
	}

	return component, nil
}

func (f mysqlFactory) connection(settings *mysqlSettings, binding containerBinding) (*sqlx.DB, error) {
	dsn := url.URL{
		User: url.UserPassword(settings.Credentials.UserName, settings.Credentials.UserPassword),
		Host: fmt.Sprintf("tcp(%s:%v)", binding.host, binding.port),
		Path: settings.Credentials.DatabaseName,
	}

	qry := dsn.Query()
	qry.Set("multiStatements", "true")
	qry.Set("parseTime", "true")
	qry.Set("charset", "utf8mb4")
	dsn.RawQuery = qry.Encode()

	client, err := sqlx.Open("mysql", dsn.String()[2:])

	if err != nil {
		return nil, fmt.Errorf("can not create client: %w", err)
	}

	return client, nil
}
