package actions

import (
	"context"

	"olympics/pkg/core/entities"
	"olympics/pkg/storage"
)

type Actions struct {
	storage storage.Storage
}

func NewActions(storage storage.Storage) *Actions {
	return &Actions{storage: storage}
}

func (a *Actions) GetRole(ctx context.Context, userName string) (entities.Role, error) {
	return a.storage.GetInfoByClient(ctx, userName)
}

func (a *Actions) AddRole(ctx context.Context, userName string, role entities.Role) error {
	return a.storage.AddRole(ctx, userName, role)
}

func (a *Actions) GetAthleteInfo(ctx context.Context, athleteName string) (entities.AthleteInfo, error) {
	athleteEntries, err := a.storage.GetAthleteInfo(ctx, athleteName)
	if err != nil {
		return entities.AthleteInfo{}, err
	}

	return entities.BuildAthleteInfo(athleteEntries), nil
}

func (a *Actions) AddAthleteEvent(ctx context.Context, event entities.Athlete) error {
	return a.storage.AddAthleteEvent(ctx, event)
}
