package model

const (
	RedShuai   = 1
	RedShi     = 2
	RedXiang   = 3
	RedMa      = 4
	RedJu      = 5
	RedPao     = 6
	RedBing    = 7
	BlackShuai = 8
	BlackShi   = 9
	BlackXiang = 10
	BlackMa    = 11
	BlackJu    = 12
	BlackPao   = 13
	BlackBing  = 14
)

var Board = [10][9]int{
	{5, 4, 3, 2, 1, 2, 3, 4, 5},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 6, 0, 0, 0, 0, 0, 6, 0},
	{7, 0, 7, 0, 7, 0, 7, 0, 7},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{14, 0, 14, 0, 14, 0, 14, 0, 14},
	{0, 13, 0, 0, 0, 0, 0, 13, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{12, 11, 10, 9, 8, 9, 10, 11, 12},
}

func Move(isRed bool, x1, y1, x2, y2 int) string {
	if x2 < 0 || x2 > 8 || y2 < 0 || y2 > 9 || (x1 == x2 && y1 == y2) {
		return "移动错误"
	}
	if Board[y1][x1] == 0 {
		return "未选中棋子"
	}
	if Board[y1][x1]/8 == 1 && isRed {
		return "未选中红方棋子"
	}
	if Board[y1][x1]/8 == 0 && !isRed {
		return "未选中黑方棋子"
	}
	//目的地位置为己方棋子
	if Board[y2][x2]/8 == 0 && Board[y2][x2] != 0 {
		return "移动错误"
	}
	switch Board[y1][x1] % 7 {
	case 1: //王
		if x2-x1+y2-y1 != 1 && x2-x1+y2-y1 != -1 {
			return "移动错误"
		}
		if x2 < 3 || x2 > 5 {
			return "移动错误"
		}
		if y2 < 3 || y2 > 6 {
			Board[y2][x2] = 1
			Board[y1][x1] = 0
			break
		}
	case 2: //士
		if x1-x2 != 1 && x1-x2 != -1 {
			return "移动错误"
		}
		if y1-y2 != 1 && y1-y2 != -1 {
			return "移动错误"
		}
		Board[y2][x2] = 2
		Board[y1][x1] = 0
	case 3: //相
		if x1-x2 != 2 && x1-x2 != -2 {
			return "移动错误"
		}
		if y1-y2 != 2 && y1-y2 != -2 {
			return "移动错误"
		}
		//撇象脚
		if Board[y1+(y1-y2)/2][x1+(x1-x2)/2] != 0 {
			return "移动错误"
		}
		Board[y2][x2] = RedXiang
		Board[y1][x1] = 0
	case 4: //马
		if x1-x2 == 1 || x1-x2 == -1 {
			if y1-y2 == 2 || y1-y2 == -2 {
				//不撇马脚
				if Board[y1+(y2-y1)/2][x1] == 0 {
					Board[y2][x2] = 4
					Board[y1][x1] = 0
					break
				}
			}
		}
		if y1-y2 == 1 || y1-y2 == -1 {
			if x1-x2 == 2 || x1-x2 == -2 {
				//不撇马脚
				if Board[y1][x1+(x2-x1)/2] == 0 {
					Board[y2][x2] = 4
					Board[y1][x1] = 0
					break
				}
			}
		}
	case 5: //车
		if x1 == x2 {
			for i := y1 + 1; i < y2; i++ {
				if Board[i][x1] != 0 {
					return "移动错误"
				}
			}
		}
		if y1 == y2 {
			for i := x1 + 1; i < x2; i++ {
				if Board[y1][i] != 0 {
					return "移动错误"
				}
			}
		}
		Board[y2][x2] = 5
		Board[y1][x1] = 0
		break
	case 6: //炮
		if Board[y2][x2] == 0 {
			return "移动错误"
		}
		fla := 0 //炮架
		if x1 == x2 {
			for i := y1 + 1; i < y2; i++ {
				if Board[i][x1] != 0 {
					fla++
				}
			}
			if fla == 1 {
				Board[y2][x2] = 6
				Board[y1][x1] = 0
				break
			}
		}
		if y1 == y2 {
			for i := x1 + 1; i < x2; i++ {
				if Board[y1][i] != 0 {
					fla++
				}
			}
			if fla == 1 {
				Board[y2][x2] = 6
				Board[y1][x1] = 0
				break
			}
		}
	case 0: //兵
		if y2-y1 < 0 {
			return "移动错误"
		}
		//过河兵
		if y1 > 4 {
			if y2-y1 == 1 && x1 == x2 {
				Board[y2][x2] = 7
				Board[y1][x1] = 0
				break
			}
			if y2 == y1 && x1-x2 == -1 || x1-x2 == 1 {
				Board[y2][x2] = 7
				Board[y1][x1] = 0
				break
			}
		} else { //未过河
			if x1 == x2 && y2-y1 == 1 {
				Board[y2][x2] = 7
				Board[y1][x1] = 0
				break
			}
		}
	}

	flag := false
	//判断王是否还在
	for i := 3; i < 6; i++ {
		for j := 0; j < 3; j++ {
			if Board[j][i] == 1 {
				flag = true
			}
		}
	}
	if flag == false {
		return "黑方胜"
	}
	for i := 3; i < 6; i++ {
		for j := 7; j < 10; j++ {
			if Board[j][i] == 1 {
				flag = true
			}
		}
	}
	if flag == false {
		return "红方胜"
	}
	return "成功移动"
}
