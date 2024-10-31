package cmd

import (
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt/v5"
	usersRepository "go-skeleton/internal/repositories/users"
	"go-skeleton/pkg/crypt"
	"go-skeleton/pkg/jwtExtractor"
	"go-skeleton/pkg/roles"
	"go-skeleton/pkg/signerVerifier"

	//{{codeGen5}}
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/validator"
)

var (
	Reg       *registry.Registry
	ApiPrefix string
)

func Setup() {
	conf := config.NewConfig()
	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	ApiPrefix = conf.ReadConfig("API_PREFIX")

	l := logger.NewLogger(
		conf.ReadConfig("ENVIRONMENT"),
		conf.ReadConfig("APP"),
		conf.ReadConfig("VERSION"),
	)

	l.Boot()

	db := database.NewMysql(
		l,
		conf.ReadConfig("DB_USER"),
		conf.ReadConfig("DB_PASS"),
		conf.ReadConfig("DB_URL"),
		conf.ReadConfig("DB_PORT"),
		conf.ReadConfig("DB_DATABASE"),
	)

	val := validator.NewValidator()

	db.Connect()
	val.Boot()

	idC := idCreator.NewIdCreator()

	encryptPkg := crypt.NewCrypt()

	sqlAdapter, adapterErr := sqladapter.NewAdapter(db.Db.DB, "mysql", "roles")
	if adapterErr != nil {
		panic(adapterErr)
	}

	enforcer, enforcerErr := casbin.NewEnforcer("pkg/roles/rbac.conf", sqlAdapter)
	if enforcerErr != nil {
		panic(enforcerErr)
	}
	rolesController := roles.NewCasbinRule(enforcer)

	jwtSignVerifierController := signerVerifier.NewSigner(conf.ReadConfig("JWT_SECRET"))
	jwtParser := jwt.NewParser()
	extractor := jwtExtractor.NewJWTExtractor(jwtParser)

	Reg = registry.NewRegistry()
	Reg.Provide("logger", l)
	Reg.Provide("validator", val)
	Reg.Provide("config", conf)
	Reg.Provide("idCreator", idC)
	Reg.Provide("crypt", encryptPkg)
	Reg.Provide("roles", rolesController)
	Reg.Provide("signerVerifier", jwtSignVerifierController)
	Reg.Provide("jwtExtractor", extractor)

	Reg.Provide("usersRepository", usersRepository.NewUsersRepository(db.Db))
	//{{codeGen6}}
}
