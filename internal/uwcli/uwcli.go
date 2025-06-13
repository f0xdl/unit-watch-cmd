package uwcli

import (
	"context"
	"github.com/f0xdl/unit-watch-lib/domain"
	"github.com/urfave/cli/v2"
	"time"
)

const DefaultPointName = "default"

type IStorage interface {
	CreateDevice(ctx context.Context, uid string) error
	SetActive(ctx context.Context, uid string, active bool) error
	UpdateExpires(ctx context.Context, uid string, t time.Time) error
	UpdateInfo(ctx context.Context, uid, label, point string) error
	CreateGroup(ctx context.Context, chatID int64) error
	GetDevice(ctx context.Context, uid string) (*domain.Device, error)
	AssignGroups(ctx context.Context, uid string, chatIds []int64) error
}

type UwCli struct {
	storage IStorage
}

func NewUwCli(storage IStorage) *UwCli {
	return &UwCli{storage: storage}
}

func (uc *UwCli) BuildCommands() []*cli.Command {
	return []*cli.Command{
		uc.AddDeviceCmd(),
		uc.UpdateActiveCmd(),
		uc.UpdateExpiresCmd(),
		uc.AddGroupCmd(),
		uc.UpdateDeviceInfoCmd(),
		uc.AssignGroupsCmd(),
		uc.GetDevice(),
	}
}

func NewUWCli(storage IStorage) *UwCli {
	return &UwCli{storage: storage}
}
