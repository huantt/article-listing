package cmd

import (
	"context"
	"github.com/huantt/article-listing/handler/collector"
	"github.com/huantt/article-listing/impl/article_service/forem"
	foremsrv "github.com/huantt/article-listing/pkg/forem"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

func UpdateWeather(use string) *cobra.Command {
	var templateFilePath string
	var outputFilePath string
	var username string
	var limit int

	command := &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			devToService := forem.NewService(foremsrv.NewService(foremsrv.DevToEndpoint))
			handler := collector.NewHandler(devToService)
			err := handler.Collect(context.Background(), username, limit, templateFilePath, outputFilePath)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			slog.Info("Updated!")
		},
	}

	command.Flags().StringVarP(&templateFilePath, "template-file", "f", "", "Readme template file path")
	command.Flags().StringVarP(&outputFilePath, "out-file", "o", "", "Output file path")
	command.Flags().StringVarP(&username, "username", "u", "", "Username")
	command.Flags().IntVar(&limit, "limit", 5, "Limit number of articles")
	err := command.MarkFlagRequired("username")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("template-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("out-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return command
}
