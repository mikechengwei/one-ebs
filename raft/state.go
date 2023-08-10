package raft

type Status int

const (
	Leader Status = iota
	Follower
	Candidate
)

type State struct {
	currentTerm int
	votedFor    *int
	log         []LogEntry
	commitIndex int
	lastApplied int
	nextIndex   []int
	matchIndex  []int
	status      Status
	votesNumber int
}

type LogEntry struct {
	Command interface{}
	Term    int
}

type AppendEntry struct {
	term         int
	leaderId     int
	preLogIndex  int
	prevLogTerm  int
	entries      []LogEntry
	leaderCommit int
}

type AppendEntriesArgs struct {
	Term         int
	LeaderId     int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}
type AppendEntriesReply struct {
	Term          int
	Success       bool
	ConflictTerm  int
	ConflictIndex int
}
