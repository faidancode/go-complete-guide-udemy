
GO=go

# =========================
# RUN
# =========================
.PHONY:  cards
cards:
	$(GO) run ./cards/main.go ./cards/deck.go

.PHONY: printslice
printslice:
	$(GO) run ./printslice/main.go