// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build solaris
// +build solaris

package unix_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"golang.org/x/sys/unix"
)

func TestStatvfs(t *testing.T) {
	if err := unix.Statvfs("", nil); err == nil {
		t.Fatal(`Statvfs("") expected failure`)
	}

	statvfs := unix.Statvfs_t{}
	if err := unix.Statvfs("/", &statvfs); err != nil {
		t.Errorf(`Statvfs("/") failed: %v`, err)
	}

	if t.Failed() {
		mount, err := exec.Command("mount").CombinedOutput()
		if err != nil {
			t.Logf("mount: %v\n%s", err, mount)
		} else {
			t.Logf("mount: %s", mount)
		}
	}
}

func TestSysconf(t *testing.T) {
	n, err := unix.Sysconf(3 /* SC_CLK_TCK */)
	if err != nil {
		t.Fatalf("Sysconf: %v", err)
	}
	t.Logf("Sysconf(SC_CLK_TCK) = %d", n)
}

// Event Ports

func TestBasicEventPort(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "eventport")
	if err != nil {
		t.Fatalf("unable to create a tempfile: %v", err)
	}
	path := tmpfile.Name()
	defer os.Remove(path)

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Failed to stat %s: %v", path, err)
	}
	port, err := unix.NewEventPort()
	if err != nil {
		t.Fatalf("NewEventPort failed: %v", err)
	}
	defer port.Close()
	cookie := stat.Mode()
	err = port.AssociatePath(path, stat, unix.FILE_MODIFIED, cookie)
	if err != nil {
		t.Errorf("AssociatePath failed: %v", err)
	}
	if !port.PathIsWatched(path) {
		t.Errorf("PathIsWatched unexpectedly returned false")
	}
	err = port.DissociatePath(path)
	if err != nil {
		t.Errorf("DissociatePath failed: %v", err)
	}
	err = port.AssociatePath(path, stat, unix.FILE_MODIFIED, cookie)
	if err != nil {
		t.Errorf("AssociatePath failed: %v", err)
	}
	bs := []byte{42}
	tmpfile.Write(bs)
	timeout := new(unix.Timespec)
	timeout.Sec = 1
	pevent, err := port.GetOne(timeout)
	if err == unix.ETIME {
		t.Errorf("GetOne timed out: %v", err)
	}
	if err != nil {
		t.Errorf("GetOne failed: %v", err)
	}
	if pevent.Path != path {
		t.Errorf("Path mismatch: %v != %v", pevent.Path, path)
	}
	err = port.AssociatePath(path, stat, unix.FILE_MODIFIED, cookie)
	if err != nil {
		t.Errorf("AssociatePath failed: %v", err)
	}
	err = port.AssociatePath(path, stat, unix.FILE_MODIFIED, cookie)
	if err == nil {
		t.Errorf("Unexpected success associating already associated path")
	}
}

func TestEventPortFds(t *testing.T) {
	_, path, _, _ := runtime.Caller(0)
	stat, err := os.Stat(path)
	cookie := stat.Mode()
	port, err := unix.NewEventPort()
	if err != nil {
		t.Errorf("NewEventPort failed: %v", err)
	}
	defer port.Close()
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("unable to create a pipe: %v", err)
	}
	defer w.Close()
	defer r.Close()
	fd := r.Fd()

	port.AssociateFd(fd, unix.POLLIN, cookie)
	if !port.FdIsWatched(fd) {
		t.Errorf("FdIsWatched unexpectedly returned false")
	}
	err = port.DissociateFd(fd)
	err = port.AssociateFd(fd, unix.POLLIN, cookie)
	bs := []byte{42}
	w.Write(bs)
	n, err := port.Pending()
	if n != 1 {
		t.Errorf("Pending() failed: %v, %v", n, err)
	}
	timeout := new(unix.Timespec)
	timeout.Sec = 1
	pevent, err := port.GetOne(timeout)
	if err == unix.ETIME {
		t.Errorf("GetOne timed out: %v", err)
	}
	if err != nil {
		t.Errorf("GetOne failed: %v", err)
	}
	if pevent.Fd != fd {
		t.Errorf("Fd mismatch: %v != %v", pevent.Fd, fd)
	}
	var c = pevent.Cookie
	if c == nil {
		t.Errorf("Cookie missing: %v != %v", cookie, c)
		return
	}
	if c != cookie {
		t.Errorf("Cookie mismatch: %v != %v", cookie, c)
	}
	port.AssociateFd(fd, unix.POLLIN, cookie)
	err = port.AssociateFd(fd, unix.POLLIN, cookie)
	if err == nil {
		t.Errorf("unexpected success associating already associated fd")
	}
}

