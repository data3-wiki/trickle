.PHONY: all decoder_anchor swagger_ui

all: decoder_anchor swagger_ui
	go build -o bin/trickle

decoder_anchor:
	npm install
	npm --prefix js/decoder/anchor run build

swagger_ui:
	npm install
	npm --prefix js/swagger-ui run build