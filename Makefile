PROG = "./lentille"

run:
	go build && ${PROG} -config=testdata/file1.yaml
