package problem0675

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	forest [][]int
	ans    int
}{

	{
		[][]int{
			{69438, 55243, 0, 43779, 5241, 93591, 73380},
			{847, 49990, 53242, 21837, 89404, 63929, 48214},
			{90332, 49751, 0, 3088, 16374, 70121, 25385},
			{14694, 4338, 87873, 86281, 5204, 84169, 5024},
			{31711, 47313, 1885, 28332, 11646, 42583, 31460},
			{59845, 94855, 29286, 53221, 9803, 41305, 60749},
			{95077, 50343, 27947, 92852, 0, 0, 19731},
			{86158, 63553, 56822, 90251, 0, 23826, 17478},
			{60387, 23279, 78048, 78835, 5310, 99720, 0},
			{74799, 48845, 60658, 29773, 96129, 90443, 14391},
			{65448, 63358, 78089, 93914, 7931, 68804, 72633},
			{93431, 90868, 55280, 30860, 59354, 62083, 47669},
			{81064, 93220, 22386, 22341, 95485, 20696, 13436},
			{50083, 0, 89399, 43882, 0, 13593, 27847},
			{0, 12256, 33652, 69301, 73395, 93440, 0},
			{42818, 87197, 81249, 33936, 7027, 5744, 64710},
			{35843, 0, 99746, 52442, 17494, 49407, 63016},
			{86042, 44524, 0, 0, 26787, 97651, 28572},
			{54183, 83466, 96754, 89861, 84143, 13413, 72921},
			{89405, 52305, 39907, 27366, 14603, 0, 14104},
			{70909, 61104, 70236, 30365, 0, 30944, 98378},
			{20124, 87188, 6515, 98319, 78146, 99325, 88919},
			{89669, 0, 64218, 85795, 2449, 48939, 12869},
			{93539, 28909, 90973, 77642, 0, 72170, 98359},
			{88628, 16422, 80512, 0, 38651, 50854, 55768},
			{13639, 2889, 74835, 80416, 26051, 78859, 25721},
			{90182, 23154, 16586, 0, 27459, 3272, 84893},
			{2480, 33654, 87321, 93272, 93079, 0, 38394},
			{34676, 72427, 95024, 12240, 72012, 0, 57763},
			{97957, 56, 83817, 45472, 0, 24087, 90245},
			{32056, 0, 92049, 21380, 4980, 38458, 3490},
			{21509, 76628, 0, 90430, 10113, 76264, 45840},
			{97192, 58807, 74165, 65921, 45726, 47265, 56084},
			{16276, 27751, 37985, 47944, 54895, 80706, 2372},
			{28438, 53073, 0, 67255, 38416, 63354, 69262},
			{23926, 75497, 91347, 58436, 73946, 39565, 10841},
			{34372, 69647, 44093, 62680, 32424, 69858, 68719},
			{24425, 4014, 94871, 1031, 99852, 88692, 31503},
			{24475, 12295, 33326, 37771, 37883, 74568, 25163},
			{0, 18411, 88185, 60924, 29028, 69789, 0},
			{34697, 75631, 7636, 16190, 60178, 39082, 7052},
			{24876, 9570, 53630, 98605, 22331, 79320, 88317},
			{27204, 89103, 15221, 91346, 35428, 94251, 62745},
			{26636, 28759, 12998, 58412, 38113, 14678, 0},
			{80871, 79706, 45325, 3861, 12504, 0, 4872},
			{79662, 15626, 995, 80546, 64775, 0, 68820},
			{25160, 82123, 81706, 21494, 92958, 33594, 5243},
		},
		5637,
	},

	{
		[][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}},
		-1,
	},

	{
		[][]int{
			{1, 2, 3},
			{0, 0, 4},
			{7, 6, 5},
		},
		6,
	},

	{
		[][]int{{18, 17, 16, 15, 14, 13}, {7, 8, 9, 10, 11, 12}, {6, 5, 4, 3, 2, 1}},
		22,
	},

	{
		[][]int{{9, 8, 7}, {4, 5, 6}, {3, 2, 1}},
		10,
	},

	{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		12,
	},

	{
		[][]int{{1, 2, 3}, {0, 1, 0}, {7, 1, 1}},
		6,
	},

	{
		[][]int{{1, 2, 3}, {0, 1, 0}, {7, 6, 5}},
		8,
	},

	{
		[][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}},
		6,
	},

	// 可以有多个 testcase
}

func Test_cutOffTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, cutOffTree(tc.forest), "输入:%v", tc)
	}
}

func Benchmark_cutOffTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			cutOffTree(tc.forest)
		}
	}
}

// 为了 100% 覆盖率
func Test_PQ_push(t *testing.T) {
	ast := assert.New(t)
	heights := []int{9, 8, 7, 4, 5, 6, 3, 2, 1}
	pq := make(PQ, 0, 9)
	for i := 0; i < len(heights); i++ {
		heap.Push(&pq, &tree{height: heights[i]})
	}

	actual := heap.Pop(&pq).(*tree).height
	expected := 1
	ast.Equal(expected, actual)

}
