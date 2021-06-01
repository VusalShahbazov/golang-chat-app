package api

import (
	"back/internal/models"
	"back/internal/services/migrations"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	Engine *gin.Engine
	Config *Config
}

func NewApiServer() *Server {

	engine := gin.Default()



	//conf := cors.DefaultConfig()
	//conf.AllowAllOrigins = true
	//engine.Use(cors.New(conf))
	return &Server{
		Engine: engine,
		Config: NewConfig(),
	}
}

func (apiServer *Server) Run() error {



	models.SetUp()


	apiServer.InitRoutes()


	if err := migrations.Run(os.Getenv("MIGRATION_DIR")); err != nil {
		return err
	}




	return apiServer.Engine.Run(apiServer.Config.BindAddr)
}





