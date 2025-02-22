#include <iostream>

using namespace std;

class Node {
public:
    int val;
    Node* next;  

    Node(int v) : val(v), next(nullptr) {}
};

class LinkedList {
private:
    Node* root;

public:
    LinkedList() {
        root = NULL;
    }

    // destructor free allocated memory
    ~LinkedList() {
        Node* curr = root;
        while (curr) {
            Node* next = curr->next;
            delete curr;
            curr = next;
        }
        root = nullptr;
    }

    void addList(int val) {
        // create new node
        Node* newNode = new Node(val);
        // if root is null make root point to first new node
        if (root == nullptr) {
            root = newNode;
            return;
        }
        // go to last node and make last node point to new node
        Node* curr = root;
        while (curr->next) 
            curr = curr->next;

        curr->next = newNode;
    }

    void printList() {
        Node* curr = root;
        while (curr) {
            printf("%d ", curr->val);
            curr = curr->next;
        }
    }
};

int main() {
    LinkedList l = LinkedList();
    l.addList(1);
    l.addList(2);
    l.addList(3);
    l.addList(4);
    
    l.printList();

    return 0;
}

