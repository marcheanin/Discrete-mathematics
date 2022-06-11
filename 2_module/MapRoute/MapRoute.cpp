#include <iostream>
#include <vector>
#include<algorithm>
#include<cmath>
using namespace std;
const int inf = 1e9;


int main() {
    int n, s, f;
    cin >> n >> s >> f;
    s--; f--;
    vector <int> d(n, inf), p(n);
    vector<int> u(n);
    vector < vector<pair<int, int> > > g(n);
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            int q;
            scanf("%d", &q);
            if (i > 0) {

            }
        }
    }
    d[s] = 0;
    for (int i = 0; i < n; ++i) {
        int v = -1;
        for (int j = 0; j < n; ++j)
            if (!u[j] && (v == -1 || d[j] < d[v]))
        v = j;
        if (d[v] == inf) break;
        u[v] = true;
        for (auto j : g[v]) {
            int to = j.first;
            int len = j.second;
            if (d[v] + len < d[to]) {
                d[to] = d[v] + len;
                p[to] = v;
            }
        }

    }
    cout << (d[f] == inf ? -1 : d[f]) << endl;
    return 0;
}
