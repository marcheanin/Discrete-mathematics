#include <bits/stdc++.h>

using namespace std;

bool check(vector <int> s, int i, int j){
    for (int k = 0; k < s.size(); k++){
        if (s[k] != i && s[k] != j && s[k] % i == 0 && j % s[k] == 0){
            return 0;
        }
    }
    return 1;
}

int main(){
    long long n;
    cin >> n;
    vector <int> s;
    for (int i = 1; i <= n; i++){
        if (n % i == 0){
            s.push_back(i);
        }
    }
    vector <pair <int, int> > ans;
    for (int i = 0; i < s.size(); i++){
        for (int j = i + 1; j < s.size(); j++){
            if (s[j] % s[i] == 0 && check(s, s[i], s[j])){
                ans.push_back({s[i], s[j]});
            }
        }
    }
    for (int i = 0; i < s.size(); i++){
        cout << s[i] << endl;
    }

    for (auto i : ans){
        cout << i.first << "--" << i.second << endl;
    }
}
