package uwcli

import (
	guuid "github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
	"time"
)

func (uc *UwCli) AddDeviceCmd() *cli.Command {
	return &cli.Command{
		Name:  "add-device",
		Usage: "Добавить новое устройство",
		Action: func(c *cli.Context) error {
			uid := guuid.New().String()
			err := uc.storage.CreateDevice(c.Context, uid)
			if err == nil {
				log.Info().Str("uid=", uid).Msg("new device created")
			}
			return err
		},
	}
}

func (uc *UwCli) UpdateActiveCmd() *cli.Command {
	return &cli.Command{
		Name:  "set-active",
		Usage: "Установить Active=true/false",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "uuid", Required: true},
			&cli.BoolFlag{Name: "active", Required: true},
		},
		Action: func(c *cli.Context) error {
			return uc.storage.SetActive(c.Context, c.String("uuid"), c.Bool("value"))
		},
	}
}

func (uc *UwCli) UpdateExpiresCmd() *cli.Command {
	return &cli.Command{
		Name:  "set-expiry",
		Usage: "Обновить ExpiresAt устройства",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "uuid", Required: true},
			&cli.StringFlag{Name: "date", Required: true, Usage: "в формате 2006-01-02"},
		},
		Action: func(c *cli.Context) error {
			t, err := time.Parse("2006-01-02", c.String("date"))
			if err != nil {
				return err
			}
			return uc.storage.UpdateExpires(c.Context, c.String("uuid"), t)
		},
	}
}

func (uc *UwCli) AssignGroupsCmd() *cli.Command {
	return &cli.Command{
		Name:  "assign-groups",
		Usage: "Назначить группы устройству",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "uuid", Required: true},
			&cli.StringFlag{Name: "groups", Usage: "-123456,1234567,12345678"},
		},
		Action: func(c *cli.Context) error {
			raw := strings.Split(c.String("groups"), ",")
			var ids []int64
			for _, s := range raw {
				id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}
			return uc.storage.AssignGroups(c.Context, c.String("uuid"), ids)
		},
	}
}

func (uc *UwCli) AddGroupCmd() *cli.Command {
	return &cli.Command{
		Name:  "add-group",
		Usage: "Добавить группу",
		Flags: []cli.Flag{
			&cli.Int64Flag{Name: "chatid", Required: true},
		},
		Action: func(c *cli.Context) error {
			return uc.storage.CreateGroup(c.Context, c.Int64("chatid"))
		},
	}
}

func (uc *UwCli) UpdateDeviceInfoCmd() *cli.Command {
	return &cli.Command{
		Name:  "set-info",
		Usage: "Обновить Label и Point",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "uuid", Required: true},
			&cli.StringFlag{Name: "label"},
			&cli.StringFlag{Name: "point"},
		},
		Action: func(c *cli.Context) error {
			uuid := c.String("uuid")
			label := c.String("label")
			if label == "" {
				label = uuid
			}
			point := c.String("point")
			if point == "" {
				point = DefaultPointName
			}
			return uc.storage.UpdateInfo(c.Context, uuid, label, point)
		},
	}
}

func (uc *UwCli) GetDevice() *cli.Command {
	return &cli.Command{
		Name:  "get-device",
		Usage: "Получить информацию про устройство",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "uuid", Required: true},
		},
		Action: func(c *cli.Context) error {
			uuid := c.String("uuid")
			device, err := uc.storage.GetDevice(c.Context, uuid)
			if err == nil {
				log.Info().Interface("device", device).Msg("information about device")
			}
			return err
		},
	}
}
