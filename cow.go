package cowsay

import (
	"math/rand"
	"strings"

	"github.com/pkg/errors"
)

// Cow struct!!
type Cow struct {
	eyes        string
	tongue      string
	typ         string
	thoughts    rune
	thinking    bool
	bold        bool
	isAurora    bool
	isRainbow   bool
	ballonWidth int

	buf strings.Builder
}

// New returns pointer of Cow struct that made by options
func New(options ...Option) (*Cow, error) {
	cow := &Cow{
		eyes:        "oo",
		tongue:      "  ",
		thoughts:    '\\',
		typ:         "cows/default.cow",
		ballonWidth: 40,
	}
	for _, o := range options {
		if err := o(cow); err != nil {
			return nil, err
		}
	}
	return cow, nil
}

// Say returns string that said by cow
func (cow *Cow) Say(phrase string) (string, error) {
	mow, err := cow.GetCow()
	if err != nil {
		return "", err
	}

	said := cow.Balloon(phrase) + mow

	if cow.isRainbow {
		return cow.Rainbow(said), nil
	}
	if cow.isAurora {
		return cow.Aurora(rand.Intn(256), said), nil
	}
	return said, nil
}

// Clone returns a copy of cow.
func (cow *Cow) Clone() *Cow {
	ret := new(Cow)
	*ret = *cow
	ret.buf.Reset()
	return ret
}

// Option defined for Options
type Option func(*Cow) error

// Eyes specifies eyes
// You must specify two length string
func Eyes(s string) Option {
	return func(c *Cow) error {
		if l := len(s); l != 2 {
			return errors.New("You should pass 2 length string because cow has only two eyes")
		}
		c.eyes = s
		return nil
	}
}

// Tongue specifies tongue
// You must specify two length string
func Tongue(s string) Option {
	return func(c *Cow) error {
		if l := len(s); l != 2 {
			return errors.New("You should pass 2 length string because cow has only two space on mouth")
		}
		c.tongue = s
		return nil
	}
}

func containCows(t string) bool {
	for _, cow := range AssetNames() {
		if t == cow {
			return true
		}
	}
	return false
}

// Type specify cow type that is file name of .cow
func Type(s string) Option {
	if s == "" {
		s = "cows/default.cow"
	}
	if !strings.HasSuffix(s, ".cow") {
		s += ".cow"
	}
	if !strings.HasPrefix(s, "cows/") {
		s = "cows/" + s
	}
	return func(c *Cow) error {
		if containCows(s) {
			c.typ = s
			return nil
		}
		return errors.Errorf("Could not find %s", s)
	}
}

// Thinking enables thinking mode
func Thinking() Option {
	return func(c *Cow) error {
		c.thinking = true
		return nil
	}
}

// Thoughts Thoughts allows you to specify
// the rune that will be drawn between
// the speech bubbles and the cow
func Thoughts(thoughts rune) Option {
	return func(c *Cow) error {
		c.thoughts = thoughts
		return nil
	}
}

// Random specifies something .cow from cows directory
func Random() Option {
	return func(c *Cow) error {
		c.typ = pickCow()
		return nil
	}
}

func pickCow() string {
	cows := AssetNames()
	n := len(cows)
	rand.Shuffle(n, func(i, j int) {
		cows[i], cows[j] = cows[j], cows[i]
	})
	return cows[rand.Intn(n)]
}

// Bold enables bold mode
func Bold() Option {
	return func(c *Cow) error {
		c.bold = true
		return nil
	}
}

// Aurora enables aurora mode
func Aurora() Option {
	return func(c *Cow) error {
		c.isAurora = true
		return nil
	}
}

// Rainbow enables raibow mode
func Rainbow() Option {
	return func(c *Cow) error {
		c.isRainbow = true
		return nil
	}
}

// BallonWidth specifies ballon size
func BallonWidth(size uint) Option {
	return func(c *Cow) error {
		c.ballonWidth = int(size)
		return nil
	}
}
