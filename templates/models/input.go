package models

import "fmt"

type Input struct {
	id    string
	name  string
	label string
	opts  InputOpts
}

func NewInput(id, name, label string, opts ...InputOptsFunc) Input {
	dopts := DefaultInputOpts()
	for _, fn := range opts {
		fn(&dopts)
	}
	return Input{id: id, name: name, label: label, opts: dopts}
}

func (input Input) GetId() string {
	return input.id
}

func (input Input) GetName() string {
	return input.name
}

func (input Input) GetOpts() *InputOpts {
	return &input.opts
}

func (input Input) GetLabel() string {
	return input.label
}

type InputOptsFunc func(*InputOpts)

type InputOpts struct {
	typ            string
	fill           string
	place          string
	val            string
	template       string
	hint           string
	pattern        string
	isStateful     bool
	isRequired     bool
	isSpellChecked bool
	isFocused      bool
	button         *Button
}

func DefaultInputOpts() InputOpts {
	return InputOpts{
		typ:            "text",
		fill:           "off",
		place:          "",
		pattern:        ".*",
		val:            "",
		template:       "",
		isStateful:     false,
		isRequired:     false,
		isSpellChecked: false,
		isFocused:      false,
		button:         nil,
	}
}

func WithInpTyp(typ string) InputOptsFunc {
	return func(io *InputOpts) {
		io.typ = typ
	}
}

func WithFill(fill string) InputOptsFunc {
	return func(io *InputOpts) {
		io.fill = fill
	}
}

func WithPlace(place string) InputOptsFunc {
	return func(io *InputOpts) {
		io.place = place
	}
}

func WithPattern(pattern string) InputOptsFunc {
	return func(io *InputOpts) {
		io.pattern = pattern
	}
}

func WithVal(val string) InputOptsFunc {
	return func(io *InputOpts) {
		io.val = val
	}
}

func WithInpTemplate(template string) InputOptsFunc {
	return func(io *InputOpts) {
		io.template = template
	}
}

func WithStateful() InputOptsFunc {
	return func(io *InputOpts) {
		io.isStateful = true
	}
}

func WithRequired() InputOptsFunc {
	return func(io *InputOpts) {
		io.isRequired = true
	}
}

func WithSpellcheck() InputOptsFunc {
	return func(io *InputOpts) {
		io.isSpellChecked = true
	}
}

func WithFocus() InputOptsFunc {
	return func(io *InputOpts) {
		io.isFocused = true
	}
}

func WithHint(hint string) InputOptsFunc {
	return func(io *InputOpts) {
		io.hint = hint
	}
}

func WithButton(name string, opts ...ButtonOptsFunc) InputOptsFunc {
	return func(io *InputOpts) {
		button := NewButton(name, opts...)
		io.button = &button
	}
}

func (io InputOpts) GetTyp() string {
	return io.typ
}

func (io InputOpts) GetFill() string {
	return io.fill
}

func (io InputOpts) GetPlace() string {
	return io.place
}

func (io InputOpts) GetPattern() string {
	return io.pattern
}

func (io InputOpts) GetVal() string {
	return io.val
}

func (io InputOpts) GetTemplate() string {
	return io.template
}

func (io InputOpts) GetStateful() string {
	return fmt.Sprintf("%t", io.isStateful)
}

func (io InputOpts) GetRequired() string {
	return fmt.Sprintf("%t", io.isRequired)
}

func (io InputOpts) GetSpellchecked() string {
	return fmt.Sprintf("%t", io.isSpellChecked)
}

func (io InputOpts) GetFocused() string {
	return fmt.Sprintf("%t", io.isFocused)
}

func (io InputOpts) GetHint() string {
	return io.hint
}

func (io InputOpts) GetButton() *Button {
	return io.button
}
