package example

import (

	"fmt"
	"strings"
	log "github.com/sirupsen/logrus"

	"github.com/longphu-thesis/go-gin-app/example/redis"
	"github.com/longphu-thesis/go-gin-app/example/stringutil"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix(`go_gin_app`)
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`yaml`)
	//viper.AddConfigPath(`./config/`)
	viper.SetConfigFile(`config.yml`)

	err := viper.ReadInConfig()

	if err != nil {
		log.Error(err)
	}

	if viper.GetBool(`debug`) {
		log.Info(fmt.Sprintf("Service RUN on DEBUG mode"))
	}

	//os.Setenv("GO_GIN_APP_MYSQL_HOST", "localhost")

	//log.Info(viper.AllKeys())
	//log.Info(viper.ConfigFileUsed())
	//log.Info(viper.Get(`mysql.host`))
}

func Run() {

	f := &Foo{
		FirstName: "FirstName",
		LastName:  "LastName",
		Age:       50,
	}

	f.reflect()

	Sum(1,1)

	redis.Connect()

	log.Info(fmt.Sprintf("%s",stringutil.Reverse("!oG ,olleH")))
}
