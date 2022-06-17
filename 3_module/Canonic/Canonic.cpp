#include <iostream>
#include <vector>

using namespace std;

struct cond{
    vector <int> T;
    vector <string> F;
    bool used;
    int New;
};

int n, m, q, order;
vector <cond> A, A1;

void dfs(int i){
    A[i].used = true;
    A[i].New = order;
    order++;
    for (int j = 0; j < m; j++){
        if (!A[A[i].T[j]].used){
            dfs(A[i].T[j]);
        }
    }
}

int main(){
    int x;
    string s;
    cin >> n >> m >> q;
    A.resize(n);
    A1.resize(n);
    for (int i = 0; i < n; i++){
        A[i].T.resize(m);
        A1[i].T.resize(m);
        for (int j = 0; j < m; j++){
            //cin >> x;
            scanf("%d", &x);
            A[i].T[j] = x;
        }
    }
    for (int i = 0; i < n; i++){
        A[i].F.resize(m);
        A1[i].F.resize(m);
        for (int j = 0; j < m; j++){
            cin >> s;
            A[i].F[j] = s;
        }
    }
    dfs(q);
    for (int i = 0; i < n; i++){
        for (int j = 0; j < m; j++) {
            A1[A[i].New].T[j] = A[A[i].T[j]].New;
			A1[A[i].New].F[j] = A[i].F[j];
        }
    }
    cout << order << endl << m << endl << 0 << endl;
    for (int i = 0; i < order; i++){
        for (int j = 0; j < m; j++){
            cout << A1[i].T[j] << " ";
        }
        cout << endl;
    }
    for (int i = 0; i < order; i++){
        for (int j = 0; j < m; j++){
            cout << A1[i].F[j] << " ";
        }
        cout << endl;
    }
}
