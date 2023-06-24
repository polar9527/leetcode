package __Singleton

type singletonLazyAndWrong struct{}

var insLWW *singletonLazyAndWrong

func GetInsLazyAndWrong() *singletonLazyAndWrong {
	if insLWW == nil {
		insLWW = &singletonLazyAndWrong{}
	}
	return insLWW
}
