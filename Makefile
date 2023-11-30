IMAGE_NAME = "tasnimzotder/artificial-life"

dkr_build:
	docker build -t $(IMAGE_NAME):latest .

dkr_run:
	docker run -it --rm -p 8080:8080 $(IMAGE_NAME):latest


.PHONY: dkr_build dkr_run