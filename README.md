# Backend models
A repo which containes backend models to be shared in golang code, to avoid "Don't repeat yourself".
There are two types of models:
a) Go-internal models placed in pkg folder
b) Protobuf models places in pkg-bin

## How-to create a protobuf model
1) create a ptofile in proto
2) add this file to Makefile
3) run ```make```