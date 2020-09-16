import java.util.List;

public class Solution03 {
    public static int segment(int x, List<Integer> space) {
        int start = 0;
        int end = space.size() - x;
        int nextIndex = -1;
        int answer = Integer.MIN_VALUE;

        while(start <= end) {
            int min = Integer.MAX_VALUE;

            for(int i = start ; i < start + x ; i++) {
                if(min > space.get(i)) {
                    min = space.get(i);
                    nextIndex = i;
                }
            }

            if(start == nextIndex) {
                start++;
            } else {
                start = nextIndex;
            }

            answer = Math.max(answer, min);
        }
        return answer;
    }
}
