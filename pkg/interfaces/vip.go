// Code generated by internal/generate/api/main.go; DO NOT EDIT.

package interfaces

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var VipOpts = api.ReqOpts{
	AddEndpoint:         "/interfaces/vip_settings/addItem",
	GetEndpoint:         "/interfaces/vip_settings/getItem",
	UpdateEndpoint:      "/interfaces/vip_settings/setItem",
	DeleteEndpoint:      "/interfaces/vip_settings/delItem",
	ReconfigureEndpoint: "/interfaces/vip_settings/reconfigure",
	Monad:               "vip",
}

// Data structs

type Vip struct {
	Interface   api.SelectedMap `json:"interface"`
	Mode        api.SelectedMap `json:"mode"`
	Address     string          `json:"address"`
	Subnet      string          `json:"subnet"`
	SubnetBit   string          `json:"subnet_bits"`
	Description string          `json:"descr"`
	Gateway     string          `json:"gateway"`
	NoExpand    string          `json:"noexpand"`
	NoBind      string          `json:"nobind"`
	Password    string          `json:"password"`
	Vhid        string          `json:"vhid"`
	AdvBase     string          `json:"advbase"`
	AdvSkew     string          `json:"advskew"`
	Peer        string          `json:"peer"`
	Peer6       string          `json:"peer6"`
	NoSync      string          `json:"nosync"`
}

// CRUD operations

func (c *Controller) AddVip(ctx context.Context, resource *Vip) (string, error) {
	return api.Add(c.Client(), ctx, VipOpts, resource)
}

func (c *Controller) GetVip(ctx context.Context, id string) (*Vip, error) {
	return api.Get(c.Client(), ctx, VipOpts, &Vip{}, id)
}

func (c *Controller) UpdateVip(ctx context.Context, id string, resource *Vip) error {
	return api.Update(c.Client(), ctx, VipOpts, resource, id)
}

func (c *Controller) DeleteVip(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, VipOpts, id)
}
