package base

type (
	Any = interface{}
	Map = map[string]Any

	Res interface {
		OK() bool
		Fail() bool
		Args() []Any
		With(args ...Any) Res
		Code() int
		Text() string
		Error() string
	}

	Vars map[string]Var
	Var  struct {
		nil      bool
		Type     string `json:"t"`
		Required bool   `json:"r"`
		Nullable bool   `json:"-"`
		Name     string `json:"n,omitempty"`
		Text     string `json:"t,omitempty"`
		Default  Any    `json:"-"`
		Setting  Map    `json:"-"`
		Children Vars   `json:"cs,omitempty"`
		Options  Map    `json:"os,omitempty"`

		Encode string `json:"enc,omitempty"`
		Decode string `json:"dec,omitempty"`

		Empty Res `json:"-"`
		Error Res `json:"-"`

		Valid func(Any, Var) bool `json:"-"`
		Value func(Any, Var) Any  `json:"-"`
	}
)

var (
	Nil = Var{nil: true}
)

func (vvv *Var) Nil() bool {
	if vvv == nil {
		return true
	}
	return vvv.nil
}
