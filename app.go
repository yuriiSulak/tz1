package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"math/big"
	"encoding/json"
	"strconv"
)
func factorial( i int)(chan string){
	ch:=make(chan string)
	go func(){
		if i>20{ch<- factorialBig(big.NewInt(int64(i))).String()  }//probably for national aerospace program :))
		res:=int64(1)
		for j:=2;j<=i;j++{
			res=res*int64(j)
		}
		ch<-strconv.FormatInt(res, 10)
	}()
	return ch
}
func factorialBig(n *big.Int) (result *big.Int) {
	//fmt.Println("n = ", n)
	b := big.NewInt(0)
	c := big.NewInt(1)
	if n.Cmp(b) == -1 {
		result = big.NewInt(1)
	}
	if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {
		// return n * factorial(n - 1);
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorialBig(n.Sub(n, c)))
	}
	return
}
type InternalFunction func(int)chan string
func Th( f InternalFunction , SlMain []int  ) []string {
	var out []string
	var chans []chan string//make chans slice
	for i := 0; i < len(SlMain); i++ {
		chans=append(chans, f(SlMain[i]))//this line is asynchronous
	}
	for i := 0; i < len(SlMain); i++ {//this line is synchoronous, waiting for channel one by one
		out=append(out,<-chans[i])
	}
	//splitting each channels
	return out
}
func Calculate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var ab map[string]int
	err := json.NewDecoder(r.Body).Decode(&ab)
	if err != nil {
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	a,isOkA:=ab["a"]
	b,isOkB:=ab["b"]
	if isOkA==false || isOkB==false || a<0 || b<0{
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
		return
	}
	results:=Th(factorial,[]int{a,b})
	w.Write([]byte( `{"a!":`+results[0]+`,"b!":`+results[1]+`}`  ))
}
func main() {
	router := httprouter.New()
	router.POST("/calculate", Calculate)
	log.Fatal(http.ListenAndServe(":8989", router))
}