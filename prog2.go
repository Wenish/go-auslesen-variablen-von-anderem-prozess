package main

/*
#cgo LDFLAGS: -lkernel32
#include <windows.h>
#include <stdio.h>

// Funktion zum Lesen des Speichers eines anderen Prozesses
int readMemory(DWORD pid, void* address, void* buffer, SIZE_T size) {
    HANDLE hProcess = OpenProcess(PROCESS_VM_READ, FALSE, pid);
    if (hProcess == NULL) {
        return -1; // Fehler beim Öffnen des Prozesses
    }

    SIZE_T bytesRead;
    BOOL success = ReadProcessMemory(hProcess, address, buffer, size, &bytesRead);
    CloseHandle(hProcess);

    if (!success) {
        return -2; // Fehler beim Lesen des Speichers
    }
    return 0; // Erfolg
}
*/
import "C"
import (
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run prog2.go <pid> <address>")
		os.Exit(1)
	}

	// Zielprozess-ID
	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Ungültige PID:", os.Args[1])
		os.Exit(1)
	}

	// Zieladresse
	address, err := strconv.ParseUint(os.Args[2], 0, 64)
	if err != nil {
		fmt.Println("Ungültige Adresse:", os.Args[2])
		os.Exit(1)
	}

	// Speicher lesen
	var buffer int32
	result := C.readMemory(C.DWORD(pid), unsafe.Pointer(uintptr(address)), unsafe.Pointer(&buffer), C.SIZE_T(unsafe.Sizeof(buffer)))
	if result == -1 {
		fmt.Println("Fehler beim Öffnen des Prozesses. Prüfen Sie die Berechtigungen.")
	} else if result == -2 {
		fmt.Println("Fehler beim Lesen des Speichers. Prüfen Sie die Adresse.")
	} else {
		fmt.Printf("Wert an Adresse 0x%x: %d\n", address, buffer)
	}
}
