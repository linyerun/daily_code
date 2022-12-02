package main

import "fmt"

func main() {
	var tomorrowDay, tomorrowMonth, tomorrowYear int
	var day, month, year int

	// 输入的年月日不一定是正确的
	fmt.Println("form: MM DD YYYY")
	_, err := fmt.Scanf("%d %d %d", &month, &day, &year)
	if err != nil {
		panic(err)
	}

	switch month {
	case 1, 3, 5, 7, 8, 10: //大月
		if day < 31 {
			tomorrowDay = day + 1
		} else {
			// 自己加的
			if day > 31 {
				fmt.Println("无效输入")
				return
			}
			tomorrowDay = 1
			tomorrowMonth = month + 1
		}
	case 4, 6, 9, 11: // 小月
		if day < 30 {
			tomorrowDay = day + 1
		} else {
			// 自己加的
			if day > 30 {
				fmt.Println("无效输入")
				return
			}
			tomorrowDay = day + 1
			tomorrowMonth = month + 1
		}
	case 12: // 特殊月1
		if day < 31 {
			tomorrowDay = day + 1
		} else {
			// 自己加的
			if day > 31 {
				fmt.Println("无效输入")
				return
			}
			tomorrowDay = 1
			tomorrowMonth = 1
		}
		if year == 2012 {
			fmt.Println("2012 is over")
			return
		} else {
			tomorrowYear = year + 1
		}
	case 2:
		if day < 28 {
			tomorrowDay = day + 1
		} else {
			if day == 28 {
				if isLeapYear(year) {
					tomorrowDay = 29
				} else {
					tomorrowDay = 1
					tomorrowMonth = 3
				}
			} else {
				if isLeapYear(year) {
					if day == 29 {
						tomorrowDay = 1
						tomorrowMonth = 3
					} else {
						fmt.Println("无效输入")
						return
					}
				} else {
					fmt.Println("无效输入")
					return
				}
			}
		}
	}
	fmt.Println("YYYY-MM-DD: ", tomorrowYear, tomorrowMonth, tomorrowDay)
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
		} else {
			return true
		}
	}
	return false
	//return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
