package group_anagrams

import "strings"

func hash(str string) string {
	var char [26]int

	for _, v := range str {
		char[v-'a']++
	}
	var res string
	for i := 0; i < len(char); i++ {
		tmp := strings.Repeat(string('a'+i), char[i])
		res += tmp
	}
	return res
}

func groupAnagrams(strs []string) [][]string {

	var hashmap = make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		key := hash(strs[i])
		val, ok := hashmap[key]
		if ok {
			val = append(val, strs[i])
		} else {
			val = []string{strs[i]}
		}
		hashmap[key] = val
	}

	var result [][]string

	for _, v := range hashmap {
		result = append(result, v)
	}

	return result

}

func groupAnagrams2(strs []string) [][]string {
	anagramGroups := make(map[[26]byte]int, len(strs))
	result := make([][]string, 0)

	for _, str := range strs {
		list := [26]byte{}
		for _, v := range str {
			list[v-'a']++
		}
		if idx, ok := anagramGroups[list]; ok {
			result[idx] = append(result[idx], str)
		} else {
			result = append(result, []string{str})
			anagramGroups[list] = len(result) - 1
		}
	}
	return result

}
