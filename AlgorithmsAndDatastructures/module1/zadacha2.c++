#include <iostream>
#include <math.h>
// Дан отсортированный массив целых чисел A[0..n-1] и массив целых чисел B[0..m-1].
// Для каждого элемента массива B[i] найдите минимальный индекс k минимального элемента массива A, равного или превосходящего B[i]: A[k] >= B[i].
// Если такого элемента нет, выведите n. n, m ≤ 10000.
//Требования:  Время работы поиска k для каждого элемента B[i]: O(log(k)).
// Внимание! В этой задаче для каждого B[i] сначала нужно определить диапазон для бинарного поиска размером порядка k с помощью экспоненциального поиска,
// а потом уже в нем делать бинарный поиск. Экспонентальный поиск основан на степенях двойки (от нуля до степеней двойки), в нем делаем бинарный поиск.
int binsearch(const int arr[], int start, int end, int znach) {
    while(start <= end) {
        int middle = start + (end - start) / 2;
        if (arr[middle] == znach) {
            return middle;
        }
        if (arr[middle] < znach) {
            start = middle + 1;
        } else {
            end = middle - 1;
        }
    }
    return end+1;
}

int expsearch(const int arr[], int lenarr, int znach){
    int exp = 2;
    int i=0;
    int power = (int)pow(exp, i);
    while (power < lenarr){
        if(znach <= arr[power]){
            return binsearch(arr, 0, power, znach);
        }
        power*=2;
    }
    if(znach <= arr[lenarr - 1]){
        return binsearch(arr, 0, lenarr - 1, znach);
    }
    return lenarr;
}

int main() {
    int n;
    int m;
    std::cin >> n >> m;
    int *arr1 = new int[n];
    int *arr2 = new int[m];
    for (int i = 0; i < n; ++i){
        std::cin>> arr1[i];
    }
    for (int j = 0; j < m; ++j){
        std::cin>> arr2[j];
    }
    for (int j = 0; j < m; ++j){
        std::cout << expsearch(arr1, n, arr2[j]) << " ";
    }
    delete[] arr1;
    delete[] arr2;
    return 0;
}
