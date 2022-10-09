package framework

type Product interface {
	// 功能方法
	Use(str string)
	// 复制方法
	Clone() Product
}
