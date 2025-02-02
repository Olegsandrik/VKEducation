// Использование АВЛ-дерева. Решение должно поддерживать передачу функции сравнения снаружи.
// Порядковые статистики. Дано число N и N строк.
// Каждая строка содержит команду добавления или удаления натуральных чисел,
// а также запрос на получение k-ой порядковой статистики. Команда добавления числа A задается положительным числом A,
// Использование АВЛ-дерева. Решение должно поддерживать передачу функции сравнения снаружи.
// Порядковые статистики. Дано число N и N строк.
// Каждая строка содержит команду добавления или удаления натуральных чисел,
// а также запрос на получение k-ой порядковой статистики. Команда добавления числа A задается положительным числом A,
// команда удаления числа A задается отрицательным числом “-A”. Запрос на получение k-ой порядковой статистики задается числом k.
// Требования: скорость выполнения запроса - O(log n).

#include <iostream>
#include <stack>

template<typename T>
struct mycomporator{
    bool operator()(const T& a, const T& b) const{
        return a < b;
    }
};

template <typename T, typename Comparator = mycomporator<T>>
class AvlTree
{
    struct Node
    {
        Node(const T &data)
                : data(data), left(nullptr), right(nullptr), height(1), weight(1)
        {
        }
        Node(const T &data, int h, int w)
                : data(data), left(nullptr), right(nullptr), height(h), weight(w)
        {
        }
        T data;
        Node *left;
        Node *right;
        size_t height;
        size_t weight;
    };

public:
    AvlTree(Comparator comp = Comparator())
            : root(nullptr), cmp(comp)
    {
    }

    ~AvlTree()
    {
        destroyTree(root);
    }

    AvlTree &operator=(const AvlTree &copy) {
        if (this == &copy){
            return *this;
        }
        if (copy.root == nullptr) {
            root = nullptr;
            cmp = copy.cmp;
            return *this;
        } else{
            destroyTree(root);
            root = new Node(copy.root->data, copy.root->height, copy.root->weight);
            std::stack<Node*> stack1new, stack2new, stack1copy, stack2copy;
            stack1new.push(root);
            stack1copy.push(copy.root);
            while (!stack1copy.empty()) {
                Node* current = stack1new.top();
                Node* currentcopy = stack1copy.top();
                stack1new.pop();
                stack2new.push(current);
                stack1copy.pop();
                stack2copy.push(currentcopy);
                current->weight = currentcopy->weight;
                current->height = currentcopy->height;
                if (current->left != nullptr) {
                    Node *newleft = new Node(currentcopy->left->data, currentcopy->left->height, currentcopy->left->weight);
                    current->left = newleft;
                    stack1new.push(current->left);
                    stack1copy.push(currentcopy->left);
                }
                if (current->right != nullptr) {
                    Node *newright = new Node(currentcopy->right->data, currentcopy->right->height, currentcopy->right->weight);
                    current->right = newright;
                    stack1new.push(current->right);
                    stack1copy.push(currentcopy->right);
                }
            }
            cmp = copy.cmp;
            return *this;
        }
    }

    void Add(const T &data)
    {
        root = addInternal(root, data);
    }

    bool Has(const T &data)
    {
        Node *tmp = root;
        while (tmp)
        {
            if (cmp(tmp->data, data)){
                tmp = tmp->right;
            }else if (cmp(data, tmp-data)){
                tmp = tmp->left;
            } else {
                return true;
            }
        }
        return false;
    }
    void Delete(const T &data)
    {
        root = deleteInternal(root, data);
    }

    T kthElem(size_t value) {
        if (!root) {
            return 0;
        }
        Node *result = kthElemInternal(root, value);
        if (result == nullptr) {
            return 0;
        } else {
            return result->data;
        }
    }
private:

    Node* kthElemInternal(Node *node, int value) {
        if (!node) {
            return nullptr;
        }
        //if (index < 1 || index > getWeight(node)) {return nullptr;}
        int delta = value - getWeight(node->left);
        if (delta > 0) {
            return kthElemInternal(node->right, delta - 1);
        } else if (delta < 0) {
            return kthElemInternal(node->left, value);
        } else {
            return node;
        }
    }
    void destroyTree(Node *node)
    {
        if (node)
        {
            destroyTree(node->left);
            destroyTree(node->right);
            delete node;
        }
    }

    Node* deleteInternal(Node *node, const T &data)
    {
        if (!node)
            return nullptr;
        if (cmp(node->data, data))
            node->right = deleteInternal(node->right, data);
        else if (cmp(data, node->data))
            node->left = deleteInternal(node->left, data);
        else
        {
            Node *left = node->left;
            Node *right = node->right;

            delete node;

            if (!right)
                return left;


            Node *max = nullptr;
            right = FindandRemoveMaxNode(right, max);

            max->left = left;
            max->right = right;

            return doBalance(max);
        }

        return doBalance(node);
    }

    Node* FindandRemoveMaxNode(Node *node, Node*& max) {
        if (!node->right) {
            max = node;
            return node->left;
        }
        node->right = FindandRemoveMaxNode(node->right, max);
        return doBalance(node);
    }

    Node* removeMin(Node *node)
    {
        if (!node->left)
            return node->right;
        node->left = removeMin(node->left);
        return doBalance(node);
    }

    Node* findMin(Node *node)
    {
        while (node->left)
            node = node->left;
        return node;
    }

    Node* addInternal(Node *node, const T &data)
    {
        if (!node)
            return new Node(data);

        if (!cmp(data, node->data)){
            node->right = addInternal(node->right, data);
        } else{
            node->left = addInternal(node->left, data);
        }
        return doBalance(node);
    }

    size_t getHeight(Node *node)
    {
        return node ? node->height : 0;
    }

    size_t getWeight(Node *node)
    {
        return node ? node->weight : 0;
    }

    void fixHeight(Node *node)
    {
        node->height = std::max(getHeight(node->left), getHeight(node->right)) + 1;
    }

    void fixWeight(Node *node)
    {
        node->weight = 1+ getWeight(node->right)+ getWeight(node->left);
    }

    int getBalance(Node *node)
    {
        return getHeight(node->right) - getHeight(node->left);
    }

    Node* rotateLeft(Node *node)
    {
        Node *tmp = node->right;
        node->right = tmp->left;
        tmp->left = node;
        fixHeight(node);
        fixHeight(tmp);
        fixWeight(node);
        fixWeight(tmp);
        return tmp;
    }

    Node* rotateRight(Node *node)
    {
        Node *tmp = node->left;
        node->left = tmp->right;
        tmp->right = node;
        fixHeight(node);
        fixHeight(tmp);
        fixWeight(node);
        fixWeight(tmp);
        return tmp;
    }

    Node* doBalance(Node *node)
    {
        fixHeight(node);
        fixWeight(node);
        switch (getBalance(node))
        {
            case 2:
            {
                if (getBalance(node->right) < 0)
                    node->right = rotateRight(node->right);
                return rotateLeft(node);
            }
            case -2:
            {
                if (getBalance(node->left) > 0)
                    node->left = rotateLeft(node->left);
                return rotateRight(node);
            }
            default:
                return node;
        }
    }
    Comparator cmp;
    Node *root;
};

int main(int argc, const char * argv[]) {
    AvlTree<int> AVLTREE;
    int N, operation, value;
    std::cin >> N;
    for (int i = 0; i<N; ++i){
        std::cin >> operation >> value;
        if (operation>=0){
            AVLTREE.Add(operation);
        }else{
            AVLTREE.Delete(-operation);
        }
        std::cout << AVLTREE.kthElem(value) << " ";
    }
    return 0;
}
