class Node {
  constructor(data) {
    this.data = data;
    this.next = null;
    this.prev = null;
  }
}

export class DoublyLinkedListJS {
  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }
  insertAtBeginning(data) {
    const newnode = new Node(data);
    if (!this.head) {
      this.head = newnode;
      this.tail = newnode;
    } else {
      newnode.next = this.head;
      this.head.prev = newnode;
      this.head = newnode;
    }
    this.length++;
  }
  insertAtEnd(data) {
    const newnode = new Node(data);
    if (!this.head) {
      this.head = newnode;
      this.tail = newnode;
    } else {
      newnode.prev = this.tail;
      this.tail.next = newnode;
      this.tail = newnode;
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
    const newnode = new Node(data);
    let current = this.head;
    for (let i = 0; i < position - 1; i++) {
      current = current.next;
    }
    newnode.next = current.next;
    newnode.prev = current;
    current.next.prev = newnode;
    current.next = newnode;
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

  deleteAndPopEndNode() {
    if (!this.head || !this.tail) return false;
    if (this.head === this.tail) {
      this.head = null
      this.tail = null
      return null;
    }
    let tailData = this.tail.data;
    let newTail = this.tail.prev
    newTail.next = null
    this.tail = newTail
    this.length--;
    return tailData;
  }
  sizereturner() {
    return this.length;
  }

  traverseBackward() {
    let current = this.tail;
    let value;
    while (current) {
      value = current.data;
      current = current.prev;
    }
    return value;
  }

  search(data) {
    let current = this.head;
    let index = 0;
    while (current) {
      if (current.data === data) {
        return index;
      }
      current = current.next;
      index++;
    }
    return -1;
  }
}
