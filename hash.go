package main

import(
   "fmt"
   "crypto/sha256"
	 "encoding/hex"
   "strconv"
   "math/big"
 )



func calcHash( bytes []byte ) string {
  h := sha256.New()
  h.Write( bytes )
	hashed := h.Sum( nil )   // check why nil para?
	return hex.EncodeToString(hashed)
}


func main() {
  fmt.Println( calcHash( []byte( "Hello, Cryptos!" ) ))
  fmt.Println( calcHash( []byte( "Hello, Cryptos! - Hello, Cryptos! - Hello, Cryptos!" ) ))

  fmt.Println( calcHash( []byte( "Your Name Here" ) ))
  fmt.Println( calcHash( []byte( "Data Data Data Data" ) ))

  fmt.Println( calcHash( []byte( `Data Data Data Data Data Data
  Data Data Data Data Data Data
  Data Data Data Data Data Data
  Data Data Data Data Data Data
  Data Data Data Data Data Data` ) ))

  fmt.Println( len(calcHash( []byte( "Hello, Cryptos!" ) )))
  fmt.Println( len(calcHash( []byte( "Hello, Cryptos! - Hello, Cryptos! - Hello, Cryptos!" ) )))

  hex := calcHash( []byte( "Hello, Cryptos!" ) )
  fmt.Println( hex )



  n, err := strconv.ParseInt( hex, 16, 64 )
  if err != nil {
    fmt.Printf( "Error in conversion: %s\n", err )
  } else {
    fmt.Println( n )
  }

// Error in conversion: strconv.ParseInt:
//  parsing "33eedea60b0662c66c289ceba71863a864cf84b00e10002ca1069bf58f9362d5":
//  value out of range !!!!
//
//
// todo:   print - biggest largest possible uint64 number

//
// how to deal with uint256
//  see https://ethereum.stackexchange.com/questions/20859/how-generate-number-type-uint256-in-go


  big := new(big.Int)
  big, ok := big.SetString( hex, 16 )
  if !ok {
    fmt.Println( "SetString: error" )
  } else {
    fmt.Println( big )
    // => 23490001543365037720284007500157053051505610714786813679598750288695740555989
    fmt.Println( big.Text( 2 ))
    // => 11001111101110110111101010011000001011000001100110001011000...
  }





  n, _ = strconv.ParseInt( "33ee", 16, 64 )
  fmt.Println( n )
  // => 13294



  n, _ = strconv.ParseInt( "33ee", 16, 64 )
  b  := strconv.FormatInt( n, 2 )
  fmt.Println( b )
  // => 11001111101110

}

/*
33eedea60b0662c66c289ceba71863a864cf84b00e10002ca1069bf58f9362d5
c4b5e2b9685062ecca5d0f6f6ba605b3f99eafed3a3729d2ae1ccaa2b440b1cc
39459289c09c33a7b516bef926c1873c6ecd2e6db09218b065d7465b6736f801
a7bbfc531b2ecf641b9abcd7ad8e50267e1c873e5a396d1919f504973090565a
*/
