build:
	docker build -t scrapbuilder .
	docker run --name temp-container scrapbuilder
	docker cp temp-container:/app/bin ./
	docker rm temp-container