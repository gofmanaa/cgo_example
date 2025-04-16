#include <stdint.h>
#ifndef ADDLIB_H
#define ADDLIB_H
// Add a value to the counter and return the new value
uint32_t add(uint32_t val);
// Concatenates two NULL-terminated arrays of strings
char* concat(char* msg, char* msg2);
char* concat2(char* msg, char* msg2);

#endif // ADDLIB_H