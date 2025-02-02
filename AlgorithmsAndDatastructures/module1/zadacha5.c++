//5_1. Реклама.
//В супермаркете решили оптимизировать показ рекламы. Известно расписание прихода и ухода покупателей (два целых числа).
// Каждому покупателю необходимо показать минимум 2 рекламы.  Рекламу можно транслировать только в целочисленные моменты времени.
// Покупатель может видеть рекламу от момента прихода до момента ухода из магазина.
//В каждый момент времени может показываться только одна реклама. Считается, что реклама показывается мгновенно.
// Если реклама показывается в момент ухода или прихода, то считается, что посетитель успел её посмотреть.
// Требуется определить минимальное число показов рекламы.
#include "iostream"
template <typename T>
void MemCopy(T* a, T* b, int len) {
    while (len > 0){
        *a = *b;
        a++;
        b++;
        len--;
    }
}
template <typename T, typename Compare=std::less<T>>
void Merge(T* a, int aLen, T* b, int bLen, T* ans, Compare cmp = Compare()) {
    while (aLen > 0 && bLen > 0){
        if (cmp(*a, *b)){
            *ans = *a;
            a++;
            aLen--;
        } else {
            *ans = *b;
            b++;
            bLen--;
        }
        ans++;
    }
    while (aLen > 0){
        *ans = *a;
        a++;
        aLen--;
        ans++;
    }
    while (bLen > 0){
        *ans = *b;
        ans++;
        b++;
        bLen--;
    }
}

template <typename T, typename Compare=std::less<T>>
void MergeSort(T* a, int aLen, Compare cmp = Compare()) {
    if (aLen <= 1 ) {
        return;
    }
    int firstLen = aLen / 2;
    int secondLen = aLen - firstLen;
    MergeSort(a, firstLen, cmp);
    MergeSort(a + firstLen, secondLen, cmp);
    T* c = new T[aLen];
    Merge(a, firstLen, a + firstLen, secondLen, c, cmp);
    MemCopy(a, c, aLen);
    delete[] c;
}

struct myuser{
    myuser(): in(0), out(0){}
    myuser(int x, int y): in(x), out(y){}
    int in;
    int out;
};

struct UserComparator {
    bool operator() (const myuser &l, const myuser &r) {
        if (l.out==r.out){
            return l.in > r.in;
        }
        return l.out < r.out;
    }
};

int main(){
    int N;
    std::cin >> N;
    myuser *users = new myuser[N];
    for (int i =0; i < N; ++i){
        int in, out;
        std::cin >> in >> out;
        users[i] = myuser(in, out);
    }
    MergeSort<myuser, UserComparator>(users, N, UserComparator());
    int count = 2;
    int curout = users[0].out;
    int curin = curout-1;
    bool prosmotr = false;
    for (int j = 1; j<N; ++j){
        // находим опорный элемент
        if (curout<=users[j].out && users[j].in<= curin){
            continue;
        }
        // проверяем его на просмотр рекламы от прошлой пачки посетителей
        if (users[j].in <= curout){
            prosmotr = true;
        }
        // дальше передвигаем наш текущий выход в зависимости от просмотра
        if (!prosmotr) {
            count+=2;
            curout = users[j].out;
            curin = curout-1;
        } else {
            curin = curout;
            count+=1;
            curout = users[j].out;
            prosmotr= false;
        }
    }
    std::cout << count;
    delete[] users;
    return 0;
}

