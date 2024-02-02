export class MaxHeap<T> {
  arr: T[];
  maxSize: number;
  heapSize: number;

  constructor(maxSize: number) {
    this.maxSize = maxSize;
    this.heapSize = 0;
    this.arr = []
  }

  parent(i: number): number {
    return Math.floor((i - 1) / 2);
  }

  leftChild(i: number): number {
    return 2 * i + 1;
  }

  rightChild(i: number): number {
    return 2 * i + 2;
  }

  swap(lhs: number, rhs: number): void {
    const temp = this.arr[lhs];
    this.arr[lhs] = this.arr[rhs];
    this.arr[rhs] = temp;
  }

  insertKey(key: T): void {
    if (this.heapSize == this.maxSize) {
      throw new RangeError()
    }

    this.heapSize++;
    let index = this.heapSize - 1;
    this.arr.push(key)

    while (index != 0 && this.arr[this.parent(index)] < this.arr[index]) {
      this.swap(this.parent(index), index);
      index = this.parent(index);
    }
  }

  heapifyAt(i: number): void {
    const left = this.leftChild(i)
    const right = this.rightChild(i);
    let largest = i;
    if (left < this.heapSize && this.arr[left] > this.arr[largest]) {
      largest = left;
    }
    if (right < this.heapSize && this.arr[right] > this.arr[largest]) {
      largest = right;
    }
    if (largest != i) {
      this.swap(largest, i);
      this.heapifyAt(largest)
    }
  }

  removeMax(): void {
    if (this.heapSize == 0) {
      throw new RangeError()
    }
    if (this.heapSize == 1 ){ 
      this.arr.pop()
      this.heapSize--;
    }
    this.swap(0, this.heapSize - 1)
    this.arr.pop()
    this.heapifyAt(0)
  }
}
