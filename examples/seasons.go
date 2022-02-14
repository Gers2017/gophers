package examples

/*
   Winter: 1, 2, and 12
   Spring: 3, 4, and 5
   Summer: 6, 7, and 8
   Autumn: 9, 10, and 11
*/
func Season(month int) string  {
	switch month {
		case 1, 2, 12: return "Winter"
		case 3, 4, 5: return "Spring"
		case 6, 7, 8: return "Summer"
		case 9, 10, 11: return "Autumn"
		default: return "Season unknown"		
	}
}