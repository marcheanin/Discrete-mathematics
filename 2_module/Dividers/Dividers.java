

import java.util.*;
class TestPair{
    int i;
    int j;
    public TestPair(int x, int y){
        i = x;
        j = y;
    }
}
public class Dividers {


    static boolean check(int[] s, int size, int i, int j){
        for (int k = 0; k < size; k++){
            if (s[k] != i && s[k] != j && s[k] % i == 0 && j % s[k] == 0){
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n;
        n = in.nextInt();
        int[] s = new int[n];
        int size = 0;
        for (int i = 1; i <= n; i++){
            if (n % i == 0){
                s[size] = i;
                size++;
            }
        }
        ArrayList<TestPair> ans = new ArrayList<>();
        for (int i = 0; i < size; i++){
            for (int j = i + 1; j < size; j++){
                if (s[j] % s[i] == 0 && check(s, size, s[i], s[j])){
                    ans.add(new TestPair(s[i], s[j]));
                }
            }
        }
        System.out.println("graph {");
        for (int i = 0; i < size; i++){
            System.out.printf("\t%d\n", s[i]);
        }
        for (int i = 0; i < ans.size(); i++){
            System.out.println("\t" + ans.get(i).i + "--" + ans.get(i).j);
        }
        System.out.println("}");
    }
}
