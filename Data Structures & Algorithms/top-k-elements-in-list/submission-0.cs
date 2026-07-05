public class Solution {
    public int[] TopKFrequent(int[] nums, int k)
    {
        var freqs = new Dictionary<int, int>();
        foreach (var number in nums)
        {
            freqs.TryAdd(number, 0);
            freqs[number]++;
        }

        var buckets = new List<int>?[nums.Length + 1];
        foreach (var number in freqs.Keys)
        {
            var freq = freqs[number];
            buckets[freq] ??= [];
            buckets[freq].Add(number);
        }

        var result = new List<int>();
        for (var i = buckets.Length - 1; i >= 0 && result.Count < k; i--)
        {
            var bucket = buckets[i];
            if (bucket is null) continue;
            
            for (var j = 0; j < bucket.Count && result.Count < k; j++)
            {
                result.Add(bucket[j]);
            }
        }
        
        return result.ToArray();
    }
}