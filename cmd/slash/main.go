package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/boojack/slash/internal/log"
	"github.com/boojack/slash/server"
	_profile "github.com/boojack/slash/server/profile"
	"github.com/boojack/slash/store"
	"github.com/boojack/slash/store/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

const (
	greetingBanner = `Welcome to Slash!`
)

var (
	profile *_profile.Profile
	mode    string
	port    int
	data    string

	rootCmd = &cobra.Command{
		Use:   "slash",
		Short: `An open source, self-hosted bookmarks and link sharing platform.`,
		Run: func(_cmd *cobra.Command, _args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			db := db.NewDB(profile)
			if err := db.Open(ctx); err != nil {
				cancel()
				log.Error("failed to open database", zap.Error(err))
				return
			}

			storeInstance := store.New(db.DBInstance, profile)
			s, err := server.NewServer(ctx, profile, storeInstance)
			if err != nil {
				cancel()
				log.Error("failed to create server", zap.Error(err))
				return
			}

			c := make(chan os.Signal, 1)
			// Trigger graceful shutdown on SIGINT or SIGTERM.
			// The default signal sent by the `kill` command is SIGTERM,
			// which is taken as the graceful shutdown signal for many systems, eg., Kubernetes, Gunicorn.
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				sig := <-c
				log.Info(fmt.Sprintf("%s received.\n", sig.String()))
				s.Shutdown(ctx)
				cancel()
			}()

			printGreetings()

			if err := s.Start(ctx); err != nil {
				if err != http.ErrServerClosed {
					log.Error("failed to start server", zap.Error(err))
					cancel()
				}
			}

			// Wait for CTRL-C.
			<-ctx.Done()
		},
	}
)

func Execute() error {
	defer log.Sync()
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "demo", `mode of server, can be "prod" or "dev" or "demo"`)
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8082, "port of server")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "data directory")

	err := viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("data", rootCmd.PersistentFlags().Lookup("data"))
	if err != nil {
		panic(err)
	}

	viper.SetDefault("mode", "demo")
	viper.SetDefault("port", 8082)
	viper.SetEnvPrefix("slash")
}

func initConfig() {
	viper.AutomaticEnv()
	var err error
	profile, err = _profile.GetProfile()
	if err != nil {
		log.Error("failed to get profile", zap.Error(err))
		return
	}

	println("---")
	println("Server profile")
	println("dsn:", profile.DSN)
	println("port:", profile.Port)
	println("mode:", profile.Mode)
	println("version:", profile.Version)
	println("---")
}

func printGreetings() {
	fmt.Println(greetingBanner)
	fmt.Printf("Version %s has been started on port %d\n", profile.Version, profile.Port)
	fmt.Println("---")
	fmt.Println("See more in:")
	fmt.Printf("👉GitHub: %s\n", "https://github.com/boojack/slash")
	fmt.Println("---")
}

func main() {
	err := Execute()
	if err != nil {
		panic(err)
	}
}
