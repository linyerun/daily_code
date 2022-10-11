//一个变量两部分组成：type、value
// type：static type(string、int....)、concrete type

package main

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	println("读书")
}

func (b *Book) WriteBook() {
	println("写书")
}

func main() {
	//pair<type: Book, value: Book地址>
	b := new(Book)

	var reader Reader
	//pair<type: Book, value: Book地址>
	reader = b
	reader.ReadBook()

	var writer Writer
	//pair<type: Book, value: Book地址>
	writer, ok := reader.(Writer) //ok可以不写，只需要writer := reader.(Writer)

	//因为reader和writer的pair一样，所以可以类型断言成功 (我觉得不对)
	//应该是因为reader的type实现了writer接口的方法
	if ok {
		writer.WriteBook()
	}
}
