package mail

import (
	"github.com/ecletus/admin"
)

func AddMailSubResource(res *admin.Resource, value interface{}, fieldName ...string) *admin.Resource {
	cfg := &admin.Config{Name: fieldName[0], Setup: func(r *admin.Resource) {
		r.SetI18nModel(&Mail{})
		PrepareMailResource(r)
	}}

	if len(fieldName) == 0 || fieldName[0] == "" {
		fieldName = []string{"Mails"}
		res.Meta(&admin.Meta{
			Name:  fieldName[0],
			Label: GetResource(res.GetAdmin()).PluralLabelKey(),
		})
	} else {
		cfg.LabelKey = res.ChildrenLabelKey(fieldName[0])
	}

	r := res.NewResource(&admin.SubConfig{FieldName: fieldName[0]}, value, cfg)
	res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
	return r
}

func PrepareMailResource(res *admin.Resource) {
	res.EditAttrs(&admin.Section{Rows: [][]string{{"Address", "Note"}}})
	res.ShowAttrs(res.EditAttrs())
	res.NewAttrs(res.EditAttrs())
	res.IndexAttrs("Address", "Note")
}

func GetResource(Admin *admin.Admin) *admin.Resource {
	return Admin.GetResourceByID("Mail")
}
