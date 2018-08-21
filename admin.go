package mail

import (
	"github.com/aghape/admin"
	"github.com/aghape/admin/admincommon"
)

func AddMailSubResource(res *admin.Resource, value interface{}, fieldName ...string) *admin.Resource {
	if len(fieldName) == 0 || fieldName[0] == "" {
		fieldName = []string{"Mails"}
	}
	r := res.NewResource(&admin.SubConfig{FieldName: fieldName[0]}, value, &admin.Config{Name: fieldName[0], Setup: func(r *admin.Resource) {
		r.SetI18nModel(&Mail{})
		PrepareMailResource(r)
	}})
	res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
	return r
}

func PrepareMailResource(res *admin.Resource) {
	admincommon.RecordInfoFields(res)
	res.EditAttrs(&admin.Section{Rows: [][]string{{"Address", "Note"}}})
	res.ShowAttrs(res.EditAttrs())
	res.NewAttrs(res.EditAttrs())
}

func GetResource(Admin *admin.Admin) *admin.Resource {
	return Admin.GetResourceByID("Mail")
}
