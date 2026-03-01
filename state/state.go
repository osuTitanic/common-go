package state

import (
	"fmt"
	"log/slog"

	"github.com/osuTitanic/common-go/config"
	"github.com/osuTitanic/common-go/database"
	"github.com/osuTitanic/common-go/email"
	"github.com/osuTitanic/common-go/logging"
	"github.com/osuTitanic/common-go/storage"
	"gorm.io/gorm"
)

type State struct {
	*Repositories

	Config   *config.Config
	Database *gorm.DB
	Logger   *slog.Logger
	Storage  storage.Storage
	Email    email.Email
}

func NewState(environmentFiles ...string) (*State, error) {
	cfg, err := config.LoadConfig(environmentFiles...)
	if err != nil {
		return nil, fmt.Errorf("state: failed to load config: %w", err)
	}

	logLevel := slog.LevelInfo
	if cfg.Debug {
		logLevel = slog.LevelDebug
	}
	logging.SetDefault("titanic", logLevel)
	logger := slog.Default()

	fs := storage.NewFileStorage(cfg.DataPath)
	if err := fs.Setup(); err != nil {
		return nil, fmt.Errorf("state: failed to setup storage: %w", err)
	}

	db, err := database.CreateSession(cfg)
	if err != nil {
		return nil, fmt.Errorf("state: failed to create database session: %w", err)
	}

	mailer, err := email.NewEmailFromConfig(cfg)
	if err != nil {
		database.CloseSession(db)
		return nil, fmt.Errorf("state: failed to create email service: %w", err)
	}

	if err := mailer.Setup(); err != nil {
		database.CloseSession(db)
		return nil, fmt.Errorf("state: failed to setup email service: %w", err)
	}

	return &State{
		Repositories: NewRepositories(db),
		Config:       cfg,
		Database:     db,
		Logger:       logger,
		Storage:      fs,
		Email:        mailer,
	}, nil
}

func (state *State) Close() error {
	if state == nil || state.Database == nil {
		return nil
	}

	return database.CloseSession(state.Database)
}
