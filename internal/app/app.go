package app

import (
	"context"
	"github.com/VeneLooool/BookHub/internal/config"
	grpc2 "github.com/VeneLooool/BookHub/internal/delivery/grpc"
	"github.com/VeneLooool/BookHub/internal/delivery/telegram"
	"github.com/VeneLooool/BookHub/internal/entity"
	desc "github.com/VeneLooool/BookHub/internal/pb"
	"github.com/VeneLooool/BookHub/internal/service/usecase"
	"github.com/VeneLooool/BookHub/internal/storage"
	"github.com/VeneLooool/BookHub/internal/storage/cache/memcache"
	"github.com/VeneLooool/BookHub/internal/storage/filemanager"
	postgres2 "github.com/VeneLooool/BookHub/internal/storage/postgres"
	"github.com/VeneLooool/BookHub/pkg/db/postgres"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"log"
	"net"
	"net/http"
)

const configPath = "./internal/config/config.yaml"

func Run() {
	viper, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	conf, err := config.ParseConfig(viper)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	dbConnection, err := postgres.NewPsqlDB(conf)
	if err != nil {
		log.Fatalf("NewPsqlDB: %v", err)
	}

	userCache := memcache.New[string, entity.User](&conf.Memcached)
	repoCache := memcache.New[string, entity.Repo](&conf.Memcached)
	bookCache := memcache.New[string, entity.Book](&conf.Memcached)

	fileManager := filemanager.NewFileManager("./fileManager")

	userStorage := postgres2.NewUserStorage(dbConnection)
	repoStorage := postgres2.NewRepoStorage(dbConnection)
	bookStorage := postgres2.NewBookStorage(dbConnection)

	userStorageAbs := storage.NewUserStorageAbs(userStorage, userCache)
	repoStorageAbs := storage.NewRepoStorageAbs(repoStorage, repoCache)
	bookStorageAbs := storage.NewBookStorageAbs(bookStorage, bookCache)

	userUseCase := usecase.NewUserService(userStorageAbs)
	repoUseCase := usecase.NewRepoService(repoStorageAbs)
	bookUseCase := usecase.NewBookService(bookStorageAbs, fileManager)

	service := grpc2.NewService(userUseCase, repoUseCase, bookUseCase)

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	desc.RegisterBookHubServiceServer(grpcServer, service)
	log.Println("Server started")
	go runRest(conf)

	log.Fatal(grpcServer.Serve(listen))

}

func runTgBot(config *config.Config) {
	telegram, err := telegram.NewTelegramBot(config.TelegramBot)
	if err != nil {
		panic(err)
	}

	if err = telegram.StartTelegramBot(); err != nil {
		panic(err)
	}
}

func runRest(config *config.Config) {
	mux := runtime.NewServeMux()
	err := desc.RegisterBookHubServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)

	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: mux,
	}
	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
