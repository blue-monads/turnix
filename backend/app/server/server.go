package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/bornjre/turnix/backend/controller"
	"github.com/bornjre/turnix/backend/controller/auth"
	"github.com/bornjre/turnix/backend/controller/common"
	"github.com/bornjre/turnix/backend/controller/project"
	"github.com/bornjre/turnix/backend/controller/self"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
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

	localSocket string
	ldListener  net.Listener

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
	LocalSocket     string
}

func New(opts Options) *Server {

	out, err := project.BuildGlobalJS(opts.ProjectBuilders)
	if err != nil {
		panic(err)
	}

	return &Server{
		db:          opts.DB,
		signer:      opts.Signer,
		rootLogger:  zerolog.New(os.Stdout).With().Str("service", "server").Logger(),
		projects:    opts.ProjectBuilders,
		globalJS:    out,
		cAuth:       opts.Controller.GetAuthController(),
		cProject:    opts.Controller.GetProjectController(),
		cSelf:       opts.Controller.GetSelfController(),
		cCommon:     opts.Controller.GetCommonController(),
		localSocket: opts.LocalSocket,
	}
}

func (a *Server) Start(port string) error {

	a.listenUnixSocket(port)

	r := gin.Default()

	r.Use(compress.Compress())
	r.Use(limits.RequestSizeLimiter(100 * 1024 * 1024))

	a.bindRoutes(r)

	go func() {
		time.Sleep(time.Second * 2)

		pp.Println(fmt.Sprintf("http://localhost%s/z/pages", port))

	}()

	return r.Run(port)
}

type LocalStatus struct {
	Port string `json:"port"`
}

func (s *Server) listenUnixSocket(port string) error {

	pp.Println("@listen_unix_socket", s.localSocket)

	// delete old socket

	os.Remove(s.localSocket)

	l, err := net.Listen("unix", s.localSocket)
	if err != nil {
		log.Println("listen_unix_socket error:", err.Error())
		return err
	}

	s.ldListener = l

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err.Error())
				return
			}

			func(c net.Conn) {
				defer c.Close()

				msg := LocalStatus{
					Port: port,
				}

				out, err := json.Marshal(msg)
				if err != nil {
					log.Fatal("json marshal error:", err.Error())
					return
				}

				_, err = c.Write(out)
				if err != nil {
					log.Fatal("Write: ", err)
				}

			}(c)

		}

	}()

	return nil
}
