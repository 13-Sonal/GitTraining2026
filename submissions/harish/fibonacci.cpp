#include <iostream>
using namespace std;

#Fibonacci series cpp program
int main() {
    int n;
    cin >> n;

    int a = 0, b = 1;

    for (int i = 1; i <= n; i++) {
        cout << a << " ";
        int next = a + b;
        a = b;
        b = next;
    }

    return 0;
}
