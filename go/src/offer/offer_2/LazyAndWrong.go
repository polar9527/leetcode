package offer_2

type singletonLazyAndWrong struct{}

var insLWW *singletonLazyAndWrong

func GetInsLazyAndWrong() *singletonLazyAndWrong {
	if insLWW == nil {
		insLWW = &singletonLazyAndWrong{}
	}
	return insLWW
}
