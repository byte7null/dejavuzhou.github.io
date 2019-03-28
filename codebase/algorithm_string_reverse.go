package codebase

func LeftShiftOne(b []byte)[]byte  {
	t := b[0]
	n := len(b)
	for i:=1; i <n;i++{
		b[i-1]=b[i]
	}
	b[n-1]=t
	return b
}
func LeftRotateStringV0(b []byte, m int) []byte {
	for i:=0; i< m;i++{
		b = LeftShiftOne(b)
	}
	return b
}

func ReverseString(b []byte,from,to int) []byte {
	var fromI,toI = from,to
	for fromI < toI  {
		t := b[fromI]
		b[fromI] = b[toI]
		b[toI] = t
		toI--
		fromI++
	}
	return b
}

func LeftRotateString(s string ,m int)string {
	b := []byte(s)
	n := len(b)
	b = ReverseString(b, 0, m - 1) //反转[0..m - 1]，套用到上面举的例子中，就是X->X^T，即 abc->cba
	b = ReverseString(b, m, n - 1)//反转[m..n - 1]，例如Y->Y^T，即 def->fed
	b = ReverseString(b, 0, n - 1) //反转[0..n - 1]，即如整个反转，(X^TY^T)^T=YX，即 cbafed->defabc。
	return string(b)
}
