class Node {
  constructor(data) {
    this.data = data;
    this.prev = null;
    this.next = null;
  }
}

export class DoublyLinkedLIst {
  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }
  insertAtBeginning(data) {
    const newNode = new Node(data);
    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      newNode.next = this.head;
      this.head.prev = newNode;
      this.head = newNode;
    }
    this.length++;
  }
  insertAtEnd(data) {
    const newNode = new Node(data);
    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      newNode.prev = this.tail;
      this.tail.next = newNode;
      this.tail = newNode;
    }
    this.length++;
  }
  insertAtPosition(data, position) {
    if (position < 0 || position > this.length) {
      return false;
    }
    if (position == this.length) {
      this.insertAtEnd(data);
    }
    if (position === 0) {
      this.insertAtBeginning(data);
    }
    const newNode = new Node(data);
    let current = this.head;
    for (let i = 0; i < position - 1; i++) {
      current = current.next;
    }
    newNode.next = current.next;
    newNode.prev = current;
    current.next.prev = newNode;
    current.next = newNode;
    this.length++;
    return true;
  }
  deleteNode(data) {
    if (!this.head) return false;
    let current = this.head;
    while (current) {
      if (current.data === data) {
        if (current === this.head && current === this.tail) {
          this.head = null;
          this.tail = null;
        } else if (current === this.head) {
          this.head = current.next;
          this.head.prev = null;
        } else if (current === this.tail) {
          this.tail = current.prev;
          this.tail.next = null;
        } else {
          current.prev.next = current.next;
          current.next.prev = current.prev;
        }
        this.length--;
        return true;
      }
      current = current.next;
    }
    return false;
  }

  traverseBackward() {
    let current = this.tail;
    while (current) {
      console.log(current.data);
      current = current.prev;
    }
  }

  search(data) {
    let current = this.head;
    let index = 0;
    while(current) {
        if(current.data === data) {
            return index;
        }
        current = current.next;
        index++;
    }
    return -1;
  }
}
