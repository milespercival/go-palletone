/*
   This file is part of go-palletone.
   go-palletone is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-palletone is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * @author PalletOne core developer Albert·Gou <dev@pallet.one>
 * @date 2018
 */

package mediatorplugin

import (
	"github.com/dedis/kyber"
	"github.com/palletone/go-palletone/common"
	"github.com/palletone/go-palletone/common/log"
	"github.com/palletone/go-palletone/core"
	"gopkg.in/urfave/cli.v1"
)

const (
	DefaultPassword              = "1"
	DefaultInitPrivKey           = "47gsj9pK3pwYUS1ZrWQjTgWMHUXWdNuCr7hXPXHySyBk"
	DefaultRequiredParticipation = 33
)

var (
	MediatorFlags = []cli.Flag{
		NoProduceUnitFlag,
		StaleProductionFlag,
		ConsecutiveProductionFlag,
		RequiredParticipationFlag,
		NoGroupSignFlag,
	}

	NoProduceUnitFlag = cli.BoolFlag{
		Name:  "noProduce",
		Usage: "Disable producing unit when start up node.",
	}
	StaleProductionFlag = cli.BoolFlag{
		Name:  "staleProduce",
		Usage: "Enable unit production, even if the chain is stale.",
	}
	ConsecutiveProductionFlag = cli.BoolFlag{
		Name:  "allowConsecutive",
		Usage: "Enable unit production, even if the last unit was generated by the same mediator.",
	}
	RequiredParticipationFlag = cli.UintFlag{
		Name:  "requiredParticipation",
		Usage: "Percent of mediators (0-99) that must be participating in order to produce units.",
		Value: DefaultRequiredParticipation,
	}
	NoGroupSignFlag = cli.BoolFlag{
		Name:  "noGroupSign",
		Usage: "Disable group-signing in this node.",
	}
)

// config data for mediator plugin
type Config struct {
	// 主程序启动时，是否立即开启unit生产
	EnableProducing bool

	// Enable Unit production, even if the chain is stale. 运行本节点开始生产unit，即使数据不是最新的
	EnableStaleProduction bool

	// Enable Unit production, even if the last unit was generated by the same mediator.
	// 允许本节点的mediator可以连续生产unit
	EnableConsecutiveProduction bool

	// Percent of mediators (0-99) that must be participating in order to produce uints
	RequiredParticipation uint32

	// 标记本节点是否开启群签名的功能
	EnableGroupSigning bool

	Mediators []*MediatorConf // the set of mediator accounts controlled by this node
}

func DefaultMediatorConf() *MediatorConf {
	return &MediatorConf{
		core.DefaultMediator,
		DefaultPassword,
		DefaultInitPrivKey,
		core.DefaultInitPubKey,
	}
}

// mediator plugin default config
var DefaultConfig = Config{
	EnableProducing:             false,
	EnableStaleProduction:       false,
	EnableConsecutiveProduction: false,
	RequiredParticipation:       DefaultRequiredParticipation,
	EnableGroupSigning:          true,
	Mediators: []*MediatorConf{
		DefaultMediatorConf(),
	},
}

func MakeConfig() Config {
	cfg := DefaultConfig
	cfg.Mediators = nil
	return cfg
}

func SetMediatorConfig(ctx *cli.Context, cfg *Config) {
	if ctx.GlobalIsSet(NoProduceUnitFlag.Name) {
		cfg.EnableProducing = false
	}
	if ctx.GlobalIsSet(StaleProductionFlag.Name) {
		cfg.EnableStaleProduction = true
	}
	if ctx.GlobalIsSet(ConsecutiveProductionFlag.Name) {
		cfg.EnableConsecutiveProduction = true
	}
	if ctx.GlobalIsSet(RequiredParticipationFlag.Name) {
		cfg.RequiredParticipation = uint32(ctx.GlobalUint(RequiredParticipationFlag.Name))
	}
	if ctx.GlobalIsSet(NoGroupSignFlag.Name) {
		cfg.EnableGroupSigning = false
	}
}

type MediatorConf struct {
	Address,
	Password,
	InitPrivKey,
	InitPubKey string
}

func (medConf *MediatorConf) configToAccount() *MediatorAccount {
	// 1. 解析 mediator 账户地址
	addr, err := core.StrToMedAdd(medConf.Address)
	if err != nil {
		log.Debugf(err.Error())
		return nil
	}

	// 2. 解析 mediator 的 DKS 初始公私钥
	sec, _ := core.StrToScalar(medConf.InitPrivKey)
	pub, _ := core.StrToPoint(medConf.InitPubKey)

	medAcc := &MediatorAccount{
		addr,
		medConf.Password,
		sec,
		pub,
	}

	return medAcc
}

type MediatorAccount struct {
	Address     common.Address
	Password    string
	InitPrivKey kyber.Scalar
	InitPubKey  kyber.Point
}
