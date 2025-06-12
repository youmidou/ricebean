/*
*
added by yh @ 2023/4/27 17:22
注意:
*/
package sys_yml

import (
	"gopkg.in/yaml.v3"
)

func Marshal(config any) (out []byte, err error) {
	// 将配置对象转换成 YAML
	data, err := yaml.Marshal(&config)
	return data, err
	/*
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	*/
}
