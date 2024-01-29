package flag

import "flag"

func ReadFlag() bool {
	var isTransimg bool

	// 将命令行参数与变量绑定
	flag.BoolVar(&isTransimg, "i", false, "Enable or disable the flag")

	// 解析命令行参数
	flag.Parse()
	return isTransimg
}
