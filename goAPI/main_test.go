package main

import(
   "testing"
)

// Test for pattern matching of the query from the article

func TestFound(t *testing.T){
    ans := Found( "Hello" , {"Id":"1","Title":"Hello","Subtitle":"Article Description","Content":"Article Content","Timestamp":"2019-11-10 23:00:00 +0000 UTC"})
    if ans == false{
      t.Error("Expected it to be true")
    }
}
