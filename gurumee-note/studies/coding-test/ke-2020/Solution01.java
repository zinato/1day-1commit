import java.util.List;

public class Solution01 {
    public static int splitIntoTwo(List<Integer> arr) {
        int answer = 0;
        int leftSum = 0;
        int all = 0;

        for(int i = 0 ; i < arr.size() ; i++) {
            all += arr.get(i);
        }

        for(int i = 0 ; i < arr.size() - 1 ; i++) {
            leftSum += arr.get(i);
            if(leftSum > all - leftSum) {
                answer++;
            }
        }

        return answer;
    }
}
