package speedy

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/ghibbo/speedy/render"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

// Speedy App
type Speedy struct {
	AppName  string
	Debug    bool //dev o production
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	JetViews *jet.Set
	config   config
}

type config struct {
	port     string
	renderer string
}

// Path Application Start
func (s *Speedy) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := s.Init(pathConfig)
	if err != nil {
		return err
	}

	err = s.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	//read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	//create loggers
	infoLog, errorLog := s.startLoggers()

	// settings
	s.InfoLog = infoLog
	s.ErrorLog = errorLog
	s.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	s.Version = version
	s.RootPath = rootPath
	s.Routes = s.routes().(*chi.Mux)
	s.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	s.JetViews = views

	s.createRenderer()

	return nil
}

// Initialize Folder
func (s *Speedy) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder
		err := s.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListenAndServe starts the web server
func (s *Speedy) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     s.ErrorLog,
		Handler:      s.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	s.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	s.ErrorLog.Fatal(err)
}

func (s *Speedy) checkDotEnv(path string) error {
	err := s.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

func (s *Speedy) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (s *Speedy) createRenderer() {
	myRender := render.Render{
		Renderer: s.config.renderer,
		RootPath: s.RootPath,
		Port:     s.config.port,
		JetViews: s.JetViews,
	}
	s.Render = &myRender
}
