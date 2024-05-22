package money

import "github.com/gin-gonic/gin"

// 计算分期付款时，需要达到支出利息收入需要的年化收益利率
// principal 分期本金
// interest 扣除的利息，也是预期获取的收益
// n 分期期数(月)
// annualizedInterestRate 年化收益率
func CountInterest(principal, interest float64, n int) (annualizedInterestRate float64) {
	all := principal
	for m := 1; m < n; m++ {
		all += principal - (principal+interest)/float64(n)*float64(m)
	}
	return interest / all
}

func Interest(c *gin.Context) {
	c.String(200, "")
}
