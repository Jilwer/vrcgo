package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jilwer/vrcgo/vrcbot/commands"
	"github.com/Jilwer/vrcgo/vrcbot/components"
	"github.com/Jilwer/vrcgo/vrcbot/handlers"

	vrcbot "github.com/Jilwer/vrcgo/vrcbot"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/handler"
)

var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	shouldSyncCommands := flag.Bool("sync-commands", false, "Whether to sync commands to discord")
	path := flag.String("config", "config.toml", "path to config")
	flag.Parse()

	cfg, err := vrcbot.LoadConfig(*path)
	if err != nil {
		slog.Error("Failed to read config", slog.Any("err", err))
		os.Exit(-1)
	}

	setupLogger(cfg.Log)
	slog.Info("Starting vrcgo...", slog.String("version", version), slog.String("commit", commit))
	slog.Info("Syncing commands", slog.Bool("sync", *shouldSyncCommands))

	b := vrcbot.New(*cfg, version, commit)

	h := handler.New()
	h.Command("/test", commands.TestHandler)
	h.Autocomplete("/test", commands.TestAutocompleteHandler)
	h.Command("/version", commands.VersionHandler(b))
	h.Component("/test-button", components.TestComponent)
	h.Command("/online", commands.OnlineHandler)
	h.Command("/config", commands.ConfigHandler)

	if err = b.SetupBot(h, bot.NewListenerFunc(b.OnReady), handlers.MessageHandler(b)); err != nil {
		slog.Error("Failed to setup vrcbot", slog.Any("err", err))
		os.Exit(-1)
	}

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		b.Client.Close(ctx)
	}()

	if *shouldSyncCommands {
		slog.Info("Syncing commands", slog.Any("guild_ids", cfg.Bot.DevGuilds))
		if err = handler.SyncCommands(b.Client, commands.Commands, cfg.Bot.DevGuilds); err != nil {
			slog.Error("Failed to sync commands", slog.Any("err", err))
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = b.Client.OpenGateway(ctx); err != nil {
		slog.Error("Failed to open gateway", slog.Any("err", err))
		os.Exit(-1)
	}

	slog.Info("Bot is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
	slog.Info("Shutting down vrcbot...")
}

func setupLogger(cfg vrcbot.LogConfig) {
	opts := &slog.HandlerOptions{
		AddSource: cfg.AddSource,
		Level:     cfg.Level,
	}

	var sHandler slog.Handler
	switch cfg.Format {
	case "json":
		sHandler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		sHandler = slog.NewTextHandler(os.Stdout, opts)
	default:
		slog.Error("Unknown log format", slog.String("format", cfg.Format))
		os.Exit(-1)
	}
	slog.SetDefault(slog.New(sHandler))
}
