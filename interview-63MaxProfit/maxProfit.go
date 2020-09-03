package interview_63MaxProfit

/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？


示例 1:

输入: [7,3,5,1,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 

限制：

0 <= 数组长度 <= 10^5
*/

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	in, out, otherIn := 0, 0, -1
	for i, v := range prices {
		if i == 0 {
			in = v
		} else {
			//进价要比之前的进价低的时候记录一下
			if v < in {
				//如果没有记录过更低进价
				if otherIn == -1 {
					otherIn = v
				} else {
					//否则判断一下是不是比之前记录过的还要低
					if v < otherIn {
						otherIn = v
					}
				}
			}
			if v > out {
				out = v
			}
			//如果进价跟新后赚的多了
			if (v-otherIn+1 > out-in) && otherIn >= 0 {
				out = v
				in = otherIn
			}
		}
	}
	return out - in
}
