#include <bits/stdc++.h>
#include <fstream>

using namespace std;

vector < vector <int> > g, num;
vector <int> used, ans, tin, tup;

int timer = 0;
void dfs(int v, int predok, int nomer){

    timer++;
    used[v] = 1;
    tin[v] = timer;
    tup[v] = timer;
    for (int i = 0; i < g[v].size(); i++){
        int to = g[v][i];
        if (used[to] == 0){
            dfs(to, v, num[v][i]);
            tup[v] = min(tup[v], tup[to]);
        }
        else if (nomer != num[v][i]){
            tup[v] = min(tup[v], tin[to]);
        }
    }
    if (tup[v] > tin[predok] && predok != -1){
        ans.push_back(nomer);
    }
}

int main(){
    int n, m,  x, y;
    cin >> n >> m;
    g.resize(n);
    used.resize(n, 0);
    tin.resize(n);
    tup.resize(n);
    num.resize(n);
    //cout<<"1"<<endl;
    for (int i = 0; i < m; i++){
        cin >> x >> y;
        g[x].push_back(y);
        g[y].push_back(x);
        num[x].push_back(i + 1);
        num[y].push_back(i + 1);
//        check.insert({x, y});
//        check.insert({y, x});
    }
    //cout<<"input correct"<<endl;
    for (int k = 0; k < n; k++){
        if (!used[k])
            dfs(k, -1, -1);
    }
   // cout<<"dfs correct"<<endl;
    sort(ans.begin(), ans.end());
    cout << ans.size() << endl;
//    for (int i = 0; i < ans.size(); i++){
//        cout << ans[i] << " ";
//    }
}
