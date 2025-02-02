// Реализуйте структуру данных типа “множество строк” на основе динамической хеш-таблицы с открытой адресацией.
// Хранимые строки непустые и состоят из строчных латинских букв.
// Хеш-функция строки должна быть реализована с помощью вычисления значения многочлена методом Горнера.
// Начальный размер таблицы должен быть равным 8-ми. Перехеширование выполняйте при добавлении элементов в случае,
// когда коэффициент заполнения таблицы достигает 3/4.
// Структура данных должна поддерживать операции добавления строки в множество,
// удаления строки из множества и проверки принадлежности данной строки множеству.
// Для разрешения коллизий используйте квадратичное пробирование. i-ая проба
// g(k, i)=g(k, i-1) + i (mod m). m - степень двойки.

// nil - data: "" && isdeleted: false
// del - data: "" && isdeleted: true
// node - data: "somestring" && isdeleted: false
#include <iostream>
#include <vector>

const size_t DEFAULT_SIZE = 8;
const double MAX_ALPHA = 0.75;

template <typename T>
struct HashtableNode
{
    T data;
    bool isdeleted= false;
};

class StringHasher
{
public:
    StringHasher(size_t prime = 997)
            : prime(prime)
    {
    }

    size_t operator()(const std::string &str)
    {
        size_t hash = 0;
        for (int i = 0; i < str.size(); i++)
        {
            hash = hash * prime + str[i];
        }
        return hash;
    }

private:
    size_t prime;
};

template <typename T, typename Hasher>
class Hashtable
{
public:
    Hashtable(size_t initial_size = DEFAULT_SIZE): size(0), table(initial_size){}

    ~Hashtable(){
        table.clear();
        table.shrink_to_fit();
    }
    Hashtable& operator=(const Hashtable& copy){
        if (this != &copy){
            if (!table.empty()) {
                table.clear();
                table.shrink_to_fit();
            }
            size = copy.size;
            table = copy.table;
            hasher = copy.hasher;
        }
        return *this;
    }

    bool Add(const HashtableNode<T> &node){
        if (size > table.size() * MAX_ALPHA)
        {
            grow();
        }
        long long firstdelete = -1;
        long long hash = hasher(node.data)%table.size();
        for (int i = 0; i < table.size(); ++i){
            if (table[hash].data==node.data && !table[hash].data.empty()){
                return false;
            }
            if (table[hash].data.empty() && table[hash].isdeleted && firstdelete<0){
                firstdelete = hash;
            }
            if (!table[hash].isdeleted && table[hash].data == ""){
                table[hash].data = node.data;
                table[hash].isdeleted = false;
                size++;
                return true;
            }
            hash = (hash + i*i) % table.size();
        }
        table[hash].data = node.data;
        table[hash].isdeleted = false;
        size++;
        return true;
    }

    bool Has(const HashtableNode<T> &node){
        long long hash = hasher(node.data)%table.size();
        for (int i=0;i<table.size(); ++i){
            if (table[hash].data==node.data){
                return true;
            }
            if (table[hash].data.empty() && !table[hash].isdeleted){
                return false;
            }
            hash = (hash+i*i)%table.size();
        }
        return false;
    }

    bool Delete(const HashtableNode<T> &node){
        long long hash = hasher(node.data)%table.size();
        for (int i=0;i<table.size();++i){
            if (table[hash].data==node.data && !table[hash].data.empty()){
                table[hash].isdeleted = true;
                table[hash].data = "";
                return true;
            }
            if (table[hash].data.empty() && !table[hash].isdeleted){
                return false;
            }
            hash = (hash+i*i)%table.size();
        }
        return false;
    }

private:
    void grow(){
        std::vector<HashtableNode<T>> newtable(2*table.size());
        for (int i = 0; i < table.size(); ++i){
            if (!table[i].data.empty() || table[i].data.empty()&&table[i].isdeleted){
                HashtableNode<T> node = table[i];
                unsigned long hash = hasher(node.data)%newtable.size();
                for (int i = 0; i < newtable.size(); ++i) {
                    if (newtable[hash].data == node.data){
                        break;
                    }
                    if (newtable[hash].data.empty()){
                        newtable[hash] = node;
                        break;
                    }
                    hash = (hash + i*i) % newtable.size();
                }
            }
        }
        table = newtable;
    }
    std::vector<HashtableNode<T>> table;
    size_t size=0;
    Hasher hasher;
};

int main() {
    Hashtable<std::string, StringHasher> table(DEFAULT_SIZE);
    char opiration;
    HashtableNode<std::string> node;
    while (std::cin >> opiration >> node.data) {
        switch (opiration) {
            case '?': {
                std::cout << (table.Has(node) ? "OK" : "FAIL") << std::endl;
                break;
            }
            case '+': {
                std::cout << (table.Add(node) ? "OK" : "FAIL") << std::endl;
                break;
            }
            case '-': {
                std::cout << (table.Delete(node) ? "OK" : "FAIL") << std::endl;
                break;
            }
            default:
                return 0;
        }
    }
    return 0;
}
