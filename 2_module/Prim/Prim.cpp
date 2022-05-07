#include <bits/stdc++.h>

using namespace std;

int main(){
    int n, m;
    cin >> n >> m;
    vector <vector <pair <int, int> > > g (n);
    int x, y, dist;
    for (int i = 0; i < m; i++){
        cin >> x >> y >> dist;
        g[x].push_back({y, dist});
        g[y].push_back({x, dist});
    }
    int used[n] = {};
    int used_v = 1, ans = 0;
    used[0] = 1;
    while(used_v < n){
        int mun = 1e9, vmun;
        for (int i = 0; i < n; i++){
            if (used[i]){
                for (auto j : g[i]){
                    if (!used[j.first]){
                        if (j.second < mun){
                            mun = j.second;
                            vmun = j.first;
                        }
                    }
                }
            }
        }
        used[vmun] = 1;
        ans += mun;
        used_v++;
    }
    cout << ans;
}
