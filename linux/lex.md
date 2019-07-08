Golang YACC
===========

I started playing around with goyacc a few days ago and not having more than a general understanding of
lex and yacc I knew it would likely have a learning curve. However the docs that I have found around
goyacc are pretty bad which kinda sucks because it's a nice tool. This is a general guide to goyacc
for someone who knowns nothing about lex or yacc.

## YACC Overview

To start with lets take a look at a very basic yacc program that I made. This program only accepts 'h'
and 'h,h' as inputs. This is not the full program but just the start of it so don't try to run this yet.

```
%{
// Simple lex program everything in these brackets can be golang code.
package main

import (
        "bufio"
        "fmt"
        "os"
)
%}

// This would hold the different values that you expect to be produced however we don't need this yet.
%union{}

// Start lets you specify which grammer to start with. Don't worry you will see what this is related to
// in just a few moments.
%start list

// Specify the tokens we expect our lexer to return. We will take a look at the lexer later.
%token HELLO

// This sets precidents of in what order instructions are executed
%left ','

%%

list: hi '\n';

hi:
        hi ',' hi
                { fmt.Println("double hello!!"); }
        | HELLO
                { fmt.Println("why hello!"); };

%%

```
