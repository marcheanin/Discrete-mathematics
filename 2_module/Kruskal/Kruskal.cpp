#include <bits/stdc++.h>

using namespace std;

struct edge {
    int x, y;
    double dist;
};

double count_dist(pair <int, int> a, pair <int, int> b){
    return sqrt((a.first - b.first) * (a.first - b.first) + (a.second - b.second) * (a.second - b.second));
}

bool compare(edge a, edge b){
    return a.dist < b.dist;
}

int main(){
    int n;
    cin >> n;
    vector <pair <int, int > > points (n);
    for (int i = 0; i < n; i++){
        cin >> points[i].first >> points[i].second;
    }
    vector <edge> edges;
    for (int i = 0; i < n; i++){
        for (int j = i + 1; j < n; j++){
            edges.push_back({i, j, count_dist(points[i], points[j])});
            //cout << edges[i].x.first << " " << edges[j].x.second << edges[i].x.first << " " << edges[j].x.second << << " " << edges.dist << endl;
        }
    }
    sort(edges.begin(), edges.end(), compare);
    for (int i = 0; i < edges.size(); i++){
        cout << edges[i].dist << endl;
    }
    vector <int> tree_id(n);
    for (int i = 0; i < n; i++){
        tree_id[i] = i;
    }
    int m = edges.size();
    double cost = 0;
    for (int i = 0; i < m; i++){
        int a = edges[i].x, b = edges[i].y;
        double l = edges[i].dist;
        if (tree_id[a] != tree_id[b]){
            cost += l;
            int old_id = tree_id[b], new_id = tree_id[a];
            for (int j = 0; j < n; j++){
                if (tree_id[j] == old_id){
                    tree_id[j] = new_id;
                }
            }
        }
    }
    cout << cost;
}
