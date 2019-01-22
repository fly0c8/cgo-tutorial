// Falls es doppelte Definitionen gibt, Bauen mit go build main.go, nicht go build!!!
// Prototypes nicht vergessen!!!
// Die Callbackfunctions sind C Typen, deswegen den schirchen C-cast nicht vergessen!!!
package main

/*
void go_callback(int a);
int go_multiply(int,int);

typedef void (*callback_f)(int a);
callback_f g_callback;
static inline void setCallbackHandler(callback_f f) {
	g_callback = f;
}

// Msghandler Callback Beispiel mit char* Param
void go_msghandler();
typedef void (*msghandler_f)(const char*);
msghandler_f g_msghandler;
static void setMsgHandler(msghandler_f f) {
	g_msghandler = f;
}
static inline int multiply(int a, int b) {
	return go_multiply(a,b);
}
static void emit() {
	//g_callback = go_callback;  // funktioniert auch -> go_callback ist eigentlich ein C typ (wg. export)
	if(g_callback) {
		g_callback(666);
	}
	if(g_msghandler) {
		g_msghandler("Hello from MSGBUS!");
	}
}
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	c := C.multiply(C.int(5), C.int(6))
	fmt.Println(c)
	C.setCallbackHandler((C.callback_f)(unsafe.Pointer(C.go_callback)))
	f := (C.callback_f)(unsafe.Pointer(C.go_callback))
	C.setCallbackHandler(f)
	C.setMsgHandler((C.msghandler_f)(unsafe.Pointer(C.go_msghandler)))

	go func() {
		time.Sleep(2 * time.Second)
		C.emit()
	}()
	time.Sleep(4 * time.Second)
}

//export go_multiply
func go_multiply(a C.int, b C.int) C.int {
	return a * b
}

//export go_callback
func go_callback(a C.int) {
	fmt.Println("in go_callback!", a)
}

//export go_msghandler
func go_msghandler(c_msg *C.char) {
	fmt.Println("In go msg handler")
	fmt.Println(C.GoString(c_msg))
}
