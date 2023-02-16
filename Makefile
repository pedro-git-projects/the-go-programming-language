OBJ_NAME = 
LDFLAGS = 
install:
	$(eval OBJ_NAME += programming)
	$(eval LDFLAGS += "-w -s")
	cd ./cmd/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
run:
	$(eval OBJ_NAME += programming)
	./bin/$(OBJ_NAME)

lissajous:
	$(eval OBJ_NAME += programming)
	./bin/$(OBJ_NAME) >out.gif

doc:
	cd ./cmd/; godoc -http=:6060
