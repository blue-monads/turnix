package server

import (
	"os"

	"github.com/bornjre/turnix/backend/controller"
	"github.com/bornjre/turnix/backend/controller/project"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	db         *database.DB
	signer     *signer.Signer
	flakeNode  *snowflake.Node
	globalJS   []byte
	projects   map[string]*xproject.Defination
	rootLogger zerolog.Logger
	controller *controller.RootController
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

	node, err := snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}

	return &Server{
		db:         opts.DB,
		signer:     opts.Signer,
		flakeNode:  node,
		rootLogger: zerolog.New(os.Stdout).With().Str("service", "server").Logger(),
		projects:   opts.ProjectBuilders,
		globalJS:   out,
		controller: opts.Controller,
	}
}

func (a *Server) Start(port string) error {

	r := gin.Default()

	a.bindRoutes(r)

	return r.Run(port)
}
