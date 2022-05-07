#include <bits/stdc++.h>

using namespace std;

int n, m;
vector <vector <int> > g;

void dfs(int v, int used[], int c){
    //cout << v << endl;
    used[v] = c;
    for (auto i : g[v]){
        if (!used[i]){
            dfs(i, used, c);
        }
    }
}

struct component {
    int minv = -1;
    set <int> v;
    multiset <pair <int, int> > e;
};

bool compare(component a, component b){
    if (a.v.size() != b.v.size()) return a.v.size() < b.v.size();
    else if (a.e.size() != b.e.size()) return a.e.size() < b.e.size();
    else return a.minv > b.minv;
}

int main(){
    cin >> n >> m;
    g.resize(n);
    vector <pair <int, int > > edges (m);
    int x, y;
    for (int i = 0; i < m; i++){
        cin >> x >> y;
        edges[i] = {x, y};
        g[x].push_back(y);
        g[y].push_back(x);
    }
    int used[n] = {};
    int c = 1;
    for (int i = 0; i < n; i++){
        if (!used[i]){
            dfs(i, used, c);
            c++;
        }
    }
    for (int i = 0; i < n; i++) used[i]--;
    //for (int i = 0; i < n; i++) cout << used[i] << " ";
    //cout << endl;
    vector <component> s (c);
    for (int i = 0; i < n; i++){
        (s[used[i]].v).insert(i);
        if (i < s[used[i]].minv || s[used[i]].minv == -1) s[used[i]].minv = i;
    }
    for (int i = 0; i < c; i++){
        for (int j = 0; j < m; j++){
            if (s[i].v.count(edges[j].first) > 0 || s[i].v.count(edges[j].second) > 0){
                s[i].e.insert(edges[j]);
            }
        }
    }
    sort(s.rbegin(), s.rend(), compare);
    auto ans = s[0];
    //cout << ans.v.size() << " " << ans.e.size() << " " << ans.minv << endl;
    for (auto i : ans.v) {
        cout << i << endl;
    }
    for (auto i : ans.e){
        cout << i.first << "--" << i.second << endl;
    }
}

/*
14
14
0 1
2 3
2 4
4 3
5 6
6 7
10 8
10 9
9 8
8 9
11 12
11 13
12 13
13 11
*/

