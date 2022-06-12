package base

//

type (
	// Any 任意类型
	Any = interface{}
	// Map map对象
	Map = map[string]interface{}

	// Res 操作结果
	Res interface {
		// OK 函数返回操作是否成功，Res 对象为空，或是 Code() 为 0，表示成功，其它均表示失败。
		OK() bool

		// Fail 函数返回操作是否失败，与 OK() 正好相反。
		Fail() bool

		// Code 表示返回的状态码，其中 0 表示成功，其它均表示失败。
		Code() int

		// State 表示结果中的原始状态信息
		State() string

		// Args 表示Res所携带的参数。
		Args() []Any

		// With 方法，使用新的参数，生成一个新的 Res对象
		// 统常在需要返回动态结果的时候使用
		With(args ...Any) Res

		// Error 为兼容 error 对象的方法
		Error() string
	}

	// Vars 参数表
	Vars = map[string]Var
	// Var 参数
	Var struct {

		// nil 是否空参数，参考 Nil
		nil bool

		// Type 参数类型
		Type string

		// Required 是否必填，不可为空
		Required bool

		// Nullable 表示此参数可不存在
		// 为 true 时，解析参数的时候，如果为空，则不生成此参数，
		// 为 false 时，则生成默认值或nil
		Nullable bool

		// Name 参数名称
		Name string

		// Text 参数说明
		Text string

		// Default 默认值
		// 可以是常量，或是函数
		Default Any

		// Setting 参数配置，自定义配置
		// 一些特定的情况下，会有一些自定义的配置，
		Setting Map

		// Options 参数选项
		// 此参数不为空的时候，表示，值只允许是其中之一，相关于枚举选一
		// 注意，key为值，value为描述
		Options Map

		// Children 子参数表
		// 如此，就是支持无限下级的参数表
		Children Vars

		// Encode 编码方式
		// 具体参考 chef 内置的 codec 模块
		Encode string

		// Decode 解码方式
		// 具体参考 chef 内置的 codec 模块
		Decode string

		// Empty 自定义参数为空时的操作结果
		Empty Res

		// Error 自定义参数类型不对时的操作结果
		Error Res

		// Valid 自定义参数类型验证的方法
		Valid func(Any, Var) bool

		// Value 自定义参数通过验证后的值包装方法
		Value func(Any, Var) Any
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
