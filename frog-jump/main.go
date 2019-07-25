package main

func canCross(stones []int) bool {

	if len(stones) < 2 {
		return false
	}

	if stones[0] != 0 && stones[1] != 1 {
		return false
	}

	type point struct {
		a int //当前值
		b int //last jump
	}

	var numMap = make(map[int]bool, 100)
	var pointMap = make(map[point]bool, 100)

	var NQ = make([]*point, 0, 100)

	//初始化numMap
	for _, num := range stones {
		numMap[num] = true
	}

	lastStone := stones[len(stones)-1]

	first := &point{0, 0}

	NQ = append(NQ, first)

	for len(NQ) > 0 {
		cur := NQ[0]
		NQ = NQ[1:]

		if cur.a == lastStone {
			return true
		}

		//check jump - 1
		if cur.b-1 > 0 {
			checkNum := cur.a + cur.b - 1
			if numMap[checkNum] {
				pointKey := &point{
					a: checkNum,
					b: cur.b - 1,
				}
				if _, ok := pointMap[*pointKey]; !ok {
					//不存在则插入队列
					NQ = append(NQ, pointKey)
					pointMap[*pointKey] = true
				}
			}
		}

		//check jump
		if cur.b > 0 {
			checkNum := cur.a + cur.b
			if numMap[checkNum] {
				pointKey := &point{
					a: checkNum,
					b: cur.b,
				}
				if _, ok := pointMap[*pointKey]; !ok {
					//不存在则插入队列
					NQ = append(NQ, pointKey)
					pointMap[*pointKey] = true
				}
			}
		}

		//check jump+1
		checkNum := cur.a + cur.b + 1
		if numMap[checkNum] {
			pointKey := &point{
				a: checkNum,
				b: cur.b + 1,
			}
			if _, ok := pointMap[*pointKey]; !ok {
				//不存在则插入队列
				NQ = append(NQ, pointKey)
				pointMap[*pointKey] = true
			}
		}

	}

	return false

}
