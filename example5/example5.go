
package main

import (
	"fmt"
	"net/http"
  )

  func myHandler(w http.ResponseWriter, r *http.Request, count int) {
	//Print to the screen the number of times this handler has been called
	fmt.Fprintf(w, "<h1>You've called myHandler %v times</h1>", count)
  }

  func mySecondHandler(w http.ResponseWriter, r *http.Request, count int) {
	//Print to the screen the number of times this handler has been called
	fmt.Fprintf(w, "<h1>You've called mySecondHandler %v times</h1>", count)
  }

  func counter(myFunc func(http.ResponseWriter, *http.Request, int)) func(http.ResponseWriter, *http.Request) {
	
	var count int

	return func(w http.ResponseWriter, r *http.Request) {
	  //increment count 
	  count++

	  /*In addition to counting the number of times this func
	  is invoked, we could also collect data from the request
	  (such as the user making the request) and log that 
	  data to a database*/

	  myFunc(w, r, count)
	}
  }
  
  func main() {
	http.HandleFunc("/myHandler", counter(myHandler))
	http.HandleFunc("/mySecondHandler", counter(mySecondHandler))
	http.ListenAndServe(":8091", nil)
  }
  
 
  
