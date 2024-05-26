
build-docker: docker-env-down
	docker \
		buildx \
			build \
				-t eldius/grafana-metrics:dev \
				.

run-docker: build-docker docker-env
	docker run -p 8080:8080 --rm eldius/grafana-metrics:dev

docker-env: docker-env-down
	docker compose -f docker-compose.yml up -d

docker-env-down:
	docker compose -f docker-compose.yml down
