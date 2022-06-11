


import java.util.*;
class TestPair{
    long i;
    long j;
    public TestPair(long x, long y){
        i = x;
        j = y;
    }
}
public class Dividers{

    static boolean check(ArrayList<Long> s, long size, long i, long j){
        for (int k = 0; k < size; k++){
            if (s.get(k) != i && s.get(k) != j && s.get(k) % i == 0 && j % s.get(k) == 0){
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        long n;
        n = in.nextLong();
        //int[] s = new int[n];
        ArrayList<Long> s = new ArrayList<>();
        int size = 0;
        for (long i = 1; i <= n; i++){
            if (n % i == 0){
                //s[size] = i;
                s.add(i);
                size++;
            }
        }
        ArrayList<TestPair> ans = new ArrayList<>();
        for (int i = 0; i < s.size(); i++){
            for (int j = i + 1; j < s.size(); j++){
                if (s.get(j) % s.get(i) == 0 && check(s, s.size(), s.get(i), s.get(j))){
                    ans.add(new TestPair(s.get(i), s.get(j)));
                }
            }
        }
        System.out.println("graph {");
        for (int i = 0; i < size; i++){
            System.out.printf("\t%d\n", s.get(i));
        }
        for (int i = 0; i < ans.size(); i++){
            System.out.println("\t" + ans.get(i).i + "--" + ans.get(i).j);
        }
        System.out.println("}");
    }
}

