IMAGE_NAME = "tasnimzotder/artificial-life"

dkr_build:
	docker build -t $(IMAGE_NAME):latest .

dkr_run:
	docker run -it --rm -p 8080:8080 $(IMAGE_NAME):latest

app-install:
	rm -rf /Applications/Artificial\ Life.app/
	fyne install -icon Icon.png

app-pack:
	fyne package -os darwin -icon Icon.png

app-copy:
	cp -r ./artificial-life.app /Applications

.PHONY: dkr_build dkr_run