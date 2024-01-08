## 1. 标准库







## 2. github

### 1. [creasty](https://github.com/creasty)/**[defaults](https://github.com/creasty/defaults)**

链接：https://github.com/creasty/defaults

使用默认值初始化结构体

- Supports almost all kind of types
  - Scalar types
    - `int/8/16/32/64`, `uint/8/16/32/64`, `float32/64`
    - `uintptr`, `bool`, `string`
  - Complex types
    - `map`, `slice`, `struct`
  - Nested types
    - `map[K1]map[K2]Struct`, `[]map[K1]Struct[]`
  - Aliased types
    - `time.Duration`
    - e.g., `type Enum string`
  - Pointer types
    - e.g., `*SampleStruct`, `*int`
- Recursively initializes fields in a struct
- Dynamically sets default values by [`defaults.Setter`](https://github.com/creasty/defaults/blob/master/setter.go) interface
- Preserves non-initial values from being reset with a default value

```go
type Gender string

type Sample struct {
	Name   string `default:"John Smith"`
	Age    int    `default:"27"`
	Gender Gender `default:"m"`

	Slice       []string       `default:"[]"`
	SliceByJSON []int          `default:"[1, 2, 3]"` // Supports JSON

	Map                 map[string]int `default:"{}"`
	MapByJSON           map[string]int `default:"{\"foo\": 123}"`
	MapOfStruct         map[string]OtherStruct
	MapOfPtrStruct      map[string]*OtherStruct
	MapOfStructWithTag  map[string]OtherStruct `default:"{\"Key1\": {\"Foo\":123}}"`
    
	Struct    OtherStruct  `default:"{}"`
	StructPtr *OtherStruct `default:"{\"Foo\": 123}"`

	NoTag  OtherStruct               // Recurses into a nested struct by default
	OptOut OtherStruct `default:"-"` // Opt-out
}

type OtherStruct struct {
	Hello  string `default:"world"` // 嵌套结构体中的标记也可以工作
	Foo    int    `default:"-"`
	Random int    `default:"-"`
}

// SetDefaults implements defaults.Setter interface
func (s *OtherStruct) SetDefaults() {
	if defaults.CanUpdate(s.Random) { // Check if it's a zero value (recommended)
		s.Random = rand.Int() // Set a dynamic value
	}
}

obj := &Sample{}
if err := defaults.Set(obj); err != nil {
	panic(err)
}
```

### 2. [duke-git](https://github.com/duke-git)/**[lancet](https://github.com/duke-git/lancet)**

链接：https://github.com/duke-git/lancet

lancet（柳叶刀）是一个全面、高效、可复用的go语言工具函数库。

1. algorithm 包实现一些基本查找和排序算法

2. concurrency 包含一些支持并发编程的功能。例如：goroutine, channel, async 等

3. condition 包含一些用于条件判断的函数。
4. convertor 转换器包支持一些常见的数据类型转换。
5. cryptor 加密包支持数据加密和解密，获取 md5，hash 值。支持 base64, md5, hmac, aes, des, rsa。
6. datetime 日期时间处理包，格式化日期，比较日期。
7. datastructure 包含一些普通的数据结构实现。例如：list, linklist, stack, queue, set, tree, graph.
8. fileutil 包含文件基本操作。
9. formatter 格式化器包含一些数据格式化处理方法。
10. function 函数包控制函数执行流程，包含部分函数式编程。
11. maputil 包括一些操作 map 的函数.



