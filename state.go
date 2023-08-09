package raft

type  State struct{
	currentTerm int32
	votedFor int32
	log[] string
	commitIndex int32
	lastApplied int32
	nextIndex[] int32
	matchIndex[] int32
}

type LogEntry struct {
	
}

type AppendEntry struct {
	term int32
	leaderId int32
	preLogIndex int32
	prevLogTerm int32
	entries []LogEntry
	leaderCommit int32
}

type RequestVote struct {
	term int32
	candidateId int32 
	lastLogIndex int32
	lastLogTerm int32
}
