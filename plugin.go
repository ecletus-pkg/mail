package mail

import (
	"github.com/aghape/admin"
	"github.com/aghape/admin/adminplugin"
	"github.com/aghape/db"
	"github.com/moisespsena/go-edis"
)

type Plugin struct {
	edis.EventDispatcher
	db.DBNames
	adminplugin.AdminNames
}

func (p *Plugin) OnRegister() {
	p.AdminNames.OnInitResources(p, func(e *adminplugin.AdminEvent) {
		e.Admin.AddResource(&QorMail{}, &admin.Config{Setup: PrepareMailResource})
	})
	db.DisNames(p).DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&QorMail{}).Error
	})
}
