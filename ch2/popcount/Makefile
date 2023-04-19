run: 
	go test -v .

exec2.3:
	go test -bench=ByTable -run=^a
	go test -bench=ByLoop -run=^a

exec2.4:
	go test -bench=ByTable -run=^a
	go test -bench=ByShifting -run=^a

exec2.5:
	go test -bench=ByTable -run=^a
	go test -bench=ByRightMostOne -run=^a