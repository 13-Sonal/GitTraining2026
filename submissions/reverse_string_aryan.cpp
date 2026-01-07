#include <bits/stdc++.h>
using namespace std;

string reverse_string (string s)
{
    int i=0;
    int j=s.size()-1;
    while (i <= j)
    {
        swap(s[i],s[j]);
        i++;
        j--;
    }
    return s;
}

int main ()
{
    string s;
    cout << "Enter a string : ";
    cin >> s;
    cout << endl << "Reversed string : " << reverse_string(s);
    return 0;
}