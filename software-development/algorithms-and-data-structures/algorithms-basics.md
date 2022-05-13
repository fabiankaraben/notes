# Algorithms Basics

## Binary Search

Binary search is one of the most efficient searching algorithms with a time complexity of O(log n). 

Two conditions need to be met for the type of binary search that this chapter describes:

- The collection must be sorted.
- The underlying collection must be able to perform random index lookup in constant time. For example Dart List or Go Array/Slice.

> Note: Sometimes it may be beneficial to sort a collection to leverage the binary search capability for looking up elements.

> Dart Note: The indexOf method on List uses a linear search with O(n) time complexity. Binary search has O(log n) time complexity, which scales much better for large data sets if you are doing repeated lookups.

## O(n²) Sorting Algorithms

O(n²) time complexity isn’t great performance, but one advantage of these algorithms is that they have constant O(1) space complexity, making them attractive for certain applications where memory is limited. For small data sets, these sorting algorithms compare very favorably against more complex sorts. 

### Bubble sort

One of the most straightforward sorts is the bubble sort, which repeatedly compares adjacent values and swaps them, if needed, to perform the sort. Therefore, the larger values in the set will “bubble up” to the end of the collection.

Bubble sort has a best time complexity of O(n) if it’s already sorted, and a worst and average time complexity of O(n²), making it one of the least appealing sorts in the known universe.

### Selection sort

Selection sort follows the basic idea of bubble sort but improves this algorithm by reducing the number of swap operations. Selection sort will only swap at the end of each pass.

During each pass, selection sort will find the lowest unsorted value and swap it into place.

Selection sort has a best, worst and average time complexity of O(n²), which is fairly dismal. It’s a simple one to understand, though, and it does perform better than bubble sort!

### Insertion sort

Insertion sort is a more useful algorithm. Like bubble sort and selection sort, insertion sort has an average time complexity of O(n²), but the performance of insertion sort can vary. The more the data is already sorted, the less work it needs to do. Insertion sort has a best time complexity of O(n) if the data is already sorted. 

Dart itself uses the insertion sort. For lists with 32 or fewer elements, the sort method defaults to an insertion sort. Only for larger collections does Dart make use of a different sorting algorithm. 

The idea of insertion sort is similar to how many people sort a hand of cards. You start with the card at one end and then go through the unsorted cards one at a time, taking each one as you come to it and inserting it in the correct location among your previously sorted cards. 

It’s worth pointing out that the best-case scenario for insertion sort occurs when the sequence of values is already in sorted order and no left shifting is necessary.

## Merge Sort

The idea behind merge sort is to divide and conquer — to break up a big problem into several smaller, easier-to-solve problems and then combine the solutions into a final result.

With a time complexity of O(n log n), merge sort is one of the fastest of the general-purpose sorting algorithms.

- Merge sort is in the category of divide-and-conquer algorithms. 
- Merge sort works by splitting the original list into many individual lists of length one. It then merges pairs of lists into larger and larger sorted lists until the entire collection is sorted. 
- There are many implementations of merge sort, and you can have different performance characteristics depending on the implementation. 
- Merge sort has a time complexity of O(n log n). It does not sort in place, so the space complexity is also O(n log n), but can be O(log n) if optimized. 
- Merge sort is a stable sorting algorithm. 

