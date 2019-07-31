package minimum_window_substring

func minWindow(s string, t string) string {

	var tmap = make(map[string]int, 10) //记录t中的字符以及个数
	var smap = make(map[string]int, 10) //记录window中命中t的字符个数

	var begin, end int //记录window的起始位置
	var count int      // 记录window中符合条件的字符个数

	var found = false
	var result = s

	//初始化tmap
	for i := 0; i < len(t); i++ {
		cur := t[i : i+1]
		if _, ok := tmap[cur]; ok {
			tmap[cur]++
		} else {
			tmap[cur] = 1
		}
	}

	for end < len(s) {

		//移动end，直到窗口符合条件
		for count < len(t) && end < len(s) {
			cur := s[end : end+1]
			val := 1
			if tv, ok := smap[cur]; ok {
				smap[cur]++
				val = tv + 1
			} else {
				smap[cur] = 1
			}
			//符合条件，count++
			if tmap[cur] >= val {
				count++
			}
			end++
		}

		//找到了窗口，开始移动begin
		for count == len(t) && begin < end {
			found = true
			if end-begin < len(result) {
				result = s[begin:end]
			}

			cur := s[begin : begin+1]
			val := smap[cur] - 1
			smap[cur]--
			if tval, ok := tmap[cur]; ok {
				//如果t中存在这个字符，且个数比窗口中的大，则count--
				if tval > val {
					count--
				}
			}
			begin++
		}

	}

	if !found {
		return ""
	}
	return result

}
