import java.util.*;

public class Solution03 {
    public int solution(int n, int[][] groups) {
        int answer = 0;
        int lightedCnt = 0;
        int start = 1;
        int end = 1;
        boolean isCheck;

        while (n != lightedCnt) {
            isCheck = false;

            for (int[] group : groups) {
                if (group[0] <= start && lightedCnt < group[1]) {
                    end = Math.max(end, group[1]);
                    isCheck = true;
                }
            }

            if (isCheck) {
                lightedCnt = end;
                start = end + 1;
            } else {
                lightedCnt += 1;
                start += 1;
            }

            answer += 1;
        }

        return answer;
    }
}
