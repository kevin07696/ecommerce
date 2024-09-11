package models

type RequestOptFunc func(*RequestOpts)

type RequestOpts struct {
	method string
	url    string
	swap   string
	target string
	params string
}

func WithRequest(method, url string) RequestOptFunc {
	return func(ro *RequestOpts) {
		ro.method = method
		ro.url = url
	}
}

func WithSwap(swap string) RequestOptFunc {
	return func(ro *RequestOpts) {
		ro.swap = swap
	}
}

func WithTarget(target string) RequestOptFunc {
	return func(ro *RequestOpts) {
		ro.target = target
	}
}

func WithParams(params string) RequestOptFunc {
	return func(ro *RequestOpts) {
		ro.params = params
	}
}

func NewRequest(opts ...RequestOptFunc) *RequestOpts {
	dopts := RequestOpts{params: "*"}
	for _, fn := range opts {
		fn(&dopts)
	}
	return &dopts
}
func (requestOpts RequestOpts) GetMethod() string {
	return requestOpts.method
}

func (requestopts RequestOpts) GetURL() string {
	return requestopts.url
}

func (requestopts RequestOpts) GetSwap() string {
	return requestopts.swap
}

func (requestopts RequestOpts) GetTarget() string {
	return requestopts.target
}

func (requestOpts RequestOpts) GetParams() string {
	return requestOpts.params
}