func TestEventPortErrors(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "eventport")
	if err != nil {
		t.Errorf("unable to create a tempfile: %v", err)
	}
	path := tmpfile.Name()
	stat, _ := os.Stat(path)
	os.Remove(path)
	port, _ := unix.NewEventPort()
	defer port.Close()
	err = port.AssociatePath(path, stat, unix.FILE_MODIFIED, nil)
	if err == nil {
		t.Errorf("unexpected success associating nonexistant file")
	}
	err = port.DissociatePath(path)
	if err == nil {
		t.Errorf("unexpected success dissociating unassociated path")
	}
	timeout := new(unix.Timespec)
	timeout.Nsec = 1
	_, err = port.GetOne(timeout)
	if err != unix.ETIME {
		t.Errorf("unexpected lack of timeout")
	}
	err = port.DissociateFd(uintptr(0))
	if err == nil {
		t.Errorf("unexpected success dissociating unassociated fd")
	}
	events := make([]unix.PortEvent, 4)
	_, err = port.Get(events, 5, nil)
	if err == nil {
		t.Errorf("unexpected success calling Get with min greater than len of slice")
	}
	_, err = port.Get(nil, 1, nil)
	if err == nil {
		t.Errorf("unexpected success calling Get with nil slice")
	}
	_, err = port.Get(nil, 0, nil)
	if err == nil {
		t.Errorf("unexpected success calling Get with nil slice")
	}
}

func ExamplePortEvent() {
	type MyCookie struct {
		Name string
	}
	cookie := MyCookie{"Cookie Monster"}
	port, err := unix.NewEventPort()
	if err != nil {
		fmt.Printf("NewEventPort failed: %v\n", err)
		return
	}
	defer port.Close()
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("os.Pipe() failed: %v\n", err)
		return
	}
	defer w.Close()
	defer r.Close()
	fd := r.Fd()

	port.AssociateFd(fd, unix.POLLIN, cookie)

	bs := []byte{42}
	w.Write(bs)
	timeout := new(unix.Timespec)
	timeout.Sec = 1
	pevent, err := port.GetOne(timeout)
	if err != nil {
		fmt.Printf("didn't get the expected event: %v\n", err)
	}

	// Use a type assertion to convert the received cookie back to its original type
	c := pevent.Cookie.(MyCookie)
	fmt.Printf("%s", c.Name)
	//Output: Cookie Monster
}

