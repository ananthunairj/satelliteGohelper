import { Searchresult } from "./interfaces";

class Node<T> {
  public next: Node<T> | null = null;
  public prev: Node<T> | null = null;
  constructor(public data: T) {}
}

interface IDoublyLinkedList<T> {
  insertAtBeginning(data: T): void;
  insertAtEnd(data: T): void;
  insertAtPosition(data: T, position: number): boolean;
  deleteNode(data: T): boolean;
  traverseBackward(): T[];
  search(data: T): Searchresult;
}

export class DoublyLinkedLIst<T> implements IDoublyLinkedList<T> {
  private head: Node<T> | null = null;
  private tail: Node<T> | null = null;
  private length: number;

  private constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  public static createLinkedlist<T>() : IDoublyLinkedList<T>  {
    return new DoublyLinkedLIst<T>();
  }

  public insertAtBeginning(data: T): void {
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

  public insertAtEnd(data: T): void {
    const newnode = new Node(data);
    if (!this.head || !this.tail) {
      this.head = newnode;
      this.tail = newnode;
    } else {
      newnode.prev = this.tail;
      this.tail.next = newnode;
      this.tail = newnode;
    }
    this.length++;
  }

  public insertAtPosition(data: T, position: number): boolean {
    if (position < 0 || position > this.length) {
      return false;
    }
    if (position == this.length) {
      this.insertAtEnd(data);
      return true;
    }
    if (position === 0) {
      this.insertAtBeginning(data);
      return true;
    }
    const newnode = new Node(data);
    let current = this.head;
    if (current == null) {
      return false;
    }
    for (let i = 0; i < position - 1; i++) {
      if(!current.next) return false
      current = current.next;
    }
    newnode.next = current.next;
    newnode.prev = current;
    if (current.next) current.next.prev = newnode;
    current.next = newnode;
    this.length++;
    return true;
  }

  public deleteNode(data: T): boolean {
      if (!this.head) return false;
      let current = this.head;
      while(current) {
        if(current.data === data) {
          if (current === this.head && current === this.tail) {
            this.head = null;
            this.tail = null;
          } else if (current === this.head) {
            this.head = current.next;
            if(!this.head) return false
            this.head.prev = null;
          } else if (current === this.tail) {
            this.tail = current.prev;
            if(!this.tail) return false;
            this.tail.next = null;
          } else {
            if(!current.prev) return false
            current.prev.next = current.next
            if(!current.next) return false
            current.next.prev = current.prev;
          }
          this.length--;
          return true;
        }
        if(!current.next) return false;
        current = current.next;
      }
     return false;
  }

  public traverseBackward(): T[] {
    let current = this.tail
    const resultnodes : T[] = []
    while(current) {
       resultnodes.push(current.data)
       current = current.prev
    }
    return resultnodes
      
  }

 public search(data: T): Searchresult {
     let current = this.head;
     let index : number = 0
     while(current) {
      if(current.data === data) {
        return {
          index : index,
          flag : true
        } ;
      }
      current = current.next;
      index ++ ;
     }
     return {
      index : null,
      flag : false
     };
 }
}

  