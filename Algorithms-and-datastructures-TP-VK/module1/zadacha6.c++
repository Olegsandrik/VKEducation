
if (pivot_pos>k){
            int tmp = pivot_pos-1;
            pivot_pos = partition(arr, l, tmp, cmp);
        } else{
            int tmp = pivot_pos+1;
            pivot_pos = partition(arr, tmp, r, cmp);
        }

template <typename T, typename Comparator = std::less<T>>
T search_statistic(T *arr, int k, int l, int r, Comparator cmp = Comparator())
{
    int pivot_pos = partition(arr, l, r, cmp);
    while(1){
        if (pivot_pos == k)
        {
            return arr[pivot_pos];
        }
        if (pivot_pos>k){
            r+=1;
            pivot_pos = partition(arr, l, r, cmp);
        } else{
            l-=1;
            pivot_pos = partition(arr, l, r, cmp);
        }
    }
}
