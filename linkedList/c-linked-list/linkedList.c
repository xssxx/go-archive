#include <stdio.h>
#include <stdlib.h>

typedef struct node {
    int val;
    struct node* next;
} Node;

Node* root = NULL;

Node* newNode(int val) {
    Node* node = (Node*) malloc(sizeof(Node));
    node->val = val;
    node->next = NULL;
    return node;
}

void addList(int val) {
    Node* node = newNode(val);

    if (root == NULL) {
        root = node;
        return;
    }
    
    
    Node* curr = root;
    while (curr->next)
        curr = curr->next;

    curr->next = node;
}

void printList(Node* root) {
    if (root == NULL) {
        puts("root is NULL");
        return;
    }

    Node* curr = root;
    while (curr) {
        printf("%d ", curr->val);
        curr = curr->next;
    }
    putchar('\n');
}

void freeList() {
    Node* curr = root;
    while (curr) {
        Node* next = curr->next;
        free(curr);
        curr = next;
    }
    root = NULL;
}

int main() {
    addList(1);
    addList(2);
    addList(3);

    printList(root);
    
    freeList();
    
    printList(root);

    return 0;
}