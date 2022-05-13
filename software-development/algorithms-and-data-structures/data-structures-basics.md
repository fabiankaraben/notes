# Data Structures Basics

## Tries

The trie, pronounced “try”, is a tree that specializes in storing data that can be represented as a collection, such as English words.

Prefix matching example (bad code):

```dart
class EnglishDictionary {
    final List<String> words = [];
    List<String> lookup(String prefix) {
        return words.where((word) {
            return word.startsWith(prefix);
        }).toList();
    }
} 
```

This algorithm is reasonable if the number of elements in the words list is small. But if you’re dealing with more than a few thousand words, the time it takes to go through the words list will be unacceptable.

The time complexity of lookup is O(k × n), where k is the longest string in the collection, and n is the number of words you need to check.

Tries provide great performance metrics for prefix matching.

Tries are relatively memory efficient since individual nodes can be shared between many different values. For example, “car,” “carbs,” and “care” can share the first three letters of the word.

## Heaps

Heaps are another classical tree-based data structure with special properties to quickly fetch the largest or smallest element. 

A heap is a complete binary tree, also known as a binary heap, that can be constructed using a list. 

Heaps come in two flavors:

1. Max-heap, in which elements with a higher value have a higher priority.
   - Parent nodes must always contain a value that is greater than or equal to the value in its children.
   - The root node will always contain the highest value. 
2. Min-heap, in which elements with a lower value have a higher priority.
   - Parent nodes must always contain a value that is less than or equal to the value in its children.
   - The root node will always contain the lowest value. 

Some practical applications of a heap include: 

- Calculating the minimum or maximum element of a collection.
- Implementing the heapsort algorithm.
- Constructing a priority queue.
- Building graph algorithms that use a priority queue, like Dijkstra’s algorithm. 

The heap data structure is good for maintaining the highest- or lowest-priority element. 

## Priority Queues

A priority queue is another version of a queue in which elements are dequeued in priority order instead of FIFO order. 

A priority queue is often used to retrieve elements in **priority order**. 

A priority queue can be either of these two:

1. Max-priority, in which the element at the front is always the largest.
2. Min-priority, in which the element at the front is always the smallest. 

This is similar with heap data structure. A priority queue creates a layer of abstraction by focusing on the key operations of a queue and leaving out the additional functionality provided by a heap. This makes the priority queue’s intent clear and concise. Its only job is to enqueue and dequeue elements, nothing else.

A priority queue has the same operations as a regular queue, so only the implementation will differ: enqueue, dequeue, isEmpty, peek.

You can create a priority queue in the following ways:

1. **Sorted list**: This is useful to obtain the maximum or minimum value of an element in O(1) time. However, insertion is slow and will require O(n) time since you have to first search for the insertion location and then shift every element after that location. 
2. **Balanced binary search tree**: This is useful in creating a double-ended priority queue, which features getting both the minimum and maximum value in O(log n) time. Insertion is better than a sorted list, also O(log n).
3. **Heap**: This is a natural choice for a priority queue. A heap is more efficient than a sorted list because a heap only needs to be partially sorted. Inserting and removing from a heap are O(log n) while simply querying the highest priority value is O(1). 

Some practical applications of a priority queue include:

- **Dijkstra’s algorithm**, which uses a priority queue to calculate the minimum cost.
- **A* pathfinding algorithm**, which uses a priority queue to track the unexplored routes that will produce the path with the shortest length.
- **Heapsort**, which can be implemented using a priority queue.
- **Huffman coding** that builds a compression tree. A min-priority queue is used to repeatedly find two nodes with the smallest frequency that do not yet have a parent node. 
 
These are just some of the use cases, but priority queues have many more applications as well. 

