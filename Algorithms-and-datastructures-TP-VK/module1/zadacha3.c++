#include <iostream>
//Реализовать очередь с помощью двух стеков.
//Требования: Очередь должна быть реализована в виде класса. Стек тоже должен быть реализован в виде класса.
// a = 2 - pop front
// a = 3 - push back
class Stack{
public:
    Stack()
    {
        capacity = 1;
        arr = new int[capacity];
        top = -1;
    }
    ~Stack() {
        delete[] arr;
    }
    Stack& operator= (const Stack& copy){
        if (this!=&copy){
            if (arr != nullptr) {
                delete[] arr;
            }
            arr = new int[copy.capacity];
            for (int i=0; i<copy.capacity; ++i){
                arr[i]=copy.arr[i];
            }
            top = copy.top;
        }
        return *this;
    }
    void push(int znach){
        if (top+1==capacity) {
            capacity *= 2;
            int *newArr = new int[capacity];
            for (int i =0; i<=top; i++) {
                newArr[i] = arr[i];
            }
            delete[] arr;
            arr = newArr;
        }
        arr[++top]=znach;
    }
    int pop() {
        if (isEmpty()) {
            return 0;
        }
        return arr[top--];
    }
    bool isEmpty() const{
        if (top==-1) {
            return true;
        } else{
            return false;
        }
    }
private:
    int *arr;
    int top;
    int capacity;
};

class Queue{
public:
    Queue(){
        in = new Stack();
        out = new Stack();
    }
    ~Queue(){
        delete in;
        delete out;
    }
    Queue& operator=(const Queue& copy){
        if (this != &copy){
            if (in != nullptr) {
                delete[] in;
            }
            if (out != nullptr) {
                delete[] out;
            }
            in = copy.in;
            out = copy.out;
        }
        return *this;
    }
    void Enqueue(int n)
    {
        in->push(n);
    }
    int Dequeue()
    {
        if (out->isEmpty())
        {
            while(!in->isEmpty())
            {
                out->push(in->pop());
            }
        }

        if (!out->isEmpty())
            return out->pop();
        else
            throw std::underflow_error("Queue is empty");
    }
private:
    Stack *in;
    Stack *out;
};
int printans(int n){
    int ans = 1;
    Queue queue;
    for (int i = 0; i<n; ++i){
        int command;
        int value;
        std::cin >> command >> value;
        if (command == 2){
            try {
                int correctval = queue.Dequeue();
                if (correctval != value){
                    ans = 0;
                }
            }
            catch (std::underflow_error& e) {
                if (value!=-1){
                    ans = 0;
                }
            }
        } else{
            queue.Enqueue(value);
        }
    }
    return ans;
}

int main() {
    int n;
    std::cin >> n;
    if (printans(n)){
        std::cout << "YES";
    } else {
        std::cout << "NO";
    }
    return 0;
}
