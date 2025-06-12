package sys_utils

import (
	"phoenix-tudou/framework/sys_int"
	"strings"
)

/*
将 1,1000;2,1000;3,1000;4,1000;5,1000;
字符串转数组map
*/
func StringToMapString(mapStr string) map[string]string {
	list := make(map[string]string)
	slist := strings.Split(mapStr, ";")
	for _, s := range slist {
		if s == "" {
			continue
		}
		alist := strings.Split(s, ",")
		k := alist[0]
		v := alist[1]
		list[k] = v
		//k, err := sys_base.StringToInt32(alist[0])
	}
	return list

}

// 允许重复Id 1,1000;1,100;2,1000;2,1000;5,1000;
func StringToItemStr(mapStr string) []string {
	slist := strings.Split(mapStr, ";")
	return slist
}

// 1,1000;2,1000;3,1000;4,1000;5,1000;
func StringToMapInt32(mapStr string) (map[int32]int32, error) {
	list := make(map[int32]int32)
	slist := strings.Split(mapStr, ";")
	for _, s := range slist {
		if s == "" {
			continue
		}
		alist := strings.Split(s, ",")
		k, err1 := sys_int.StringToInt32(alist[0])
		if err1 != nil {
			return nil, err1
		}
		v, err2 := sys_int.StringToInt32(alist[1])
		if err2 != nil {
			return nil, err2
		}
		list[k] = v

	}
	return list, nil

}

//fmt.Sprintf("Hello %s, your age is %d", name, age)
/*
在Golang中，可以使用fmt包提供的函数进行格式化输出。常用的函数包括：

- fmt.Printf(format string, a ...interface{})：按照指定格式输出字符串，其中format是格式化字符串，a是可变参数列表。
- fmt.Sprintf(format string, a ...interface{})：按照指定格式返回字符串，其中format是格式化字符串，a是可变参数列表。
- fmt.Println(a ...interface{})：输出一行字符串，其中a是可变参数列表。
- fmt.Fprint(w io.Writer, a ...interface{})：将字符串输出到指定的io.Writer中，其中w是io.Writer接口类型的对象，a是可变参数列表。

在格式化字符串中，可以使用以下占位符：

- %v：按照默认格式输出变量的值。
- %d：输出整数。
- %f：输出浮点数。
- %s：输出字符串。
- %t：输出布尔值。
- %p：输出指针地址。
- %q：输出带有双引号的字符串。

例如，可以使用以下代码将变量a和b的值输出到控制台：

a := 123
b := "hello"
fmt.Printf("a=%d, b=%s\n", a, b)

输出结果为：a=123, b=hello
*/
