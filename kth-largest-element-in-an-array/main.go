package kth_largest_element_in_an_array

import "log"

func quickSort(nums []int, lo int, hi int) {
	if lo < hi {
		log.Printf("%+v", nums)
		pos := partition2(nums, lo, hi)
		log.Printf("%+v,lo:%v,hi:%v,pos:%v", nums, lo, hi, pos)
		quickSort(nums, lo, pos-1)
		quickSort(nums, pos+1, hi)
	}
}

/***

// Do partition in arr[begin, end), with the first element as the pivot.
int partition(vector<int>&arr, int begin, int end){
    int pivot = arr[begin];
    // Last position where puts the no_larger element.
    int pos = begin;
    for(int i=begin+1; i!=end; i++){
        if(arr[i] <= pivot){
            pos++;
            if(i!=pos){
                swap(arr[pos], arr[i]);
            }
        }
    }
    swap(arr[begin], arr[pos]);
    return pos;
}

*/
// 3 6 1 4 5 3 7 3
func partition1(nums []int, lo int, hi int) int {
	pivot := nums[lo]
	i := lo
	for j := lo + 1; j <= hi; j++ {
		if nums[j] <= pivot {
			i++
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	nums[i], nums[lo] = nums[lo], nums[i]
	return i
}

/**
int partition(vector<int>&arr, int begin, int end)
{
    int pivot = arr[begin];
    while(begin < end)
    {
        while(begin < end && arr[--end] >= pivot);
        arr[begin] = arr[end];
        while(begin < end && arr[++begin] <= pivot);
        arr[end] = arr[begin];
    }
    arr[begin] = pivot;
    return begin;
}

*/

func partition2(nums []int, lo int, hi int) int {

	pivot := nums[lo]
	for lo < hi {
		for lo < hi && nums[hi] >= pivot {
			hi--
		}
		nums[lo] = nums[hi]
		lo++
		for lo < hi && nums[lo] <= pivot {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = pivot
	return lo
}

func findKthLargest(nums []int, k int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	pos := partition2(nums, 0, len(nums)-1)

	kLargePos := len(nums) - k

	if pos == kLargePos {
		log.Printf("%+v,%v", nums, k)
		return nums[pos]
	}

	if pos > kLargePos {
		//find in left part
		return findKthLargest(nums[:pos], k-pos-1)
	}
	//find in right part
	return findKthLargest(nums[pos+1:], k)

}
