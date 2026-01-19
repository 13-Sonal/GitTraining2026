#include <iostream>
using namespace std;

int main(){
    int n;
    cin>>n;
    long long factorial=1;
    for(int i=n;i>=1;i--){
        factorial*=i;
    }
    cout<<"Factorial of "<<n<<" is "<<factorial<<endl;
}