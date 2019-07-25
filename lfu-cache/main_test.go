package lfu_cache

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
)

/**
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.get(3);       // returns 3.
cache.put(4, 4);    // evicts key 1.
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4

["LFUCache","put"
"get"]
[[0],[0,0],[0]]

["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
[null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]

10 3 6

*/

func TestCases(t *testing.T) {

	tests := []struct {
		op     string
		input  string
		expect string
	}{
		{`["LFUCache","put","put","get","put","get","get","put","get","get","get"]`,
			`[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]`,
			`[null,null,null,1,null,-1,3,null,-1,3,4]`},
		{
			`["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]`,
			`[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]`,
			`[null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]`,
		},
	}

	for _, tt := range tests {

		var oparr []string
		var input [][]int
		var expect []int

		json.Unmarshal([]byte(tt.op), &oparr)
		json.Unmarshal([]byte(tt.input), &input)
		json.Unmarshal([]byte(tt.expect), &expect)

		var cache = Constructor(input[0][0])

		for i := 1; i < len(oparr); i++ {

			if oparr[i] == "put" {
				log.Println("put", input[i][0], input[i][1])
				cache.Put(input[i][0], input[i][1])
			} else if oparr[i] == "get" {
				log.Println("get", input[i][0])
				val := cache.Get(input[i][0])
				if val != expect[i] {
					t.Fatalf("index:%v, key:%v , expected:%v, got:%v", i, input[i][0], expect[i], val)
				}
			}
			cap := 0
			for j := cache.list.Front(); j != nil; j = j.Next() {
				cap += 4
				fmt.Println(i, strings.Repeat(" ", cap), "↓:")
				fmt.Print(strings.Repeat(" ", cap*2))
				for k := j.Value.(*ParentItem).sublist.Front(); k != nil; k = k.Next() {
					vv := k.Value.(*Item)
					fmt.Print(vv.key, ":", vv.val, "->")
				}
				fmt.Println()
			}

		}

	}

}

func TestLFUCache(t *testing.T) {

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	val := cache.Get(1) // returns 1
	if val != 1 {
		t.Errorf("expect %d, got %d", 1, val)
	}
	cache.Put(3, 3)    // evicts key 2
	val = cache.Get(2) // returns -1 (not found)
	if val != -1 {
		t.Errorf("expect %d, got %d", -1, val)
	}
	val = cache.Get(3)
	if val != 3 {
		t.Errorf("expect %d, got %d", 3, val)
	}
	cache.Put(4, 4)    // evicts key 1
	val = cache.Get(1) // returns -1 (not found)
	if val != -1 {
		t.Errorf("expect %d, got %d", -1, val)
	}
	val = cache.Get(3) // returns 3
	if val != 3 {
		t.Errorf("expect %d, got %d", 3, val)
	}
	val = cache.Get(4) // returns 4
	if val != 4 {
		t.Errorf("expect %d, got %d", 4, val)
	}

}
