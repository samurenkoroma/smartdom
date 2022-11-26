package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"smartdom/pkg/store/postgres"
	deliveryHttp "smartdom/services/contact/internal/delivery/http"
	repositoryStorage "smartdom/services/contact/internal/repository/storage/postgres"
	useCaseContact "smartdom/services/contact/internal/useCase/contact"
	useCaseGroup "smartdom/services/contact/internal/useCase/group"
	"syscall"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	repoStorage, err := repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
	if err != nil {
		panic(err)
	}
	var (
		ucContact    = useCaseContact.New(repoStorage, useCaseContact.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		listenerHttp = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d\n", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			fmt.Print(err)
		}
	}()

	waitSignal()
}

func waitSignal() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
