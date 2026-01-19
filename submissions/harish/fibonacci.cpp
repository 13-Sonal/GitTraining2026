#include <iostream>
using namespace std;

// Fibonacci series cpp program
int main() {
    int n;
    cin >> n;

    if (n < 0) {
        cout << "Invalid input. Please enter a non-negative number." << endl;
        return 0;
    }

    long long a = 0, b = 1;

    for (int i = 1; i <= n; i++) {
        cout << a << " ";
        long long next = a + b;
        a = b;
        b = next;
    }

    return 0;
}
