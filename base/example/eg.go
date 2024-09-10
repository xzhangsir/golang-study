package example

/* // 出现次数最多的N个数
type NumCount struct {
	num   int
	count int
}
type ByCount []NumCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].count > a[j].count }

func solution(nums []int, k int) string {
	// Please write your code here
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	var arr []NumCount
	for num, count := range m {
		arr = append(arr, NumCount{num: num, count: count})
	}
	sort.Sort(ByCount(arr))
	str := []string{}
	for _, v := range arr[:k] {
		str = append(str, strconv.Itoa(v.num))
	}
	return strings.Join(str, ",")

}

func EgFunc() {
	fmt.Println(solution([]int{1, 1, 1, 2, 2, 3}, 2) == "1,2")
	fmt.Println(solution([]int{1}, 1) == "1")
}
*/
func EgFunc() {}
