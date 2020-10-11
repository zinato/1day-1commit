import java.util.*;

public class Solution03 {
    public int solution(int k, int[] score) {
        Map<Integer, Integer> map = new HashMap<>();
        Set<Integer> gapFound = new HashSet<>();
        int n = score.length;

        for(int i = 0 ; i < n - 1 ; i++) {
            int gap = score[i] - score[i+1];
            map.merge(gap, 1, Integer::sum);
        }

        for(Integer key : map.keySet()) {
            if(map.get(key) >= k) {
                gapFound.add(key);
            }
        }


        Set<Integer> indexFound = new HashSet<>();

        for(int i = 1 ; i < n ; i++) {
            int gap = score[i - 1] - score[i];

            if(gapFound.contains(gap)) {
                indexFound.add(i);
                indexFound.add(i + 1);
            }
        }

        int answer = n - indexFound.size();
        return answer;
    }
}
