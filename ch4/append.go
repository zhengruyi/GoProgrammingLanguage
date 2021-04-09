package main

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	//如果切片还有空间,那么就可以直接将要添加进去的元素直接放进去
	if zlen <= cap(x){
		z = x[:zlen]
	}else{
		//扩容并且拷贝原始数组里面的元素到新的数组里面
		zcap := zlen
		if zcap <= 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return  z
}
