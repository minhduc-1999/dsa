import { MaxHeap } from "./heap";

describe("MaxHeap", () => {
  it("Swap should work properly", () => {
    const maxHeap = new MaxHeap<number>(10);
    maxHeap.insertKey(8);
    maxHeap.insertKey(4);
    maxHeap.insertKey(5);
    maxHeap.insertKey(1);
    maxHeap.insertKey(2);

    maxHeap.insertKey(10)

    expect(maxHeap.arr).toEqual([10, 4, 8, 1, 2, 5]);
  });

  it("Remove max should work", () => {
    const maxHeap = new MaxHeap<number>(10);
    maxHeap.insertKey(15);
    maxHeap.insertKey(5);
    maxHeap.insertKey(7);
    maxHeap.insertKey(2);
    maxHeap.insertKey(3)

    maxHeap.removeMax();

    expect(maxHeap.arr).toEqual([7, 5, 3, 2]);
  });
});
