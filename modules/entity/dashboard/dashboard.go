package dashboard

type FavouriteProducts struct {
	Name        string
	TotalOrders int64
}
type MonthlyRevenue struct {
	Month   string
	Revenue float64
}

type WeeklyRevenue struct {
	Day     string
	Revenue float64
}
