#include <stdio.h>
#include <stdlib.h>

typedef struct QNode{
	struct QNode *prv, *next;
	unsigned pageNumber;
} QNode;
