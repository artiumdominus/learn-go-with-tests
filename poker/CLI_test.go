package poker_test

import (
	"testing"
	"strings"
	"whatever/m/poker"
	"bytes"
)

var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

}

func userSends(messages ...string) *strings.Reader {
	return strings.NewReader(strings.Join(messages, "\n") + "\n")
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, want)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, want int) {
	t.Helper()
	got := game.StartedWith
	if got != want {
		t.Errorf("got started wiht %d, want %d", got, want)
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, want string) {
	t.Helper()
	got := game.FinishedWith
	if got != want {
		t.Errorf("got finished with %q, want %q", got, want)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
