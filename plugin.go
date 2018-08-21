package mail

import (
	"github.com/aghape/admin"
	"github.com/aghape/admin/adminplugin"
	"github.com/aghape/db"
	"github.com/aghape/plug"
)

type Plugin struct {
	plug.EventDispatcher
	db.DBNames
	adminplugin.AdminNames
}

func (p *Plugin) OnRegister() {
	p.AdminNames.OnInitResources(p, func(e *adminplugin.AdminEvent) {
		e.Admin.AddResource(&Mail{}, &admin.Config{Setup: PrepareMailResource, Invisible: true})
	})
	db.Events(p).DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&Mail{}).Error
	})
}
