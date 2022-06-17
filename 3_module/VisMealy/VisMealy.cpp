#include <iostream>
#include <vector>

using namespace std;

int main(){
    int n, m, q;
    cin >> n >> m >> q;
    vector <vector <int> > gm (n, vector <int> (m));
    vector <vector <string> > em (n, vector <string> (m));
    for (int i = 0; i < n; i++){
        for (int j = 0; j < m; j++){
            cin >> gm[i][j];
        }
    }
    for (int i = 0; i < n; i++){
        for (int j = 0; j < m; j++){
            cin >> em[i][j];
        }
    }
    cout << "digraph{" << endl << "\trankdir = LR" << endl;
    for (int i = 0; i < n; i++) {
		for (int j = 0; j < m; j++) {
            cout << "\t"  << i << " -> " << gm[i][j] << " [label = \"" << char(97+j) << "(" << em[i][j] << ")\"]\n";
			//printf("\t%d -> %d [label = \"%c(%c)\"]\n", i, gm[i][j], alph[j], em[i][j]);
		}
		q++;
	}
	cout << "}";
}
