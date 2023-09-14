package main

import (
	"github.com/huantt/article-listing/cmd"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

func main() {
	var loggingLevel = new(slog.LevelVar)
	loggingLevel.Set(slog.LevelInfo)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: loggingLevel}))
	slog.SetDefault(logger)

	root := &cobra.Command{}
	root.AddCommand(cmd.UpdateArticles("update-articles"))
	err := root.Execute()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