func TestPortEventSlices(t *testing.T) {
	port, err := unix.NewEventPort()
	if err != nil {
		t.Fatalf("NewEventPort failed: %v", err)
	}
	// Create, associate, and delete 6 files
	for i := 0; i < 6; i++ {
		tmpfile, err := os.CreateTemp("", "eventport")
		if err != nil {
			t.Fatalf("unable to create tempfile: %v", err)
		}
		path := tmpfile.Name()
		stat, err := os.Stat(path)
		if err != nil {
			t.Fatalf("unable to stat tempfile: %v", err)
		}
		err = port.AssociatePath(path, stat, unix.FILE_MODIFIED, nil)
		if err != nil {
			t.Fatalf("unable to AssociatePath tempfile: %v", err)
		}
		err = os.Remove(path)
		if err != nil {
			t.Fatalf("unable to Remove tempfile: %v", err)
		}
	}
	n, err := port.Pending()
	if err != nil {
		t.Errorf("Pending failed: %v", err)
	}
	if n != 6 {
		t.Errorf("expected 6 pending events, got %d", n)
	}
	timeout := new(unix.Timespec)
	timeout.Nsec = 1
	events := make([]unix.PortEvent, 4)
	n, err = port.Get(events, 3, timeout)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if n != 4 {
		t.Errorf("expected 4 events, got %d", n)
	}
	e := events[:n]
	for _, p := range e {
		if p.Events != unix.FILE_DELETE {
			t.Errorf("unexpected event. got %v, expected %v", p.Events, unix.FILE_DELETE)
		}
	}
	n, err = port.Get(events, 3, timeout)
	if err != unix.ETIME {
		t.Errorf("unexpected error. got %v, expected %v", err, unix.ETIME)
	}
	if n != 2 {
		t.Errorf("expected 2 events, got %d", n)
	}
	e = events[:n]
	for _, p := range e {
		if p.Events != unix.FILE_DELETE {
			t.Errorf("unexpected event. got %v, expected %v", p.Events, unix.FILE_DELETE)
		}
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("unable to create a pipe: %v", err)
	}
	port.AssociateFd(r.Fd(), unix.POLLIN, nil)
	port.AssociateFd(w.Fd(), unix.POLLOUT, nil)
	bs := []byte{41}
	w.Write(bs)

	n, err = port.Get(events, 1, timeout)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if n != 2 {
		t.Errorf("expected 2 events, got %d", n)
	}
	err = w.Close()
	if err != nil {
		t.Errorf("w.Close() failed: %v", err)
	}
	err = r.Close()
	if err != nil {
		t.Errorf("r.Close() failed: %v", err)
	}
	err = port.Close()
	if err != nil {
		t.Errorf("port.Close() failed: %v", err)
	}
}

func TestLifreqSetName(t *testing.T) {
	var l unix.Lifreq
	err := l.SetName("12345678901234356789012345678901234567890")
	if err == nil {
		t.Fatal(`Lifreq.SetName should reject names that are too long`)
	}
	err = l.SetName("tun0")
	if err != nil {
		t.Errorf(`Lifreq.SetName("tun0") failed: %v`, err)
	}
}

func TestLifreqGetMTU(t *testing.T) {
	// Find links and their MTU using CLI tooling
	// $ dladm show-link -p -o link,mtu
	// net0:1500
	out, err := exec.Command("dladm", "show-link", "-p", "-o", "link,mtu").Output()
	if err != nil {
		t.Fatalf("unable to use dladm to find data links: %v", err)
	}
	lines := strings.Split(string(out), "\n")
	tc := make(map[string]string)
	for _, line := range lines {
		v := strings.Split(line, ":")
		if len(v) == 2 {
			tc[v[0]] = v[1]
		}
	}
	ip_fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		t.Fatalf("could not open udp socket: %v", err)
	}
	// SIOCGLIFMTU is negative which confuses the compiler if used inline:
	// Using "unix.IoctlLifreq(ip_fd, unix.SIOCGLIFMTU, &l)" results in
	// "constant -1065850502 overflows uint"
	reqnum := int(unix.SIOCGLIFMTU)
	var l unix.Lifreq
	for link, mtu := range tc {
		err = l.SetName(link)
		if err != nil {
			t.Fatalf("Lifreq.SetName(%q) failed: %v", link, err)
		}
		if err = unix.IoctlLifreq(ip_fd, uint(reqnum), &l); err != nil {
			t.Fatalf("unable to SIOCGLIFMTU: %v", err)
		}
		m := l.GetLifruUint()
		if fmt.Sprintf("%d", m) != mtu {
			t.Errorf("unable to read MTU correctly: expected %s, got %d", mtu, m)
		}
	}
}
