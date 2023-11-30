IMAGE_NAME = "tasnimzotder/artificial-life"

build:
	docker build -t $(IMAGE_NAME):latest .

run:
	docker run -it --rm $(IMAGE_NAME):latest


.PHONY: build run