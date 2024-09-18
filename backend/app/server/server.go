package server

import (
	"os"

	"github.com/bornjre/turnix/backend/controller"
	"github.com/bornjre/turnix/backend/controller/auth"
	"github.com/bornjre/turnix/backend/controller/common"
	"github.com/bornjre/turnix/backend/controller/project"
	"github.com/bornjre/turnix/backend/controller/self"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/aurowora/compress"
	limits "github.com/gin-contrib/size"
)

type Server struct {
	db         *database.DB
	signer     *signer.Signer
	globalJS   []byte
	projects   map[string]*xproject.Defination
	rootLogger zerolog.Logger

	// controllers

	cAuth    *auth.AuthController
	cProject *project.ProjectController
	cSelf    *self.SelfController
	cCommon  *common.CommonController
}

type Options struct {
	DB              *database.DB
	Signer          *signer.Signer
	ProjectBuilders map[string]*xproject.Defination
	Controller      *controller.RootController
}

func New(opts Options) *Server {

	out, err := project.BuildGlobalJS(opts.ProjectBuilders)
	if err != nil {
		panic(err)
	}

	return &Server{
		db:         opts.DB,
		signer:     opts.Signer,
		rootLogger: zerolog.New(os.Stdout).With().Str("service", "server").Logger(),
		projects:   opts.ProjectBuilders,
		globalJS:   out,
		cAuth:      opts.Controller.GetAuthController(),
		cProject:   opts.Controller.GetProjectController(),
		cSelf:      opts.Controller.GetSelfController(),
		cCommon:    opts.Controller.GetCommonController(),
	}
}

func (a *Server) Start(port string) error {

	r := gin.Default()

	r.Use(compress.Compress())
	r.Use(limits.RequestSizeLimiter(100 * 1024 * 1024))

	a.bindRoutes(r)

	return r.Run(port)
}
