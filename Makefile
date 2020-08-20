all:
	go build -buildmode=c-shared -o out_pulsar.so

clean:
	rm -rf *.so *.h *~
