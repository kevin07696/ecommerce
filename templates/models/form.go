package models

type Form struct {
	id   string
	opts RequestOpts
}

func NewForm(id string, opts ...RequestOptFunc) *Form {
	dopts := RequestOpts{}
	for _, fn := range opts {
		fn(&dopts)
	}
	return &Form{id: id, opts: dopts}
}

func (form *Form) GetId() string {
	return form.id
}

func (form *Form) GetOpts() RequestOpts {
	return form.opts
}
