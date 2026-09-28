ifneq (,$(wildcard ./.env))
include .env
export
endif

# make migration-create <name>  ->  the name is the second goal
NAME := $(word 2,$(MAKECMDGOALS))
ifneq (,$(NAME))
$(NAME):
	@:
endif
.PHONY: $(NAME)

migrate:
	goose -dir ./database/migrations postgres "$(DATABASE_URL)" up

rollback:
	goose -dir ./database/migrations postgres "$(DATABASE_URL)" down

status:
	goose -dir ./database/migrations postgres "$(DATABASE_URL)" status

migration-create:
	@test -n "$(NAME)" || (echo "usage: make migration-create <name>" && exit 1)
	@file=database/migrations/$$(date +%Y%m%d%H%M%S)_$(NAME).sql; \
	printf '%s\n' '-- +goose Up' '' '-- +goose Down' > $$file; \
	echo "created $$file"

seed:
	goose -dir ./database/seed postgres "$(DATABASE_URL)" up

.PHONY: migrate rollback status migration-create seed
