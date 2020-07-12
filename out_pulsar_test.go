package main

import (
	"C"
	"unsafe"
)
import "testing"

func TestFLBPluginRegister(t *testing.T) {
	type args struct {
		ctx unsafe.Pointer
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FLBPluginRegister(tt.args.ctx); got != tt.want {
				t.Errorf("FLBPluginRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFLBPluginInit(t *testing.T) {
	type args struct {
		plugin unsafe.Pointer
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FLBPluginInit(tt.args.plugin); got != tt.want {
				t.Errorf("FLBPluginInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFLBPluginFlushCtx(t *testing.T) {
	type args struct {
		ctx    unsafe.Pointer
		data   unsafe.Pointer
		length C.int
		tag    *C.char
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FLBPluginFlushCtx(tt.args.ctx, tt.args.data, tt.args.length, tt.args.tag); got != tt.want {
				t.Errorf("FLBPluginFlushCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBool(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBool(tt.args.s); got != tt.want {
				t.Errorf("parseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFLBPluginExit(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FLBPluginExit(); got != tt.want {
				t.Errorf("FLBPluginExit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
