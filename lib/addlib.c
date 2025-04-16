#include "addlib.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// Concat msg + msg2 into a single string inside a char**
char* concat(char* msg, char* msg2) {
    size_t len1 = strlen(msg);
    size_t len2 = strlen(msg2);
    char* result = malloc(len1 + len2 + 1); // +1 for null terminator
    if (!result) return NULL;

    strcpy(result, msg);
    strcat(result, msg2);
    return result;
}

char* concat2(char* msg, char* msg2) {
    // Calculate the total length of the combined string
    size_t len1 = strlen(msg);
    size_t len2 = strlen(msg2);
    
    // Allocate enough memory for both strings plus the null terminator
    char* result = (char*)malloc(len1 + len2 + 1);
    if (!result) {
        return NULL; // Return NULL if malloc fails
    }
    
    // Use snprintf to concatenate both strings into result buffer
    snprintf(result, len1 + len2 + 1, "%s%s", msg, msg2);
    
    return result;
}


// Add 1 to the input value
uint32_t add(uint32_t val) {
    return val + 1;
}