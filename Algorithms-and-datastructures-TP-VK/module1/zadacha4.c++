//Имеется лог-файл, в котором хранятся пары для N пользователей (Идентификатор пользователя, посещаемость сайта).
//Напишите программу, которая выбирает K пользователей, которые чаще других заходили на сайт, и выводит их
// в порядке возрастания посещаемости. Количество заходов и идентификаторы пользователей не повторяются.
#include "iostream"
template <typename T, typename Comparator = std::less<T>>
class MyHeap{
public:
    MyHeap(){
        capacity = 1;
        arr = new T[capacity];
        count = 0;
    }
    ~MyHeap(){
        delete[] arr;
    }
    MyHeap& operator=(const MyHeap& copy){
        if (this != &copy){
            if (arr != nullptr) {
                delete[] arr;
            }
            capacity = copy.capacity;
            arr = new T[capacity];
            for (int i=0; i<capacity; ++i){
                arr[i]=copy.arr[i];
            }
            count = copy.count;
        }
        return *this;
    }
    // Добавление элемента в кучу
    void Insert(T element)
    {
        if (count==capacity) {
            capacity *= 2;
            T *newArr = new T[capacity];
            for (int i =0; i<=count; i++) {
                newArr[i] = arr[i];
            }
            delete[] arr;
            arr = newArr;
        }
        arr[count] = element;
        siftUp(count);
        count++;
    }
    // Извлечение максимального элемента.
    T ExtractMax(){
        if (count != 0) {
            T result = arr[0];
            arr[0] = arr[count - 1];
            count--;
            siftDown(0);
            return result;
        }
    }
    int size(){
        return count;
    }
private:
    // Проталкивание элемента наверх.
    void siftUp( int index )
    {
        while( index > 0 ) {
            int parent = ( index - 1 ) / 2;
            if(!cmp(arr[parent] , arr[index] )) {
                return;
            }
            std::swap( arr[index], arr[parent] );
            index = parent;
        }
    }
    void siftDown( int i )
    {
        int left = 2 * i + 1;
        int right = 2 * i + 2;
        // Ищем большего сына, если такой есть.
        int largest = i;
        if( left < count && cmp(arr[i], arr[left]) )
            largest = left;
        if( right < count &&  cmp(arr[largest],arr[right]) )
            largest = right;
        // Если больший сын есть, то проталкиваем корень в него.
        if( largest != i ) {
            std::swap( arr[i], arr[largest] );
            siftDown( largest );
        }
    }
    // Построение кучи.
    void buildHeap()
    {
        for(int i = count / 2 - 1; i >= 0; --i) {
            siftDown(i);
        }
    }
    Comparator cmp;
    T* arr;
    int count;
    int capacity;
};
struct myuser{
    myuser(): id(0), posesh(0){}
    myuser(int x, int y): id(x), posesh(y){}
    int id;
    int posesh;

};

struct UserComparator {
    bool operator() (const myuser &l, const myuser &r) {
        return l.posesh < r.posesh;
    }
};
void result(int N, int K, int *ans){
    MyHeap<myuser, UserComparator> myheap;
    // вводится индефикатор и посещаемость, то есть нужно сортить по посещаемости
    for(int i = 0; i < N; ++i) {
        int id, posesh;
        std::cin >> id >> posesh;
        myheap.Insert(myuser(id, posesh));
        if (myheap.size()>K){ // ограничеваем размер кучи
            MyHeap<myuser, UserComparator> myheapsmall;
            for(int j=0; j<K; ++j){
                myheapsmall.Insert(myheap.ExtractMax());
            }
            myheap = myheapsmall;
        }
    }
    for(int i = 0; i < K; ++i){
        myuser max_user = myheap.ExtractMax();
        ans[i]=max_user.id;
    }
}
int main(){
    int N, K;
    std::cin >> N >> K;
    int *ans = new int[K];
    result(N, K, ans);
    for (int j=K-1; j>=0; j--){
        std::cout << ans[j] << " ";
    }
    delete[] ans;
}
