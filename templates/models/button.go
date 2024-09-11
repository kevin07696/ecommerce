package models

type Button struct {
	name string
	opts *ButtonOpts
}

type ButtonOpts struct {
	style    string
	template string
	typ      string
	ropts    *RequestOpts
}

type ButtonOptsFunc func(*ButtonOpts)

func DefaultButtonOpts() *ButtonOpts {
	return &ButtonOpts{
		template: "basic",
		typ:      "submit",
		ropts:    &RequestOpts{},
	}
}

func NewButton(name string, opts ...ButtonOptsFunc) Button {
	dopts := DefaultButtonOpts()
	for _, fn := range opts {
		fn(dopts)
	}
	return Button{name: name, opts: dopts}
}

func WithBtnRequest(opts ...RequestOptFunc) ButtonOptsFunc {
	return func(bo *ButtonOpts) {
		bo.ropts = NewRequest(opts...)
	}
}

func WithBtnStyle(style string) ButtonOptsFunc {
	return func(bo *ButtonOpts) {
		bo.style = style
	}
}

func WithBtnTyp(typ string) ButtonOptsFunc {
	return func(bo *ButtonOpts) {
		bo.typ = typ
	}
}

func WithBtnTempl(template string) ButtonOptsFunc {
	return func(bo *ButtonOpts) {
		bo.template = template
	}
}

func (button Button) GetName() string {
	return button.name
}

func (button Button) GetOpts() *ButtonOpts {
	return button.opts
}

func (opts *ButtonOpts) GetTyp() string {
	return opts.typ
}

func (opts *ButtonOpts) GetStyle() string {
	return opts.style
}

func (opts *ButtonOpts) GetTemplate() string {
	return opts.template
}

func (opts *ButtonOpts) GetRequestOpts() *RequestOpts {
	return opts.ropts
}
