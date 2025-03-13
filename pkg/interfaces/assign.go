// Code generated by internal/generate/api/main.go; DO NOT EDIT.

package interfaces

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var AssignOpts = api.ReqOpts{
	AddEndpoint:         "/interfaces/assign_settings/addItem",
	GetEndpoint:         "/interfaces/assign_settings/getItem",
	UpdateEndpoint:      "/interfaces/assign_settings/setItem",
	DeleteEndpoint:      "/interfaces/assign_settings/delItem",
	ReconfigureEndpoint: "/interfaces/assign_settings/reconfigure",
	Monad:               "assign",
}

// Data structs

type Assign struct {
	Interface        string `json:"interface"`
	Device           string `json:"device"`
	Ip               string `json:"ipaddr"`
	Subnet           string `json:"subnet"`
	Gateway          string `json:"gateway"`
	GatewayInterface string `json:"gateway_interface"`
	Enable           string `json:"enable"`
	Description      string `json:"descr"`
}

// CRUD operations

func (c *Controller) AddAssign(ctx context.Context, resource *Assign) (string, error) {
	return api.Add(c.Client(), ctx, AssignOpts, resource)
}

func (c *Controller) GetAssign(ctx context.Context, id string) (*Assign, error) {
	return api.Get(c.Client(), ctx, AssignOpts, &Assign{}, id)
}

func (c *Controller) UpdateAssign(ctx context.Context, id string, resource *Assign) error {
	return api.Update(c.Client(), ctx, AssignOpts, resource, id)
}

func (c *Controller) DeleteAssign(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, AssignOpts, id)
}
