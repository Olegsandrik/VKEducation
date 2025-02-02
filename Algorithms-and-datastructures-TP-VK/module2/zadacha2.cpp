// Дано число N < 106 и последовательность целых чисел из [-231..231] длиной N.
// Требуется построить бинарное дерево, заданное наивным порядком вставки.
// Т.е., при добавлении очередного числа K в дерево с корнем root, если root→Key ≤ K,
// то узел K добавляется в правое поддерево root; иначе в левое поддерево root.
// Требования: Рекурсия запрещена. Решение должно поддерживать передачу функции сравнения снаружи.
// 2_3. Выведите элементы в порядке post-order (снизу вверх)
// переделанный файл

#include <queue>
#include <stack>
#include "iostream"

template<typename T>
struct mycomporator{
    bool operator()(const T& a, const T& b) const{
        return a > b;
    }
};


template<typename T>
struct Node{
    Node(const T &data): data(data), left(nullptr), right(nullptr){}
    T data;
    Node *left;
    Node *right;
};


template<typename T, typename Comparator = std::less<T>>
class BinaryTree {
public:
    BinaryTree(): root(nullptr){}
    BinaryTree(T value): root(value){}

    ~BinaryTree(){
        DeleteTree();
    }


    BinaryTree &operator=(const BinaryTree &copy) {
        if (this == &copy){
            return *this;
        }
        if (copy.root == nullptr) {
            root = nullptr;
            cmp = copy.cmp;
            return *this;
        } else{
            destroyTree(root);
            root = new Node<T>(copy.root->data);
            std::stack<Node<T>*> stack1new, stack2new, stack1copy, stack2copy;
            stack1new.push(root);
            stack1copy.push(copy.root);
            while (!stack1copy.empty()) {
                Node<T>* current = stack1new.top();
                Node<T>* currentcopy = stack1copy.top();
                stack1new.pop();
                stack2new.push(current);
                stack1copy.pop();
                stack2copy.push(currentcopy);
                if (current->left != nullptr) {
                    Node<T> *newleft = new Node<T>(currentcopy->left->data);
                    current->left = newleft;
                    stack1new.push(current->left);
                    stack1copy.push(currentcopy->left);
                }
                if (current->right != nullptr) {
                    Node<T> *newright = new Node<T>(currentcopy->right->data);
                    current->right = newright;
                    stack1new.push(current->right);
                    stack1copy.push(currentcopy->right);
                }
            }
            cmp = copy.cmp;
            return *this;
        }
    }

    void Add(T value) {
        auto* nownode = new Node<int>(value);
        Node<T>* current = root;
        while (true) {
            if (root == nullptr) {
                root = nownode;
                return;
            }
            if (cmp(current->data, nownode->data)) {
                if (current->left == nullptr) {
                    current->left = nownode;
                    return;
                } else {
                    current = current->left;
                }
            } else {
                if (current->right == nullptr) {
                    current->right = nownode;
                    return;
                } else {
                    current = current->right;
                }
            }
        }
    }

    void BFS(std::vector<T>& ans) {
        if (root == nullptr) {
            return;
        }
        std::stack<Node<T>*> stack1, stack2;
        stack1.push(root);

        while (!stack1.empty()) {
            Node<T>* current = stack1.top();
            stack1.pop();
            stack2.push(current);

            if (current->left != nullptr) {
                stack1.push(current->left);
            }
            if (current->right != nullptr) {
                stack1.push(current->right);
            }
        }
        while (!stack2.empty()) {
            Node<T>* current = stack2.top();
            ans.push_back(current->data);
            stack2.pop();
        }
    }
private:
    Node<T>* root;
    Comparator cmp;
    void DeleteTree(){
        if (root == nullptr) {
            return;
        }
        std::vector<Node<T>*> queue;
        queue.push_back(root);
        while(!queue.empty()){
            root = queue[0];
            queue.erase(queue.begin());
            if (root->left != nullptr) {
                queue.push_back(root->left);
            }
            if (root->right != nullptr) {
                queue.push_back(root->right);
            }
            delete root;
        }
    }
};


int main(){
    int N, value;
    std::cin >> N >> value;
    BinaryTree<int, mycomporator<int>> tree;
    tree.Add(value);
    for (int i=0; i<N-1; ++i){
        std::cin >> value;
        tree.Add(value);
    }
    std::vector<int> ans;
    tree.BFS(ans);
    for (int i = 0; i < ans.size(); ++i){
        std::cout << ans[i] << " ";
    }
}
