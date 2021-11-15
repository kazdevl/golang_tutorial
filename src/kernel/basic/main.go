package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	// fmt os type
	ret, _ := syscall.Sysctl("kern.ostype")
	fmt.Printf("%s\n", ret)

	// fmt memsize
	// https://itchyny.hatenablog.com/entry/2017/12/13/170000
	memsizeStr, err := unix.Sysctl("hw.memsize")
	if err != nil {
		log.Fatal(err)
	}
	memsizeStr += "\x00"
	fmt.Printf("len: %d, content: %s", len(memsizeStr), memsizeStr)
	memsize := uint64(binary.LittleEndian.Uint64([]byte(memsizeStr)))
	fmt.Printf("memsize info\n%d[B]\n%d[KB]\n%d[MB]\n%d[GB]\n", memsize, memsize/1024, memsize/1024/1024, memsize/1024/1024/1024)

	memsizeByte, err := unix.SysctlRaw("hw.memsize")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len: %d, content: %s", len(memsizeByte), memsizeByte)
	memsize = uint64(binary.LittleEndian.Uint64(memsizeByte))
	fmt.Printf("memsize info\n%d[B]\n%d[KB]\n%d[MB]\n%d[GB]\n", memsize, memsize/1024, memsize/1024/1024, memsize/1024/1024/1024)

	// fmt mem used
	vm_stat, err := exec.LookPath("vm_stat")
	fmt.Printf("vm_stat_path: %s\n", vm_stat)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(vm_stat)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(buf.String(), "\n")
	pageSize := uint64(unix.Getpagesize())
	var free, inactive uint64
	for _, line := range lines {
		texts := strings.Split(line, ":")
		if len(texts) != 2 {
			continue
		}
		key := strings.TrimSpace(texts[0])
		value := strings.Trim(texts[1], " .") //第二引数には削除したい文字セットを入力する
		switch key {
		case "Pages free":
			free, _ = strconv.ParseUint(value, 10, 64)
			fmt.Printf("value: %s, free: %d, pages free: %d\n", value, free, free*pageSize)
			free = free * pageSize
		case "Pages inactive":
			inactive, _ = strconv.ParseUint(value, 10, 64)
			fmt.Printf("value: %s, inactive: %d, pages inactive: %d\n", value, inactive, inactive*pageSize)
			inactive = inactive * pageSize
		}
	}
	fmt.Printf("free: %d[B], inactive: %d[B]\n", free, inactive)
	used := memsize - (free + inactive)
	fmt.Printf("used mem info\n%d[B]\n%d[KB]\n%d[MB]\n%d[GB]\n", used, used/1024, used/1024/1024, used/1024/1024/1024)
}
