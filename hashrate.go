package main

import(
  "fmt"
  "crypto/sha256"
  "strings"
  "time"
  "encoding/hex"
  "strconv"
)



func binToHex( bytes []byte ) string {
	return hex.EncodeToString( bytes )
}

func int64ToStr( num int64 ) string {
  return strconv.FormatInt( num, 10 )
}



func calcHash( data string ) string {
  hashed := sha256.Sum256( []byte(data) )
  return binToHex( hashed[:] )   // note: use [:] to convert [32]byte to slice ([]byte)
}

func computeHashWithProofOfWork( data string, difficulty string ) (int64,string) {
  nonce := int64( 0 )
  for {
    hash := calcHash( int64ToStr(nonce) + data )

    if strings.HasPrefix( hash, difficulty ) {
        return nonce,hash    // bingo! proof of work if hash starts with leading zeros (00)
    } else {
        nonce += 1           // keep trying (and trying and trying)
    }
  }
}



func main() {

  factors := []int{1,2,3,4,5,6}   // add 7 (~10 mins)

  for _, factor := range factors {
     difficulty := strings.Repeat( "0", factor )
     fmt.Printf( "Difficulty: %s (%d bits)\n", difficulty, len(difficulty)*4 )
  }


  for _, factor := range factors {
     difficulty := strings.Repeat( "0", factor )
     fmt.Printf( "\nDifficulty: %s (%d bits)\n", difficulty, len(difficulty)*4 )

     fmt.Println( "Starting search..." )
     t1 := time.Now()
     nonce, _ := computeHashWithProofOfWork( "Hello, Cryptos!", difficulty )
     t2 := time.Now()

     delta := t2.Sub( t1 )
     fmt.Printf( "Elapsed Time: %s, Hashes Calculated: %d\n", delta, nonce )

     if delta.Seconds() > 0.001 {
       hashrate := float64(nonce)/delta.Seconds()
       fmt.Printf( "Hash Rate: %d hashes per second\n", int64(hashrate))
     }
  }
}
