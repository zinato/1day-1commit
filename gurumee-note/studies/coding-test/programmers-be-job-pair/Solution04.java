import java.util.*;

class Solution04 {
    static class SimpleEntry implements Comparable<SimpleEntry> {
        public final String key;
        public final int value;

        public SimpleEntry(String key, int value) {
            this.key = key;
            this.value = value;
        }

        @Override
        public int compareTo(SimpleEntry o) {
            int result = Integer.compare(value, o.value);

            if (result == 0) {
                result = o.key.compareTo(key);
            }

            return result;
        }
    }
    
    public String solution(String[] votes, int k) {
        Map<String, Integer> map = new HashMap<>();

        for (String v : votes) {
            map.putIfAbsent(v, 0);
            map.put(v, map.get(v) + 1);
        }

        List<SimpleEntry> entries = map.entrySet().stream()
                .map(entry -> new SimpleEntry(entry.getKey(), entry.getValue()))
                .sorted(Comparator.reverseOrder())
                .collect(Collectors.toList());
        PriorityQueue<SimpleEntry> pq = new PriorityQueue<>(entries);
        int sum = entries.stream()
                .limit(k)
                .map(entry -> entry.value)
                .reduce(0, Integer::sum);
        int last = 0;
        String answer = "";

        while (!pq.isEmpty() && last + pq.peek().value < sum) {
            SimpleEntry e = pq.remove();
            answer = e.key;
            last += e.value;
        }

        return answer;
    }
}