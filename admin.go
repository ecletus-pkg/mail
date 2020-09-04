package mail

import (
	"github.com/ecletus/admin"
)

func AddMailSubResource(setup func(res *admin.Resource), res *admin.Resource, value interface{}, fieldName ...string) error {
	return res.GetAdmin().OnResourcesAdded(func(e *admin.ResourceEvent) error {
		cfg := &admin.Config{Name: fieldName[0], Setup: func(r *admin.Resource) {
			r.SetI18nModel(&Mail{})
			PrepareMailResource(r)
			res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
			if setup != nil {
				setup(r)
			}
		}}

		if len(fieldName) == 0 || fieldName[0] == "" {
			fieldName = []string{"Mails"}
			res.Meta(&admin.Meta{
				Name:  fieldName[0],
				Label: e.Resource.PluralLabelKey(),
			})
		} else {
			cfg.LabelKey = res.ChildrenLabelKey(fieldName[0])
		}
		res.NewResource(&admin.SubConfig{FieldName: fieldName[0]}, value, cfg)
		return nil
	}, ResourceID)
}

func PrepareMailResource(res *admin.Resource) {
	res.EditAttrs(&admin.Section{Rows: [][]string{{"Address", "Note"}}})
	res.ShowAttrs(res.EditAttrs())
	res.NewAttrs(res.EditAttrs())
	res.IndexAttrs("Address", "Note")
}

const ResourceID = "Mail"
