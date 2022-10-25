package testing1

const(
	Lang1 = "Hello "
	Lang2 = "你好呀 "
	Lang3 = "Salom"
)

func Hello(lang , name string) string {
	if Lang1 == lang || Lang2 == lang || Lang3 == lang {
		return lang + " " + name
	} else {
		return ""
	}
}