# todo: make this use make dep checks

echo:
	(cd maelstrom-echo && go install) && \
	./maelstrom/maelstrom test -w echo --bin ~/go/bin/maelstrom-echo --node-count 1 --time-limit 10

unique-ids:
	(cd maelstrom-unique-ids && go install) && \
	./maelstrom/maelstrom test -w unique-ids --bin ~/go/bin/maelstrom-unique-ids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

# broadcast tests (a-c)
BROADCAST_A := ./maelstrom/maelstrom test -w broadcast --bin ~/go/bin/maelstrom-broadcast --node-count 1 --time-limit 20 --rate 10
BROADCAST_B := ./maelstrom/maelstrom test -w broadcast --bin ~/go/bin/maelstrom-broadcast --node-count 5 --time-limit 20 --rate 10
BROADCAST_C := ./maelstrom/maelstrom test -w broadcast --bin ~/go/bin/maelstrom-broadcast --node-count 5 --time-limit 20 --rate 10 --nemesis partition
broadcast:
	(cd maelstrom-broadcast && go install) && \
	$(BROADCAST_B)

.PHONY: echo unique-ids broadcast