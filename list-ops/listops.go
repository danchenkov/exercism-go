package listops

type Lister interface {
	length() int
}

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

type IntList []int
type StringList []string

func (list IntList) Length() int {
	return len(list)
}

func (list IntList) Foldl(fn binFunc, init int) int {
	l := len(list)
	switch l {
	case 0:
		return init
	case 1:
		return fn(init, list[0])
	default:
		return (list[1:l]).Foldl(fn, fn(list[0], init))
	}
}

func (list IntList) Foldr(fn binFunc, init int) int {
	l := len(list)
	switch l {
	case 0:
		return init
	case 1:
		return fn(list[0], init)
	default:
		return (list[0 : l-1]).Foldr(fn, fn(list[l-1], init))
	}
}

func (list IntList) Reverse() IntList {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func (list IntList) Filter(fn predFunc) IntList {
	filtered := []int{}
	for _, v := range list {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (list IntList) Map(fn unaryFunc) IntList {
	mapped := make([]int, len(list))
	for i, v := range list {
		mapped[i] = fn(v)
	}
	return mapped
}

func (list IntList) Append(appendList IntList) IntList {
	for _, v := range appendList {
		list = append(list, v)
	}
	return list
}

func (list IntList) Concat(concatList []IntList) IntList {
	for _, c := range concatList {
		for _, v := range c {
			list = append(list, v)
		}
	}
	return list
}
