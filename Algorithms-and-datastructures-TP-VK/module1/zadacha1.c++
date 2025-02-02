#include <iostream>
// Если в числе содержится только один бит со значением 1, записать в выходной поток OK. Иначе записать FAIL
int scanner(unsigned long n){
    int result = {0b000};
    while (n){
        int x = n & 1;
        if (x) {
            result = result << 1;
            result = result | x;
        }
        n = n>>1;
    }
    return result;
}
void printresult(int result){
    if (result ^ 1){
        std::cout << "FAIL" << std::endl;
    } else{
        std::cout << "OK" << std::endl;
    }
}

int main() {
    unsigned long int n;
    std::cin >> n;
    int result = scanner(n);
    printresult(result);
    return 0;
}


