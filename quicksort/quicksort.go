package quicksort

func Quicksort(arr []int)[]int{
	sort(arr,0,len(arr)-1);
	return arr
}
func sort(arr []int, low int,high int){
	if(low>=high){
		return 
	}
	pivot:=arr[high];
	left:=0;right:=high
	for(left<right){
		for(arr[left]<=pivot&&left<right){
			left++;
		}
		for(arr[right]>=pivot&&right>left){
			right--;
		}
		swap(arr,left,right)
	}
	swap(arr,left,high)
	sort(arr,low,left-1)
	sort(arr,left+1,high)
}	
func swap(arr []int,left int,right int){
	temp:=arr[left];
	arr[left]=arr[right]
	arr[right]=temp
}